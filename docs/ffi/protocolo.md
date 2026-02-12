# Protocolo FFI — Especificação

> Documentação do protocolo JSON-RPC usado entre o runtime Fig e o `ffi-helper`.

---

## Gerenciamento de memória no helper

O `ffi-helper` comunica-se com bibliotecas C usando `cgo`, o que exige alocação e liberação manual de memória C (`C.CString`, `C.free`).

### Regra fundamental

**Toda memória C alocada dentro de uma iteração do loop de requests deve ser liberada antes de iniciar a próxima iteração.**

Nunca usar `defer C.free(...)` dentro do loop `for req := range reqCh`, pois `defer` em Go é vinculado ao escopo da função — e não do bloco. Isso significa que defers dentro de um loop só executam quando a função retorna, causando vazamento de memória proporcional ao número de requests processados.

### Padrão correto para `load` e `sym`

```go
// ✅ CORRETO — free imediato após uso
cp := C.CString(path)
hdl := C.dl_open(cp)
C.free(unsafe.Pointer(cp)) // libera imediatamente

// ❌ INCORRETO — defer acumula até serve() retornar
cp := C.CString(path)
defer C.free(unsafe.Pointer(cp)) // NUNCA dentro de loop
hdl := C.dl_open(cp)
```

### Padrão correto para `call` tipo `string`

Quando múltiplos `C.CString` são criados para argumentos de uma chamada de função C, usa-se um slice `toFree` com cleanup via IIFE (immediately invoked function expression):

```go
case "string", "str":
    if err := func() error {
        var toFree []*C.char
        defer func() {
            for _, p := range toFree {
                C.free(unsafe.Pointer(p))
            }
        }()
        // O defer agora é vinculado à função anônima,
        // e executa ao final de CADA iteração
        // ...
        return nil
    }(); err != nil {
        return err
    }
```

### Por que não usar `defer` diretamente?

Em Go, `defer` é function-scoped. Dentro de um `for` loop:

```go
for req := range reqCh {
    cp := C.CString(...)
    defer C.free(unsafe.Pointer(cp)) // acumula sem executar!
    // ...
}
// todos os defers executam aqui, quando serve() retorna
```

Se o helper processar 10.000 requests, terá 10.000 defers pendentes, cada um mantendo memória C alocada. Com o padrão IIFE, o defer é vinculado à função anônima e executa ao final de cada iteração.

---

## Tratamento de erros no helper

O helper utiliza funções seguras para extração de tipos, evitando panics por type assertion inválida.

### Funções de extração segura

```go
// safeFloat64 — extrai float64 de interface{}, aceita float64, int, int64
func safeFloat64(v interface{}) (float64, error)

// safeInt — extrai int via safeFloat64 + truncamento
func safeInt(v interface{}) (int, error)

// safeString — extrai string de interface{}
func safeString(v interface{}) (string, error)
```

### Padrão de uso

Todas as leituras de argumentos JSON usam as funções safe:

```go
// ✅ CORRETO — retorna erro legível em vez de panic
a0, aErr := safeInt(args[0])
if aErr != nil {
    sendErr(fmt.Sprintf("type error: arg 0: %v", aErr))
    continue
}
a := C.int(a0)

// ❌ INCORRETO — panic se args[0] não for float64
a := C.int(int(args[0].(float64)))
```

### Resultado

Se o client enviar um tipo errado (ex.: string onde espera-se number), o helper retorna:
```json
{"ok": false, "error": "type error: arg 0: expected number, got string"}
```

O helper **não crasha** e continua processando os próximos requests normalmente.

### Timeout em callbacks

A função `call_cb_from_go` (chamada por código C para invocar callbacks do Fig) possui timeout de 5 segundos. Se o runtime não responder nesse intervalo, a função retorna `nil` em vez de bloquear indefinidamente.

---

## Metadados de tipos (`arg_types`)

Algumas assinaturas usam `arg_types` para permitir dispatch correto no helper, principalmente em chamadas com tipos mistos.

### Exemplo (sym + call)

```json
{
    "cmd": "sym",
    "id": 12,
    "handle": "lib-1",
    "name": "db_bind_text",
    "rtype": "int",
    "arg_types": ["int", "int", "string", "int"]
}
```

O `arg_types` é reenviado no `call` para permitir o dispatch correto:

```json
{
    "cmd": "call",
    "id": 13,
    "symbol": "sym-9",
    "args": [1, 1, "bob@fig.dev", -1],
    "arg_types": ["int", "int", "string", "int"]
}
```

Sem `arg_types`, o helper usa um fallback que assume argumentos homogêneos.

---

## Identificadores (handles e symbols)

Todos os identificadores trocados entre o client e o helper são **strings opacas**. Nenhum dos lados deve interpretar o conteúdo — apenas armazená-los e reenviá-los.

### Formato atual

| Tipo | Formato | Exemplo |
|------|---------|---------|
| Handle (biblioteca) | `lib-N` | `"lib-1"`, `"lib-42"` |
| Symbol (função) | `sym-N` | `"sym-1"`, `"sym-200"` |
| Memória | `m-N` | `"m-1"`, `"m-99"` |

### Por que strings e não números?

Antes, IDs eram `int64` no helper e passavam por round-trip de conversões:

1. Helper gera `int64` → envia como JSON number
2. Client lê como `float64` → converte via `fmt.Sprintf("%v", val)`
3. Client armazena como `string`
4. Client reenvia → `ParseInt` → `float64` → helper lê → `int64`

Esse pipeline falhava para valores grandes (`float64(1e15)` → `"1e+15"` → `ParseInt` quebra). Com strings, o ID é transmitido sem nenhuma conversão.

### Regras

1. **Helper**: gera IDs com `fmt.Sprintf("lib-%d", counter)` / `fmt.Sprintf("sym-%d", counter)`.
2. **Client**: recebe `result.handle` e `result.symbol` como strings diretas — sem `ParseInt`, sem `Sprintf`.
3. **Client → Helper**: envia IDs como strings no JSON (`"handle": "lib-1"`, `"symbol": "sym-5"`).
4. **Helper**: lê IDs com `req["handle"].(string)` / `req["symbol"].(string)` — sem `.(float64)`.
5. IDs são **imutáveis** após criação e **únicos** dentro da sessão do helper.

---

## Logging e diagnóstico

O helper e o client usam um sistema de log por níveis, controlado via:

- **Flag**: `--log-level <nível>` (apenas helper)
- **Variável de ambiente**: `FFI_LOG_LEVEL=<nível>` (helper e client)
- **Precedência**: flag > env var > padrão (`warn`)

### Níveis disponíveis

| Nível | Quando usar | Exemplos |
|-------|------------|----------|
| `error` | Falhas irrecuperáveis | encode error, malloc failed |
| `warn` | Situações anormais mas não fatais | invalid json, callback timeout, crashing |
| `info` | Eventos de ciclo de vida | server started, socket listening |
| `debug` | Detalhes de cada request/response | req ping id=1, cb response, mem_write |

### Padrão em produção

O nível padrão é `warn`. Isso significa que:
- Requests normais (ping, load, sym, call) **não geram nenhuma saída** no stderr.
- Somente erros e avisos (timeout, json inválido) aparecem.

### Modo debug

Para diagnóstico, usar `FFI_LOG_LEVEL=debug`:

```bash
FFI_LOG_LEVEL=debug fig run meu_programa.fig
```

Nesse modo, cada request e response é logado no stderr, permitindo inspecionar o protocolo completo.

### No client (builtins/ffi_client.go)

O client também respeita `FFI_LOG_LEVEL`:
- `debug`: loga callbacks recebidos, resultados de callbacks
- `warn`: loga erros de callback, mensagens desconhecidas

---

## Handshake e versionamento

Ao estabelecer uma nova conexão com o helper, o client envia um comando `handshake` para verificar a compatibilidade de protocolo.

### Request do client

```json
{"cmd": "handshake", "version": "1.0", "id": 1}
```

### Response do helper

```json
{
  "ok": true,
  "result": {
    "version": "1.0",
    "supported_ops": ["ping", "handshake", "load", "sym", "call", "alloc", "free", "strdup", "mem_write", "mem_read", "sleep", "crash"]
  },
  "id": 1
}
```

### Comportamento de compatibilidade

| Cenário | Resultado |
|---------|-----------|
| Versões iguais (ex: client=1.0, helper=1.0) | Sucesso silencioso |
| Mesma versão major, minor diferente (1.0 vs 1.1) | Warning, continua operação |
| Versão major diferente (1.0 vs 2.0) | Erro fatal, conexão rejeitada |
| Helper legado (sem suporte a handshake) | Warning, continua em modo compatibilidade |

### Constante de versão

```go
// No helper (tools/ffi-helper/main.go)
const FFIProtocolVersion = "1.0"

// No client (builtins/ffi_client.go)
const FFIProtocolVersion = "1.0"
```

A versão segue semver simplificado: `major.minor`. Mudanças no major indicam incompatibilidade.

