# FFI — Coerção e Validação de Tipos

> Regras de coerção automática aplicadas pelo runtime Fig antes de enviar argumentos ao helper.

---

## Quando a coerção é aplicada

A coerção ocorre **somente quando `argTypes` foi declarado** na chamada `sym()`. Sem `argTypes`, os valores são passados diretamente sem nenhuma conversão.

```fig
// Com argTypes — coerção ativa
var sym = ffi.sym(handle, "sum3", "int", ["int", "int", "int"])
ffi.call(sym, 10.7, 20.9, 30.3)  // → sum3(10, 20, 30) = 60

// Sem argTypes — pass-through
var sym2 = ffi.sym(handle, "sum3", "int")
ffi.call(sym2, 10.7, 20.9, 30.3)  // → sum3(10.7, 20.9, 30.3) — helper decide
```

---

## Regras de coerção

| argType declarado | Tipo do valor Fig | Ação | Resultado |
|---|---|---|---|
| `"int"` | `Number` | Truncamento (`int(n)`) | `42.7` → `42` |
| `"int"` | `String` | `strconv.Atoi` | `"42"` → `42` |
| `"int"` | `String` (inválido) | **Erro** | `"abc"` → erro `"expects int, got string"` |
| `"double"` | `Number` | Pass-through | `3.14` → `3.14` |
| `"double"` | `String` | `ParseFloat` | `"3.14"` → `3.14` |
| `"string"` | `Number` | `fmt.Sprintf("%g", n)` | `42` → `"42"` |
| `"string"` | `String` | Pass-through | `"hello"` → `"hello"` |

### Validação de range

Quando argType é `"int"`, valores que excedem `math.MaxInt32` (≈2.1 bilhões) geram erro para evitar overflow silencioso no lado C.

---

## Mensagens de erro

Erros de coerção seguem o padrão:

```
call: arg N expects TYPE, got string "VALUE" (not a valid integer)
```

Exemplos:
- `call: arg 0 expects int, got string "abc" (not a valid integer)`
- `call: arg 1 value 3e+15 overflows int32 range`

---

## Sem argTypes

Quando `sym()` é chamado sem o 4º argumento (argTypes), **nenhuma coerção é feita**. Os valores são enviados ao helper exatamente como recebidos do runtime Fig:

- Números → `float64` no JSON
- Strings → string no JSON
- Booleanos → boolean no JSON

O helper é responsável por interpretar os tipos corretamente.
