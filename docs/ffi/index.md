# FFI — Visão Geral

> Interface para chamadas a bibliotecas nativas (C/C++) a partir de programas Fig.

---

## O que é FFI em Fig?

**FFI** (_Foreign Function Interface_) é o subsistema que permite a programas escritos em Fig invocar funções de bibliotecas nativas compiladas em C (ou qualquer linguagem que exporte símbolos com convenção C). Com FFI, é possível:

- Reutilizar bibliotecas C/C++ existentes sem reescrevê-las;
- Acessar APIs do sistema operacional (POSIX, Win32, etc.);
- Executar código performance-critical em linguagem nativa;
- Criar extensões de alta performance para o runtime Fig.

---

## Arquitetura

O FFI do Fig utiliza um **processo auxiliar** (`ffi-helper`) que age como ponte entre o runtime Fig e as bibliotecas nativas via `dlopen`/`dlsym` (POSIX) ou `LoadLibrary`/`GetProcAddress` (Windows).

```
┌──────────────┐     JSON-RPC      ┌──────────────┐     dlopen/dlsym    ┌──────────────┐
│  Runtime Fig  │ ◄──────────────► │  ffi-helper   │ ◄────────────────► │ Biblioteca C │
│  (ffi_client) │   unix socket    │  (processo    │                    │  (.so/.dylib  │
│               │   ou stdio pipe  │   standalone) │                    │   /.dll)      │
└──────────────┘                   └──────────────┘                    └──────────────┘
```

### Fluxo de uma chamada FFI

1. O código Fig invoca `ffi.load()`, `ffi.sym()` ou `ffi.call()`;
2. O módulo `ffi` (em Go) serializa a requisição em JSON e envia ao `ffi-helper` via **unix socket** (Linux/macOS) ou **stdio pipe** (Windows);
3. O `ffi-helper` decodifica a requisição, executa a operação nativa (dlopen, dlsym, chamada de função, etc.) e retorna o resultado em JSON;
4. O runtime Fig recebe a resposta e a converte de volta para valores Fig.

---

## Quando usar FFI

| Cenário | Exemplo |
|---------|---------|
| Chamar bibliotecas C/C++ existentes | OpenSSL, SQLite, libcurl |
| Código performance-critical | Processamento de imagem, criptografia |
| APIs do sistema operacional | Chamadas POSIX, syscalls, Win32 API |
| Extensões nativas | Plugins compilados em C para Fig |
| Interoperabilidade com outras linguagens | Qualquer linguagem que exporte C ABI |

---

## Referência rápida de funções

| Função | Descrição |
|--------|-----------|
| `ffi.load(path)` | Carrega uma biblioteca nativa (`.so`, `.dylib`, `.dll`) e retorna um handle |
| `ffi.sym(handle, name, retType, [argTypes])` | Resolve um símbolo e registra `argTypes` como array de strings |
| `ffi.call(sym, ...args)` | Invoca uma função nativa previamente resolvida com `sym()` |
| `ffi.alloc(size)` | Aloca `size` bytes de memória no espaço do helper e retorna um `mem_id` |
| `ffi.free(mem_id)` | Libera memória previamente alocada com `alloc()` |
| `ffi.strdup(str)` | Copia uma string Fig para memória C no helper e retorna um `mem_id` |
| `ffi.mem_write(mem_id, offset, data)` | Escreve dados binários (base64) em uma região de memória alocada |
| `ffi.mem_read(mem_id, offset, length)` | Lê dados binários de uma região de memória e retorna em base64 |

### Exemplo mínimo

```fig
use "ffi"

let lib = ffi.load("./libmylib.so")
let add = ffi.sym(lib, "add", "int", ["int", "int"])
let resultado = ffi.call(add, 2, 3)
print(resultado)  // 5
```

---

## Configuração

O FFI deve ser habilitado no arquivo `fig.toml` do projeto:

```toml
[ffi]
enabled = true
helper = "/caminho/para/ffi-helper"
call_timeout = 5000    # opcional, timeout em ms (padrão: 3000)
api_version = "1.0"    # opcional, versão do protocolo
```

| Campo          | Tipo   | Obrigatório | Padrão  | Descrição                          |
|----------------|--------|-------------|---------|------------------------------------|
| `enabled`      | bool   | sim         | `false` | Ativa o módulo FFI                 |
| `helper`       | string | sim         | —       | Caminho para o binário ffi-helper  |
| `call_timeout` | int    | não         | `3000`  | Timeout de chamada em milissegundos|
| `api_version`  | string | não         | `"1.0"` | Versão do protocolo FFI            |

---

## Índice da documentação FFI

| Documento | Descrição |
|-----------|-----------|
| [Guia Rápido](quickstart.md) | Tutorial passo a passo para sua primeira chamada FFI |
| [Protocolo](protocolo.md) | Especificação do protocolo JSON-RPC entre runtime e helper |
| [Tipos](tipos.md) | Coerção e validação de tipos entre Fig e C |
| [Erros](erros.md) | Códigos de erro estruturados retornados pelo helper |
| [Testes](testes.md) | Guia de testes FFI e ferramentas de validação |
| [Helper — Referência](helper.md) | Documentação completa do processo auxiliar ffi-helper |
| [Exemplo SQLite](exemplos/sqlite.md) | Wrapper SQLite com CRUD e prepared statements |

---

## Veja também

- [FFI — Guia do Usuário](../13-ffi.md) — documentação completa com exemplos avançados (structs, callbacks, memória, sandbox)
- [Checklist FFI](../ffi-complete-checklist.md) — roadmap e status de implementação do subsistema FFI
