# Códigos de Erro FFI

O helper FFI retorna erros estruturados no formato:

```json
{
  "ok": false,
  "error": {
    "code": "ERR_DLOPEN_FAILED",
    "message": "dlopen: libfoo.so: cannot open shared object file"
  },
  "id": 42
}
```

No lado Go (client), esses erros são parseados para o tipo `*FFIError`:

```go
type FFIError struct {
    Code    string // ex: "ERR_DLOPEN_FAILED"
    Message string // mensagem descritiva
}
```

O método `Error()` formata como `[ERR_DLOPEN_FAILED] dlopen: ...`.

## Tabela de códigos

| Código | Quando ocorre |
|--------|---------------|
| `ERR_INVALID_JSON` | JSON malformado recebido pelo helper |
| `ERR_UNKNOWN_CMD` | Comando não reconhecido |
| `ERR_MISSING_PARAM` | Parâmetro obrigatório ausente (ex: path, name, mem_id) |
| `ERR_INVALID_HANDLE` | Handle de biblioteca inválido ou não encontrado |
| `ERR_INVALID_SYMBOL` | ID de símbolo inválido ou não encontrado |
| `ERR_DLOPEN_FAILED` | Falha ao abrir a biblioteca compartilhada (dlopen) |
| `ERR_DLSYM_FAILED` | Falha ao resolver o símbolo na biblioteca (dlsym) |
| `ERR_CALL_FAILED` | Erro durante a execução da chamada FFI |
| `ERR_TYPE_ERROR` | Tipo de argumento incompatível com a assinatura |
| `ERR_UNSUPPORTED_ARGS` | Número de argumentos não suportado |
| `ERR_MALLOC_FAILED` | Falha ao alocar memória no helper (malloc) |
| `ERR_INVALID_MEM_ID` | ID de memória inválido ou não encontrado |
| `ERR_OUT_OF_BOUNDS` | Acesso de memória fora dos limites alocados |
| `ERR_INVALID_BASE64` | Dados base64 inválidos em mem_write |

## Retrocompatibilidade

O parser de erros do client (`ffiParseError`) é retrocompatível:

- **Formato novo**: `{"code": "ERR_...", "message": "..."}` → retorna `*FFIError`
- **Formato antigo**: string simples → retorna `fmt.Errorf("prefix: %v", raw)`

## Uso em testes

O tipo `FFIError` é exportado e pode ser verificado com `errors.As`:

```go
var ffiErr *builtins.FFIError
if errors.As(err, &ffiErr) {
    switch ffiErr.Code {
    case "ERR_DLOPEN_FAILED":
        // tratar
    case "ERR_INVALID_HANDLE":
        // tratar
    }
}
```

## Acesso direto ao client (testes)

Para testes que precisam acessar o client diretamente:

```go
hc, err := builtins.GetHelperForTest(projectDir, helperBinPath)
_, err = hc.Load("/invalid.so")
// err é *FFIError com Code="ERR_DLOPEN_FAILED"
```
