# Núcleo

```js
http.request(method, url, body, headers)
```

Retorna:

```js
{
  status: number,
  headers: object,
  body: string
}
```

Com isso você já cobre:

- GET / POST / PUT / DELETE
- headers
- auth
- JSON
- upload simples

# Atalhos (só conveniência)
```js
http.get(url)
http.post(url, body)
```

internamente chamam http.request.

# Utilitários essenciais
```js
http.download(url, path)
http.setTimeout(ms)
http.setHeader(key, value)
http.clearHeaders()
```

# Helpers de resposta
```js
http.isOk(res)
http.raiseForStatus(res)
```

# Uso real
```js
let res = http.request("GET", "https://api.site.com/data", nil, {
    "Authorization": "Bearer token"
})

if (http.isOk(res)) {
    let obj = json.parse(res.body)
    print(obj)
}
```

# Download:
```js
http.download("https://site.com/file.zip", "file.zip")
```

# POST JSON:
```js
http.setHeader("Content-Type", "application/json")
http.post("https://api.site.com", json.stringify(data))
```

# API final mínima
```js
http.request(method, url, body, headers)
http.get(url)
http.post(url, body)

http.download(url, path)

http.setTimeout(ms)
http.setHeader(key, value)
http.clearHeaders()

http.isOk(res)
http.raiseForStatus(res)
```