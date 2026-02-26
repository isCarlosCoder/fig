# FFI — Sandbox e Segurança

O módulo FFI executa código nativo fora do controle do runtime Fig. O sistema de sandbox oferece camadas de proteção para limitar o que bibliotecas carregadas podem fazer.

## Arquitetura de isolamento

```
┌──────────────────────────────┐
│  Fig Runtime                 │
│  (processo principal)         │
│                              │
│  ┌────────────────────────┐  │
│  │  FFI Client            │  │
│  │  (builtins/ffi*.go)    │  │
│  └──────────┬─────────────┘  │
│             │ socket/pipe     │
└─────────────┼────────────────┘
              │
┌─────────────┼────────────────┐
│  FFI Helper │                │
│  (processo separado)          │
│             ▼                │
│  ┌────────────────────────┐  │
│  │  dlopen / dlsym / call │  │
│  │  (código C nativo)     │  │
│  └────────────────────────┘  │
└──────────────────────────────┘
```

O helper roda como **processo separado** — um crash no código C não derruba o runtime Fig.

## Configuração de sandbox

```toml
[ffi]
enabled = true
helper = "./ffi-helper"

[ffi.sandbox]
# Bibliotecas permitidas (glob patterns)
allowed_libs = ["./libs/*.so", "/usr/lib/libsqlite3.so"]

# Bloquear carregamento de bibliotecas do sistema
block_system_libs = false

# Timeout de chamada (ms) — proteção contra loops infinitos
# 0 = sem limite (usar com cuidado); omitir usa o padrão de 3000ms
call_timeout = 5000

# Limite de memória alocável (bytes) — 0 = sem limite
memory_limit = 0

# Limite de bibliotecas carregadas simultaneamente
max_loaded_libs = 10
```

## Camadas de proteção

### 1. Isolamento de processo

O helper roda como processo separado. Falhas (SIGSEGV, SIGBUS, etc.) no código C são contidas no processo do helper.

- **Crash recovery:** Se o helper morrer, o runtime Fig detecta e pode reiniciá-lo automaticamente.
- **Sem acesso ao heap Fig:** O helper não tem acesso direto à memória do runtime.

### 2. Timeout de chamadas

Toda chamada FFI tem um timeout configurável (padrão: 3000ms):

```toml
[ffi]
call_timeout = 5000  # 5 segundos
```

Se a função C não retornar dentro do timeout, a chamada falha com erro.

### 3. Lista de bibliotecas permitidas

Restrinja quais bibliotecas podem ser carregadas:

```toml
[ffi.sandbox]
allowed_libs = ["./libs/*.so"]  # Apenas libs locais
```

Tentar carregar uma biblioteca fora da lista resulta em erro `ERR_DLOPEN_FAILED`.

### 4. Limites de recursos

| Configuração | Descrição | Padrão |
|-------------|-----------|--------|
| `memory_limit` | Máximo de bytes alocáveis via `ffi.alloc()` | 0 (sem limite) |
| `max_loaded_libs` | Máximo de bibliotecas carregadas simultaneamente | 10 |
| `call_timeout` | Timeout por chamada (ms) | 3000 |

### Inspecionando o estado em tempo de execução

O builtin `ffi.sandbox_status()` retorna um objeto com os limites
configurados e os contadores atuais (bibliotecas carregadas, alocações
ativas, reinícios do helper). Isso é útil para monitorar ou alertar em scripts
Fig.

```fig
use "ffi"
let st = ffi.sandbox_status()
print("loaded libs", st.loaded_libs)
print("live allocs", st.live_allocs)
```

Os campos retornados são idênticos aos definidos em `[ffi.sandbox]` mais os
contadores:
`loaded_libs`, `live_allocs`, `restarts`.

## Nível de log para auditoria

Use `--log-level info` no helper para auditar todas as operações:

```bash
FFI_LOG_LEVEL=info ./ffi-helper --socket /tmp/fig-ffi.sock
```

Isso logará cada `load`, `sym`, `call`, `alloc`, `free` — útil para detectar comportamento inesperado.

## Recomendações

1. **Sempre use a sandbox em produção** — configure `allowed_libs` para limitar o que pode ser carregado
2. **Defina timeouts** — proteja contra funções C que entram em loop infinito
3. **Teste com ASAN/Valgrind** — detecte erros de memória antes de ir para produção
4. **Revise bibliotecas C** — código nativo pode fazer qualquer coisa (sistema de arquivos, rede, etc.)
5. **Monitore logs** — use `log-level info` para auditar operações FFI

## Ver também

- [Guia de segurança](guia-seguranca.md)
- [Helper — flags e configuração](helper.md)
- [Gerenciamento de memória](memoria.md)
