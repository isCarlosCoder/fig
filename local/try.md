# Try

Do mais simples ao mais complexo, usando:

```ini
x = try <expr> onerror(e) {...}
```

# 1. Caso trivial

x recebe 0 pelo return

```js
let x = try 10 / 0 onerror() {
    return 0
}
```

# 2. Capturando mensagem

```js
let x = try toInt("abc") onerror(e) {
    print("Error: ", e)
    return -1
}
```

# 3. Encadeado com função
```js
fn safeDiv(a, b) {
    return try a / b onerror(e) {
        return 0
    }
}

print(safeDiv(10, 0)) # 0
print(safeDiv(10, 2)) # 5
```

# 4. Array com elemento protegido

```js
let arr = [1, 2, try risky() onerror { return 0 }, 4]
```

# 5. Pipeline lógico
```js
let data = try http.get(url) onerror {
    return {"status": 0, "body": ""}
}
```

# 6. Fallback em cascata
```js
let cfg = try io.readFile("config.json") onerror {
    return try io.readFile("default.json") onerror {
        return "{}"
    }
}
```

# 7. Função de alto nivel
```js
fn loadUser(id) {
    let res = try http.get("/user/" + id) onerror(e) {
        print("http error: " + e)
        return null
    }

    return json.parse(res.body)
}
```

# 8. Em função funcional
```js
let values = array.map(inputs, fn(x) {
    return try toInt(x) onerror {
        return 0
    }
})
```

# 9. Condicional com try
```js
if (try check() onerror { return false }) {
    print("ok")
}
```

# 10. Struct + try
```js
struct User {
    name
    age

    fn init(data) {
        this.name = try data.name onerror { return "anon" }
        this.age  = try data.age  onerror { return 0 }
    }
}
```

# 11. Loop resiliente
```js
for x in inputs {
    let n = try parseInt(x) onerror {
        continue
    }
    print(n * 2)
}

```

# 12. Com funções como valor
```js
fn safe(f) {
    return fn(x) {
        return try f(x) onerror { return nil }
    }
}

let safeParse = safe(parseInt)
print(safeParse("123"))
print(safeParse("abc"))
```