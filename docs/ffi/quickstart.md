# FFI — Guia Rápido

> Tutorial passo a passo para realizar sua primeira chamada a uma biblioteca C a partir de Fig.

---

## Pré-requisitos

- **Go** instalado (para compilar o `ffi-helper`)
- **GCC** ou outro compilador C (para compilar a biblioteca de exemplo)
- **Fig** instalado e funcionando
- Terminal Linux ou macOS (Windows requer adaptações — veja [Diferenças Windows](#diferenças-windows))

---

## Passo 1 — Criar uma biblioteca C

Crie um arquivo `mylib.c` com uma função simples:

```c
// mylib.c
#include <stdio.h>

int add(int a, int b) {
    return a + b;
}

double multiply(double a, double b) {
    return a * b;
}

void greet(const char* name) {
    printf("Olá, %s!\n", name);
}
```

---

## Passo 2 — Compilar a biblioteca compartilhada

```bash
gcc -shared -fPIC -o libmylib.so mylib.c
```

Verifique que o arquivo `libmylib.so` foi gerado:

```bash
ls -la libmylib.so
# -rwxrwxr-x 1 user user 16K ... libmylib.so
```

> **macOS:** o comando é o mesmo, mas a extensão será `.dylib`:
> ```bash
> gcc -shared -fPIC -o libmylib.dylib mylib.c
> ```

---

## Passo 3 — Compilar o ffi-helper

O `ffi-helper` é o processo auxiliar que executa as chamadas nativas. Compile-o a partir do diretório do projeto Fig:

```bash
cd tools/ffi-helper
go build -o ffi-helper .
```

Anote o caminho completo do binário gerado (por exemplo, `/home/user/FigLang/tools/ffi-helper/ffi-helper`).

---

## Passo 4 — Configurar o `fig.toml`

No diretório do seu projeto Fig, crie ou edite o arquivo `fig.toml`:

```toml
[ffi]
enabled = true
helper = "/caminho/completo/para/ffi-helper"
call_timeout = 5000
```

| Campo          | Descrição                                      |
|----------------|-------------------------------------------------|
| `enabled`      | Deve ser `true` para ativar o módulo FFI        |
| `helper`       | Caminho absoluto para o binário `ffi-helper`    |
| `call_timeout` | Timeout em milissegundos (padrão: 3000)         |

---

## Passo 5 — Escrever o programa Fig

Crie um arquivo `programa.fig`:

```fig
use "ffi"

// 1. Carregar a biblioteca
let lib = ffi.load("./libmylib.so")

// 2. Resolver os símbolos (funções)
let add = ffi.sym(lib, "add", "int", ["int", "int"])
let multiply = ffi.sym(lib, "multiply", "double", ["double", "double"])
let greet = ffi.sym(lib, "greet", "void", ["string"])

// 3. Chamar as funções
let soma = ffi.call(add, 10, 32)
print("10 + 32 = " + soma)  // 10 + 32 = 42

let produto = ffi.call(multiply, 3.14, 2.0)
print("3.14 * 2 = " + produto)  // 3.14 * 2 = 6.28

ffi.call(greet, "Fig")  // Olá, Fig!
```

---

## Passo 6 — Executar

```bash
fig run programa.fig
```

Saída esperada:

```
10 + 32 = 42
3.14 * 2 = 6.28
Olá, Fig!
```

---

## Exemplo com memória: alloc, write, read, free

Para operações que exigem gerenciamento manual de memória:

```fig
use "ffi"

let lib = ffi.load("./libmylib.so")

// Alocar 64 bytes de memória no helper
let mem = ffi.alloc(64)

// Escrever dados (base64-encoded)
ffi.mem_write(mem, 0, "SGVsbG8gRkZJ")  // "Hello FFI" em base64

// Ler dados de volta
let data = ffi.mem_read(mem, 0, 9)
print(data)  // "SGVsbG8gRkZJ"

// Liberar memória
ffi.free(mem)
```

---

## Exemplo com strings C: strdup

```fig
use "ffi"

// Copiar uma string Fig para memória C
let str_id = ffi.strdup("Olá mundo")

// O str_id pode ser passado para funções C que esperam char*
// ...

// Liberar a string quando não for mais necessária
ffi.free(str_id)
```

---

## Troubleshooting

### Erros comuns e seus códigos

| Erro | Código | Causa provável | Solução |
|------|--------|----------------|---------|
| Biblioteca não encontrada | `ERR_DLOPEN_FAILED` | Caminho incorreto ou biblioteca ausente | Verifique o caminho passado a `ffi.load()` e se o `.so` existe |
| Símbolo não encontrado | `ERR_DLSYM_FAILED` | Nome da função incorreto ou não exportado | Verifique o nome exato da função no código C |
| Tipo incompatível | `ERR_TYPE_ERROR` | Argumento com tipo incorreto | Confira os `argTypes` em `ffi.sym()` |
| Número de args errado | `ERR_UNSUPPORTED_ARGS` | Assinatura de wrapper não disponível | Verifique a assinatura suportada (quantidade e tipos de args) |
| Helper não responde | (timeout) | Helper não está rodando ou caminho errado | Confira `helper` no `fig.toml` e se o binário existe |
| JSON inválido | `ERR_INVALID_JSON` | Bug interno ou protocolo corrompido | Ative debug: `FFI_LOG_LEVEL=debug fig run ...` |
| Versão incompatível | `ERR_VERSION_MISMATCH` | Helper e runtime com versões diferentes | Recompile o helper: `go build -o ffi-helper ./tools/ffi-helper` |
| Memória inválida | `ERR_INVALID_MEM_ID` | `mem_id` não existe ou já foi liberado | Não use `mem_id` após `ffi.free()` |

### Ativando logs de debug

Para diagnosticar problemas, ative o logging detalhado:

```bash
# Via variável de ambiente
FFI_LOG_LEVEL=debug fig run programa.fig

# Ou diretamente no helper
./ffi-helper --server --log-level debug
```

Níveis disponíveis: `error`, `warn`, `info`, `debug`.

---

## Diferenças Windows

No Windows, o FFI utiliza **stdio pipes** em vez de unix sockets para comunicação com o helper. As adaptações necessárias são:

- A biblioteca compartilhada deve ser um `.dll` (compilada com MinGW ou MSVC);
- O helper é iniciado com a flag `--server` (modo stdio) em vez de `--socket`;
- O `fig.toml` não muda — o runtime detecta automaticamente a plataforma.

```bash
# Compilar DLL no Windows (MinGW)
gcc -shared -o mylib.dll mylib.c
```

---

## Próximos passos

- [FFI — Visão Geral](index.md) — arquitetura e referência completa de funções
- [Protocolo](protocolo.md) — detalhes do protocolo JSON-RPC
- [Tipos](tipos.md) — coerção e validação de tipos Fig ↔ C
- [Erros](erros.md) — tabela completa de códigos de erro
- [Helper — Referência](helper.md) — documentação completa do processo auxiliar
