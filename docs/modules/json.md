# Módulo json

```js
use "json"
```

Serialização e parsing de JSON.

## json.parse(texto)

Converte uma string JSON em um valor Fig (número, string, booleano, null, array ou objeto):

```js
let obj = json.parse('{"nome": "Carlos", "idade": 30}')
print(obj.nome)   # Carlos
print(obj.idade)  # 30

let arr = json.parse('[1, 2, 3]')
print(arr)  # [1, 2, 3]

let num = json.parse('42')
print(num)  # 42
```

Se o JSON for inválido, gera um erro:

```js
let dados = try json.parse("não é JSON") onerror(e) {
    print("Erro: " + e)
    return null
}
```

## json.stringify(valor)

Converte um valor Fig em uma string JSON compacta (sem formatação):

```js
let obj = {nome: "Fig", versao: 1}
let texto = json.stringify(obj)
print(texto)  # {"nome":"Fig","versao":1}
```

```js
let arr = [1, true, "abc", null]
print(json.stringify(arr))  # [1,true,"abc",null]
```

## json.serialize(valor)

Converte um valor Fig em uma string JSON **formatada** (com indentação):

```js
let obj = {nome: "Fig", versao: 1}
let texto = json.serialize(obj)
print(texto)
```

Saída:

```json
{
  "nome": "Fig",
  "versao": 1
}
```

## json.deserialize(texto)

Alias para `json.parse` — funciona de forma idêntica:

```js
let dados = json.deserialize('{"ativo": true}')
print(dados.ativo)  # true
```

## Exemplo completo

```js
use "json"
use "io"

# Criando dados
let config = {
    host: "localhost",
    port: 5432,
    debug: false,
    tags: ["dev", "test"]
}

# Salvando como JSON
let texto = json.serialize(config)
io.writeFile("config.json", texto)

# Lendo de volta
let lido = io.readFile("config.json")
let dados = json.parse(lido)
print(dados.host)  # localhost
print(dados.port)  # 5432

# Limpando
io.deleteFile("config.json")
```

## Referência rápida

| Função                | Descrição                         |
|----------------------|-----------------------------------|
| `json.parse(s)`       | JSON string → valor Fig           |
| `json.stringify(v)`   | Valor Fig → JSON compacto         |
| `json.serialize(v)`   | Valor Fig → JSON formatado        |
| `json.deserialize(s)` | Alias de `json.parse`             |
