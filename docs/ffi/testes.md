# Cobertura de Testes FFI

## Arquivos de teste

| Arquivo | Escopo | Qty |
|---------|--------|-----|
| `tests/ffi_test.go` | Testes de integração: ping, load, sym, call, callbacks, restart | ~12 |
| `tests/ffi_three_str_test.go` | Chamadas com 3 argumentos string | 1 |
| `tests/ffi_nested_struct_test.go` | Structs aninhados via FFI | 1 |
| `tests/ffi_stress_test.go` | Stress tests: concorrência, memory leak | ~3 |
| `tests/ffi_type_safety_test.go` | Proteção contra panics em type assertions | 3 |
| `tests/ffi_string_ids_test.go` | IDs de string (lib-N, sym-N), uniqueness | 2 |
| `tests/ffi_log_levels_test.go` | Sistema de log levels (warn, debug, env var) | 3 |
| `tests/ffi_coercion_test.go` | Coerção de tipos (int truncation, string→int, num→string) | 4 |
| `tests/ffi_error_codes_test.go` | Códigos de erro estruturados (FFIError) | 5 |
| `tests/ffi_handshake_test.go` | Handshake de versão do protocolo | 2 |
| `tests/ffi_protocol_test.go` | Testes de protocolo: edge cases, campos ausentes, burst | 7+1skip |

## Cenários cobertos pelo protocolo

| Cenário | Código esperado | Arquivo |
|---------|----------------|---------|
| Load com path vazio | `ERR_MISSING_PARAM` | `ffi_protocol_test.go` |
| Load com path inexistente | `ERR_DLOPEN_FAILED` | `ffi_protocol_test.go`, `ffi_error_codes_test.go` |
| Sym com handle inválido | `ERR_INVALID_HANDLE` | `ffi_protocol_test.go`, `ffi_error_codes_test.go` |
| Sym com nome inexistente | `ERR_DLSYM_FAILED` | `ffi_protocol_test.go`, `ffi_error_codes_test.go` |
| Call com symbol inválido | `ERR_INVALID_SYMBOL` | `ffi_protocol_test.go` |
| Free com mem_id inválido | `ERR_INVALID_MEM_ID` | `ffi_protocol_test.go`, `ffi_error_codes_test.go` |
| 100 requests em rajada | `ERR_DLOPEN_FAILED` x 100 | `ffi_protocol_test.go` |
| Formato FFIError [CODE] msg | — | `ffi_error_codes_test.go` |

## Cenários pendentes

- JSON malformado enviado diretamente ao helper (requer raw socket/pipe)
- Request sem campo `cmd` (requer raw socket/pipe)
- `mem_write` com offset fora dos limites
- `mem_read` com offset+len fora dos limites
- `mem_write` com base64 inválido
- `alloc` com size 0 ou negativo

> Esses cenários exigem acesso raw ao pipe do helper (enviar bytes diretamente), o que não é possível via `HelperForTest`. Podem ser adicionados futuramente com um helper de teste raw.
