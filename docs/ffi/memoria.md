# FFI — Gerenciamento de Memória

O módulo FFI dá acesso direto à memória nativa. É responsabilidade do programador gerenciar alocações e evitar vazamentos.

## Funções de memória

| Função | Descrição |
|--------|-----------|
| `ffi.alloc(tamanho)` | Aloca `tamanho` bytes (retorna ID de memória) |
| `ffi.free(id)` | Libera memória previamente alocada |
| `ffi.strdup(string)` | Duplica uma string Fig como C string (aloca memória) |
| `ffi.mem_write(id, bytes)` | Escreve bytes (base64) na região alocada |
| `ffi.mem_read(id, offset, len)` | Lê bytes de uma região alocada |

## Modelo de ownership

```
┌─────────┐   alloc/strdup   ┌────────────┐
│  Fig     │ ───────────────→ │  Helper    │
│ runtime  │                  │  (C heap)  │
│          │ ←─ retorna ID ── │            │
│          │                  │            │
│          │   free(ID)       │            │
│          │ ───────────────→ │  libera    │
└─────────┘                  └────────────┘
```

**Regra fundamental:** Quem aloca, libera.

- `ffi.alloc()` e `ffi.strdup()` alocam no heap C, via o helper.
- O ID retornado é um handle opaco (`"m-1"`, `"m-2"`, etc.).
- `ffi.free(id)` libera a memória. Após o free, o ID é inválido.
- Usar um ID após free resulta em `ERR_INVALID_MEM_ID`.

## Exemplo: alloc + write + read

```js
use "ffi"

// Alocar 256 bytes
var buf = ffi.alloc(256)

// Escrever dados (base64 encoded)
ffi.mem_write(buf, "SGVsbG8gRmlnIQ==")  // "Hello Fig!" em base64

// Ler 10 bytes a partir do offset 0
var data = ffi.mem_read(buf, 0, 10)
print(data)  // base64 dos bytes lidos

// Liberar
ffi.free(buf)
```

## Exemplo: strdup

```js
use "ffi"

// Duplicar string para memória C
var cstr = ffi.strdup("Hello, FFI!")

// cstr agora é um ponteiro para uma C string no heap
// Pode ser passado para funções C que esperam char*
var lib = ffi.load("./libmylib.so")
var process = ffi.sym(lib, "process_string", "int", ["string"])
ffi.call(process, cstr)

// Liberar quando não precisar mais
ffi.free(cstr)
```

## Erros comuns

| Erro | Causa | Solução |
|------|-------|---------|
| `ERR_INVALID_MEM_ID` | ID inexistente ou já liberado | Verifique se o ID existe e não foi liberado |
| `ERR_MALLOC_FAILED` | Falha ao alocar memória | Sistema sem memória; reduza o tamanho da alocação |
| `ERR_OUT_OF_BOUNDS` | Offset/tamanho fora da região | Verifique limites com o tamanho alocado |
| `ERR_INVALID_BASE64` | Dados base64 inválidos | Verifique a codificação dos bytes |

## Boas práticas

1. **Sempre libere memória alocada** — use `ffi.free()` assim que não precisar mais
2. **Não use IDs após free** — o ID se torna inválido
3. **Prefira `ffi.strdup()`** para strings — mais seguro que alloc + write manual
4. **Alocações grandes** — considere alocar uma vez e reutilizar em vez de alocar/liberar repetidamente

## Detecção de vazamentos

### AddressSanitizer (ASAN)

Compile o helper com ASAN para detectar vazamentos:

```bash
CGO_ENABLED=1 \
CC="gcc" \
CGO_CFLAGS="-fsanitize=address -g" \
CGO_LDFLAGS="-fsanitize=address" \
go build -o ffi-helper-asan ./tools/ffi-helper
```

### Valgrind

```bash
valgrind --leak-check=full ./ffi-helper --server
```

## Ver também

- [Protocolo — comandos de memória](protocolo.md)
- [Erros](erros.md)
- [Guia de segurança](guia-seguranca.md)
