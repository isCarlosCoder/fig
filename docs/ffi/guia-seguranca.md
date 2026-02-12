# FFI — Guia de Segurança

Boas práticas e padrões seguros para usar FFI no Fig.

## Regra de ouro

> **Código nativo é perigoso.** FFI dá acesso direto a `dlopen`/`dlsym` — qualquer biblioteca C pode executar código arbitrário, acessar o sistema de arquivos, rede, etc. Trate toda biblioteca FFI como código não-confiável até ser auditada.

## Checklist de segurança

- [ ] Biblioteca compilada com `-fPIC` e sem warnings
- [ ] Testada com AddressSanitizer (sem erros)
- [ ] Testada com Valgrind (sem leaks)
- [ ] Não usa variáveis globais mutáveis (thread safety)
- [ ] Todas as alocações são pareadas com free
- [ ] Funções documentam quem é responsável por liberar memória
- [ ] `fig.toml` configurado com sandbox habilitada
- [ ] Timeout definido para todas as chamadas

## Erros comuns e soluções

### 1. Vazamento de memória

**Problema:** Chamar `ffi.alloc()` ou `ffi.strdup()` sem `ffi.free()`.

```js
// ❌ ERRADO — vazamento!
var buf = ffi.alloc(1024)
// ... usa buf ...
// Esqueceu de liberar!

// ✅ CORRETO
var buf = ffi.alloc(1024)
// ... usa buf ...
ffi.free(buf)
```

### 2. Use-after-free

**Problema:** Usar um ID de memória após liberá-lo.

```js
// ❌ ERRADO — use-after-free!
var buf = ffi.alloc(1024)
ffi.free(buf)
ffi.mem_write(buf, data)  // ERR_INVALID_MEM_ID

// ✅ CORRETO — não usar após free
var buf = ffi.alloc(1024)
ffi.mem_write(buf, data)
ffi.free(buf)
```

### 3. Buffer overflow

**Problema:** Escrever além do tamanho alocado.

```js
// ❌ PERIGOSO — possível overflow
var buf = ffi.alloc(10)
ffi.mem_write(buf, dados_grandes)  // ERR_OUT_OF_BOUNDS

// ✅ CORRETO — alocar o tamanho necessário
var tamanho = calcular_tamanho(dados)
var buf = ffi.alloc(tamanho)
ffi.mem_write(buf, dados)
```

### 4. Biblioteca maliciosa

**Problema:** Carregar uma `.so` não confiável.

```toml
# ✅ Restringir quais libs podem ser carregadas
[ffi.sandbox]
allowed_libs = ["./libs/*.so"]
block_system_libs = true
```

### 5. Timeout não configurado

**Problema:** Função C entra em loop infinito, travando o programa.

```toml
# ✅ Sempre definir timeout
[ffi]
call_timeout = 5000
```

## Testando com ASAN

Compile bibliotecas e o helper com AddressSanitizer:

```bash
# Compilar lib C com ASAN
gcc -shared -fPIC -fsanitize=address -g -o libmylib.so mylib.c

# Compilar helper com ASAN
CGO_ENABLED=1 CC="gcc" \
CGO_CFLAGS="-fsanitize=address -g" \
CGO_LDFLAGS="-fsanitize=address" \
go build -o ffi-helper-asan ./tools/ffi-helper

# Rodar com ASAN
ASAN_OPTIONS=detect_leaks=1 ./ffi-helper-asan --server
```

## Testando com Valgrind

```bash
valgrind --leak-check=full --show-leak-kinds=all ./ffi-helper --server
```

## Padrões recomendados

### Wrapper seguro

Encapsule chamadas FFI em funções Fig que validam entrada e garantem cleanup:

```js
use "ffi"

func safe_load(path) {
    if !ffi.enabled() {
        print("ERRO: FFI não habilitado")
        return nil
    }
    return ffi.load(path)
}

func safe_call_with_buf(sym, tamanho, dados) {
    var buf = ffi.alloc(tamanho)
    ffi.mem_write(buf, dados)
    var resultado = ffi.call(sym, buf)
    ffi.free(buf)  // Sempre liberar
    return resultado
}
```

### Validação de tipos

Sempre declare `argTypes` em `ffi.sym()` para habilitar coerção e validação:

```js
// ✅ Com tipos declarados — validação automática
var add = ffi.sym(lib, "add", "int", ["int", "int"])

// ❌ Sem tipos — sem validação, possível crash
var add = ffi.sym(lib, "add")
```

## Ver também

- [Sandbox — configuração](sandbox.md)
- [Memória — gerenciamento](memoria.md)
- [Erros — códigos](erros.md)
- [Testes — como rodar](testes.md)
