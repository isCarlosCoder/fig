# FFI — Callbacks

> **Status: Experimental** — O sistema de callbacks bidirecionais está em desenvolvimento. Esta documentação descreve o design planejado.

## Visão geral

Callbacks permitem que código C chame funções Fig e vice-versa. Isso é útil para APIs que esperam ponteiros de função (qsort, event handlers, etc.).

## Arquitetura

```
┌──────────┐  register_callback  ┌────────────┐
│  Fig     │ ──────────────────→ │  Helper    │
│ runtime  │                     │            │
│          │ ← callback invoke ─ │  C code    │
│          │ ── result ────────→ │            │
└──────────┘                     └────────────┘
```

## Registro de callback (Fig → Helper)

```js
use "ffi"

// Registrar uma função Fig como callback
var cb_id = ffi.register_callback(func(x) {
    return x * 2
})

// Passar o callback para uma função C
var lib = ffi.load("./libmylib.so")
var apply = ffi.sym(lib, "apply_func", "int", ["callback", "int"])
var resultado = ffi.call(apply, cb_id, 21)
print(resultado)  // 42

// Desregistrar quando não precisar mais
ffi.unregister_callback(cb_id)
```

## Timeout de callbacks

Callbacks têm um timeout de 5 segundos por padrão para evitar deadlocks:

```toml
[ffi]
callback_timeout = 5000  # ms
```

Se o callback não retornar dentro do timeout, o helper retorna um valor padrão (0 para int, 0.0 para double, "" para string) e loga um aviso.

## Protocolo de callback

### Registro (Fig → Helper)

```json
{"cmd": "register_callback", "id": 1, "retType": "int", "argTypes": ["int"]}
```

Resposta:
```json
{"ok": true, "result": "cb-1", "id": 1}
```

### Invocação (Helper → Fig)

Quando o código C chama o callback, o helper envia uma mensagem para o runtime:

```json
{"cmd": "callback_invoke", "cb_id": "cb-1", "args": [21]}
```

O runtime executa a função Fig e responde:

```json
{"ok": true, "result": 42}
```

## Limitações atuais

1. **Não reentrante** — um callback não pode invocar outro callback
2. **Tipos simples** — callbacks só suportam tipos primitivos (int, double, string)
3. **Um argumento** — callbacks atualmente suportam no máximo 4 argumentos
4. **Timeout fixo** — sem override per-callback (apenas global via fig.toml)

## Ver também

- [Protocolo](protocolo.md)
- [Tipos suportados](tipos.md)
- [Helper](helper.md)
