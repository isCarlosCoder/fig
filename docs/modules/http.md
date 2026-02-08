# Módulo http

```js
use "http"
```

Cliente e servidor HTTP.

---

## Cliente HTTP

### http.get(url)

Faz uma requisição GET e retorna um objeto de resposta:

```js
let res = http.get("https://api.example.com/dados")
print(res.status)   # 200
print(res.body)     # corpo da resposta como string
print(res.headers)  # objeto com headers da resposta
```

### http.post(url, body?)

Faz uma requisição POST:

```js
use "json"

let dados = json.stringify({nome: "Fig", tipo: "linguagem"})
let res = http.post("https://api.example.com/criar", dados)
print(res.status)
print(res.body)
```

### http.request(método, url, body?, headers?)

Requisição HTTP genérica com controle total:

```js
let res = http.request("PUT", "https://api.example.com/atualizar", corpo, {
    "Content-Type": "application/json",
    "Authorization": "Bearer token123"
})
print(res.status)
```

### http.download(url, caminho)

Baixa um arquivo e salva no caminho especificado:

```js
http.download("https://example.com/foto.jpg", "foto.jpg")
```

### http.setTimeout(ms)

Define o timeout das requisições em milissegundos:

```js
http.setTimeout(5000)  # 5 segundos de timeout
```

### http.setHeader(chave, valor)

Define um header padrão que será enviado em **todas** as requisições:

```js
http.setHeader("Authorization", "Bearer meu-token")
http.setHeader("Accept", "application/json")

# Agora todas as requisições incluem esses headers
let res = http.get("https://api.example.com/protegido")
```

### http.clearHeaders()

Remove todos os headers padrão definidos com `setHeader`:

```js
http.clearHeaders()
```

### http.isOk(res)

Retorna `true` se o status da resposta é entre 200 e 299:

```js
let res = http.get("https://api.example.com/dados")
if (http.isOk(res)) {
    print("Sucesso!")
} else {
    print("Erro: status " + res.status)
}
```

### http.raiseForStatus(res)

Gera um erro se o status da resposta for >= 400:

```js
let res = http.get("https://api.example.com/dados")
try http.raiseForStatus(res) onerror(e) {
    print("Requisição falhou: " + e)
}
```

---

## Servidor HTTP

### http.route(método, caminho, handler)

Registra uma rota no servidor HTTP. O handler recebe `req` e `res`:

```js
http.route("GET", "/", fn(req, res) {
    res.send("Olá, mundo!")
})

http.route("GET", "/sobre", fn(req, res) {
    res.json({nome: "Fig", versao: "0.1.0"})
})
```

### Objeto `req` (request)

O handler recebe informações da requisição:

| Propriedade   | Descrição                           |
|--------------|-------------------------------------|
| `req.method`  | Método HTTP (GET, POST, etc.)       |
| `req.path`    | Caminho da URL                      |
| `req.body`    | Corpo da requisição (string)        |
| `req.headers` | Headers como objeto                 |
| `req.query`   | Parâmetros da query string (objeto) |

### Objeto `res` (response)

O handler responde ao cliente com:

#### res.send(corpo)

Envia uma resposta de texto/HTML:

```js
http.route("GET", "/", fn(req, res) {
    res.send("<h1>Olá!</h1>")
})
```

#### res.json(objeto)

Envia uma resposta JSON (define `Content-Type: application/json`):

```js
http.route("GET", "/api/dados", fn(req, res) {
    res.json({mensagem: "sucesso", total: 42})
})
```

#### res.status(código)

Define o código de status HTTP:

```js
http.route("GET", "/not-found", fn(req, res) {
    res.status(404)
    res.send("Página não encontrada")
})
```

### http.render(caminho, dados?)

Lê um arquivo HTML e substitui placeholders `{{chave}}` pelos valores do objeto `dados`:

```js
# template.html contém: <h1>Olá, {{nome}}!</h1>
http.route("GET", "/", fn(req, res) {
    let html = http.render("template.html", {nome: "Carlos"})
    res.send(html)
})
```

### http.listen(porta)

Inicia o servidor HTTP na porta especificada. **Esta função bloqueia** a execução:

```js
print("Servidor rodando na porta 8080...")
http.listen(8080)
```

---

## Exemplo completo: API REST

```js
use "http"
use "json"

let tarefas = []

# Listar tarefas
http.route("GET", "/tarefas", fn(req, res) {
    res.json(tarefas)
})

# Criar tarefa
http.route("POST", "/tarefas", fn(req, res) {
    let dados = json.parse(req.body)
    arrays.push(tarefas, dados)
    res.status(201)
    res.json({mensagem: "Criado", total: arrays.len(tarefas)})
})

# Página inicial
http.route("GET", "/", fn(req, res) {
    res.send("<h1>API de Tarefas</h1><p>Use /tarefas</p>")
})

print("Servidor em http://localhost:3000")
http.listen(3000)
```

## Exemplo completo: Cliente HTTP

```js
use "http"
use "json"

# Configurar headers padrão
http.setHeader("Accept", "application/json")
http.setTimeout(10000)

# Fazer requisição
let res = http.get("https://jsonplaceholder.typicode.com/todos/1")

if (http.isOk(res)) {
    let dados = json.parse(res.body)
    print("Tarefa: " + dados.title)
    print("Concluída: " + dados.completed)
} else {
    print("Erro: " + res.status)
}
```

## Referência rápida

### Cliente

| Função                               | Descrição                          |
|-------------------------------------|------------------------------------|
| `http.get(url)`                      | Requisição GET                     |
| `http.post(url, body?)`             | Requisição POST                    |
| `http.request(method, url, body?, headers?)` | Requisição genérica       |
| `http.download(url, path)`           | Baixar arquivo                     |
| `http.setTimeout(ms)`               | Definir timeout                    |
| `http.setHeader(key, value)`         | Header padrão para todas as req.   |
| `http.clearHeaders()`               | Limpar headers padrão              |
| `http.isOk(res)`                    | Status 2xx?                        |
| `http.raiseForStatus(res)`           | Erro se status >= 400              |

### Servidor

| Função                              | Descrição                          |
|------------------------------------|------------------------------------|
| `http.route(method, path, handler)` | Registrar rota                     |
| `http.listen(port)`                 | Iniciar servidor (bloqueia)        |
| `http.render(path, data?)`          | Template HTML com substituição     |
| `res.send(body)`                    | Responder com texto/HTML           |
| `res.json(obj)`                     | Responder com JSON                 |
| `res.status(code)`                  | Definir código de status           |
