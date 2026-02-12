# FFI Helper — Referência

> Documentação completa do processo auxiliar `ffi-helper`, responsável por executar chamadas a bibliotecas nativas em nome do runtime Fig.

---

## Arquitetura

O `ffi-helper` é um **processo standalone** escrito em Go (com CGo) que atua como ponte entre o runtime Fig e bibliotecas nativas C. Ele:

- Recebe requisições em formato **JSON-RPC** (uma linha JSON por requisição);
- Executa operações nativas usando `dlopen`/`dlsym` (POSIX) ou `LoadLibrary`/`GetProcAddress` (Windows);
- Retorna resultados em JSON;
- É executado como processo filho do runtime Fig, isolando falhas em código nativo do processo principal.

```
 Runtime Fig                    ffi-helper                    Biblioteca C
┌───────────┐   JSON request   ┌───────────┐   CGo/dlopen   ┌───────────┐
│ffi_client  │ ──────────────► │  serve()   │ ─────────────► │  .so/.dll │
│            │ ◄────────────── │            │ ◄───────────── │           │
└───────────┘   JSON response  └───────────┘   resultado     └───────────┘
      │                              │
   unix socket                  wrappers.c
   ou stdio pipe              (chamadas tipadas)
```

### Por que um processo separado?

- **Isolamento de falhas:** um _segfault_ na biblioteca C derruba apenas o helper, não o runtime Fig inteiro;
- **Segurança:** o helper pode ser sandboxed com limites de CPU, memória e syscalls;
- **Portabilidade:** unix sockets no POSIX, stdio pipes no Windows;
- **Simplicidade:** o protocolo baseado em JSON facilita depuração e extensibilidade.

---

## Compilação

A partir da raiz do projeto Fig:

```bash
cd tools/ffi-helper
go build -o ffi-helper .
```

Ou diretamente:

```bash
go build -o ffi-helper ./tools/ffi-helper
```

> **Dependências CGo:** o compilador C (GCC ou Clang) é necessário. No Linux, o pacote `libc-dev` deve estar instalado.

---

## Flags de linha de comando

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `--server` | bool | `false` | Modo **stdio**: lê JSON de `stdin` e responde em `stdout`. Usado no Windows e para testes. |
| `--socket <path>` | string | `""` | Modo **unix socket**: escuta conexões no caminho especificado. Usado em Linux/macOS. |
| `--log-level <level>` | string | `"warn"` | Nível de log: `error`, `warn`, `info`, `debug`. Sobrescreve `FFI_LOG_LEVEL`. |

### Exemplos de uso

```bash
# Modo stdio (Windows ou testes)
./ffi-helper --server

# Modo unix socket (Linux/macOS)
./ffi-helper --socket /tmp/fig-ffi.sock

# Com debug habilitado
./ffi-helper --server --log-level debug

# Sem flags: apenas exibe mensagem de ajuda
./ffi-helper
# ffi-helper: stub helper (run with --server or --socket to act as RPC helper)
```

---

## Variáveis de ambiente

| Variável | Valores | Padrão | Descrição |
|----------|---------|--------|-----------|
| `FFI_LOG_LEVEL` | `error`, `warn`, `info`, `debug` | `warn` | Define o nível de log quando `--log-level` não é especificado |

**Prioridade de resolução do nível de log:**

1. Flag `--log-level` (maior prioridade)
2. Variável `FFI_LOG_LEVEL`
3. Padrão: `warn`

---

## Níveis de log

| Nível | Valor | O que é registrado |
|-------|-------|--------------------|
| `error` | 0 | Apenas erros fatais e falhas críticas |
| `warn` | 1 | Avisos importantes (padrão) |
| `info` | 2 | Informações operacionais (conexões, carregamentos) |
| `debug` | 3 | Detalhes completos de cada requisição/resposta |

Os logs são escritos em `stderr` no formato:

```
ffi-helper: [LEVEL] mensagem
```

---

## Protocolo

### Versão

A versão atual do protocolo é **1.0**. A constante está definida em:

- **Helper:** `FFIProtocolVersion = "1.0"` em `tools/ffi-helper/main.go`
- **Client:** `FFIProtocolVersion = "1.0"` em `builtins/ffi_client.go`

### Handshake

Ao conectar, o runtime envia um comando `handshake` para verificar compatibilidade:

```json
{"cmd": "handshake", "version": "1.0", "id": 1}
```

Resposta de sucesso:

```json
{"ok": true, "version": "1.0", "id": 1}
```

Se as versões não forem compatíveis, o helper retorna erro `ERR_VERSION_MISMATCH`.

### Formato das mensagens

**Requisição:** uma linha JSON com os campos `cmd`, `id` e parâmetros específicos do comando.

```json
{"cmd": "load", "path": "./libmylib.so", "id": 2}
```

**Resposta:** uma linha JSON com `ok`, `id` e campos de resultado ou erro.

```json
{"ok": true, "handle": "h1", "id": 2}
```

---

## Comandos suportados

| Comando | Descrição | Parâmetros | Retorno |
|---------|-----------|------------|---------|
| `handshake` | Verifica compatibilidade de versão | `version` | `version` |
| `ping` | Teste de conectividade | — | `{"ok": true}` |
| `load` | Carrega biblioteca nativa (`dlopen`) | `path` | `handle` |
| `sym` | Resolve símbolo na biblioteca (`dlsym`) | `handle`, `name`, `ret_type` | `sym_id` |
| `call` | Invoca uma função nativa | `sym_id`, `args`, `arg_types` (opcional) | `result` |
| `alloc` | Aloca memória no espaço do helper | `size` | `mem_id` |
| `free` | Libera memória previamente alocada | `mem_id` | `{"ok": true}` |
| `strdup` | Copia string para memória C | `value` | `mem_id` |
| `mem_write` | Escreve dados binários (base64) em memória | `mem_id`, `offset`, `data` | `{"ok": true}` |
| `mem_read` | Lê dados binários de memória (retorna base64) | `mem_id`, `offset`, `length` | `data` |

### Exemplos de requisições

#### load

```json
{"cmd": "load", "path": "./libmylib.so", "id": 1}
// Resposta: {"ok": true, "handle": "h1", "id": 1}
```

#### sym

```json
{"cmd": "sym", "handle": "h1", "name": "add", "ret_type": "int", "id": 2}
// Resposta: {"ok": true, "sym_id": "s1", "id": 2}
```

#### call

```json
{"cmd": "call", "sym_id": "s1", "args": [10, 32], "id": 3}
// Resposta: {"ok": true, "result": 42, "id": 3}
```

#### alloc / mem_write / mem_read / free

```json
{"cmd": "alloc", "size": 64, "id": 4}
// Resposta: {"ok": true, "mem_id": "m1", "id": 4}

{"cmd": "mem_write", "mem_id": "m1", "offset": 0, "data": "SGVsbG8=", "id": 5}
// Resposta: {"ok": true, "id": 5}

{"cmd": "mem_read", "mem_id": "m1", "offset": 0, "length": 5, "id": 6}
// Resposta: {"ok": true, "data": "SGVsbG8=", "id": 6}

{"cmd": "free", "mem_id": "m1", "id": 7}
// Resposta: {"ok": true, "id": 7}
```

---

## Diferenças entre POSIX e Windows

| Aspecto | POSIX (Linux/macOS) | Windows |
|---------|---------------------|---------|
| **Comunicação** | Unix socket (`--socket`) | Stdio pipe (`--server`) |
| **Carregamento de libs** | `dlopen` / `dlsym` | `LoadLibrary` / `GetProcAddress` |
| **Extensão de biblioteca** | `.so` (Linux), `.dylib` (macOS) | `.dll` |
| **Flags de compilação** | `-ldl -rdynamic` (Linux), `-ldl` (macOS) | `-lkernel32` |
| **Socket path** | Ex.: `/tmp/fig-ffi-<pid>.sock` | N/A (usa stdio) |
| **Sandbox** | `prlimit` (RLIMIT_AS, RLIMIT_CPU) | Não disponível nativamente |

### Header portável

O helper utiliza o header `dl_portable.h` que abstrai as diferenças entre plataformas:

- **POSIX:** usa `<dlfcn.h>` com `dlopen`, `dlsym`, `dlclose`, `dlerror`
- **Windows:** mapeia para `LoadLibrary`, `GetProcAddress`, `FreeLibrary`, `FormatMessage`

---

## Códigos de erro

Quando uma operação falha, o helper retorna um erro estruturado:

```json
{
  "ok": false,
  "error": {"code": "ERR_DLOPEN_FAILED", "message": "dlopen: libfoo.so: cannot open shared object file"},
  "id": 1
}
```

| Código | Descrição |
|--------|-----------|
| `ERR_INVALID_JSON` | JSON malformado na requisição |
| `ERR_UNKNOWN_CMD` | Comando não reconhecido |
| `ERR_MISSING_PARAM` | Parâmetro obrigatório ausente |
| `ERR_INVALID_HANDLE` | Handle de biblioteca inválido |
| `ERR_INVALID_SYMBOL` | Identificador de símbolo inválido |
| `ERR_DLOPEN_FAILED` | Falha ao carregar biblioteca (`dlopen`) |
| `ERR_DLSYM_FAILED` | Falha ao resolver símbolo (`dlsym`) |
| `ERR_CALL_FAILED` | Falha na execução da função nativa |
| `ERR_TYPE_ERROR` | Tipo de argumento incompatível |
| `ERR_UNSUPPORTED_ARGS` | Assinatura de argumentos não suportada |
| `ERR_MALLOC_FAILED` | Falha na alocação de memória C |
| `ERR_INVALID_MEM_ID` | Identificador de memória inválido |
| `ERR_OUT_OF_BOUNDS` | Acesso fora dos limites da memória alocada |
| `ERR_INVALID_BASE64` | Dados base64 inválidos em `mem_write` |
| `ERR_VERSION_MISMATCH` | Versão do protocolo incompatível |

Para detalhes completos, consulte [Códigos de Erro FFI](erros.md).

---

## Veja também

- [FFI — Visão Geral](index.md) — arquitetura e referência de funções
- [Guia Rápido](quickstart.md) — tutorial passo a passo
- [Protocolo](protocolo.md) — especificação completa do protocolo JSON-RPC
- [Erros](erros.md) — tabela completa de códigos de erro
- [FFI — Guia do Usuário](../13-ffi.md) — documentação completa com exemplos avançados
