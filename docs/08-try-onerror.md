# Tratamento de Erros — try / onerror

Fig usa a expressão `try/onerror` para capturar erros em tempo de execução e fornecer valores alternativos ou executar ações de recuperação.

## Sintaxe básica

```
try <expressão> onerror {
    # bloco executado se a expressão causar erro
}
```

Se a expressão não causar erro, seu valor é retornado normalmente. Se causar erro, o bloco `onerror` é executado.

## Exemplo simples

```js
let resultado = try 10 / 0 onerror {
    return -1
}
print(resultado)  # -1
```

Sem `try`, a divisão por zero causaria um erro fatal. Com `try`, o erro é capturado e o bloco `onerror` fornece um valor alternativo via `return`.

## Try com bloco guardado

A parte protegida do `try` pode ser uma **expressão** ou um **bloco**. Usar um bloco permite executar código mais complexo (como `return`, `break`, `continue`) dentro da seção protegida.

Sintaxe:

```js
try {
    // código que pode usar return, break, continue
} onerror(e) {
    // tratamento
}
```

Comportamento:

- Se o bloco guardado executar sem erro:
  - Se fizer `return <valor>`, esse valor é o resultado do `try`.
  - Se o bloco cair fora sem `return`, o resultado do `try` é `null`.
  - `break` e `continue` dentro do bloco afetam o loop que contém o `try` (se aplicável).
- Se o bloco causar erro, o `onerror` é executado como no caso da expressão.

Exemplos:

```js
let x = try {
    return 7
} onerror {
    return 0
}
print(x)  # 7
```

```js
let x = try {
    let a = 1
} onerror {
    return 0
}
print(x)  # null
```

```js
use "arrays"
let res = []
for v in [1,2,3] {
    let n = try {
        if v == 1 { continue }
        return v
    } onerror { }
    arrays.push(res, n)
}
print(res)  # [2, 3]


## Capturando a mensagem de erro

Use `onerror(variável)` para receber a mensagem de erro:

```js
let resultado = try 10 / 0 onerror(erro) {
    print("Erro capturado: " + erro)
    return 0
}
```

Saída:

```
Erro capturado: division by zero
```

## Retornando valor de fallback

O `return` dentro do bloco `onerror` define o valor que a expressão `try` retorna quando há erro:

```js
use "json"

let dados = try json.parse("JSON inválido!") onerror {
    return {}  # retorna objeto vazio como fallback
}

print(dados)  # {}
```

## Sem return — resultado é null

Se o bloco `onerror` não tiver `return`, a expressão `try` retorna `null`:

```js
let x = try 10 / 0 onerror {
    print("Aconteceu um erro")
}
print(x)  # null
```

## try dentro de condicionais

```js
let valor = try minhaFuncaoArriscada() onerror(e) {
    print("Falhou: " + e)
    return null
}

if (valor != null) {
    print("Sucesso: " + valor)
} else {
    print("Usando valor padrão")
}
```

## try com onerror() — parênteses vazios

Você pode usar `onerror()` sem nomear a variável de erro:

```js
let resultado = try operacaoPerigosa() onerror() {
    return "fallback"
}
```

## try dentro de loops

O bloco `onerror` pode usar `break` e `continue` para controlar loops:

```js
let valores = [10, 0, 5, 0, 2]

for v in valores {
    let resultado = try 100 / v onerror {
        continue  # pula divisões por zero
    }
    print(resultado)
}
# Saída: 10, 20, 50
```

```js
let dados = ["abc", "123", "xyz"]

for d in dados {
    let num = try types.toInt(d) onerror {
        break  # para no primeiro erro
    }
    print(num)
}
```

## Executando ações sem retorno

O bloco `onerror` pode apenas logar ou executar efeitos sem retornar:

```js
try operacaoQuePodemFalhar() onerror(e) {
    print("Aviso: " + e)
    # sem return — resultado do try será null
}
```

## Exemplos práticos

### Leitura segura de arquivo

```js
use "io"

let conteudo = try io.readFile("config.txt") onerror(e) {
    print("Não foi possível ler config: " + e)
    return "configuração padrão"
}
```

### Parse de JSON seguro

```js
use "json"

fn parsearJSON(texto) {
    return try json.parse(texto) onerror(e) {
        print("JSON inválido: " + e)
        return null
    }
}

let dados = parsearJSON('{"nome": "Fig"}')
print(dados)  # {nome: Fig}

let invalido = parsearJSON("não é JSON")
print(invalido)  # null
```

### Acesso seguro a propriedades

```js
fn obterValor(obj, campo) {
    return try obj[campo] onerror {
        return null
    }
}
```

## Resumo

| Formato                            | Descrição                                      |
|------------------------------------|-------------------------------------------------|
| `try expr onerror { ... }`         | Captura erro, sem acesso à mensagem             |
| `try expr onerror(e) { ... }`      | Captura erro com mensagem em `e`                |
| `try expr onerror() { ... }`       | Captura erro, parênteses vazios (sem variável)  |
| `return valor` no bloco            | Define o valor de fallback                       |
| Sem `return`                       | Expressão `try` retorna `null`                   |
| `break` / `continue` no bloco     | Controla o loop que contém o `try`               |
