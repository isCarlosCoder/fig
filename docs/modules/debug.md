# Módulo debug

```js
use "debug"
```

Ferramentas para depuração e diagnóstico.

## debug.dump(x)

Retorna uma representação detalhada e formatada de qualquer valor, incluindo tipo e estrutura:

```js
let pessoa = {nome: "Carlos", idade: 30, hobbies: ["música", "código"]}
print(debug.dump(pessoa))
```

Saída (exemplo):

```
(object) {
  nome: (string) "Carlos"
  idade: (number) 30
  hobbies: (array) [
    (string) "música"
    (string) "código"
  ]
}
```

Funciona com todos os tipos:

```js
print(debug.dump(42))          # (number) 42
print(debug.dump("hello"))     # (string) "hello"
print(debug.dump(null))        # (nil) null
print(debug.dump([1, [2, 3]])) # mostra estrutura aninhada
```

## debug.inspect(x)

Retorna uma representação compacta `<tipo: valor>`:

```js
print(debug.inspect(42))          # <number: 42>
print(debug.inspect("hello"))     # <string: hello>
print(debug.inspect(true))        # <boolean: true>
print(debug.inspect(null))        # <nil: null>
print(debug.inspect([1, 2]))      # <array: [1, 2]>
```

## debug.type(x)

Retorna o nome do tipo como string (funciona como `types.type`):

```js
print(debug.type(42))         # number
print(debug.type("hi"))       # string
print(debug.type([]))         # array
print(debug.type({a: 1}))     # object
```

## debug.assert(condição, mensagem?)

Verifica se a condição é verdadeira. Se for falsa, gera um erro com a mensagem:

```js
let x = 5
debug.assert(x > 0, "x deve ser positivo")           # ok, nada acontece
debug.assert(x > 10, "x deveria ser maior que 10")   # ERRO!
```

Mensagem padrão (quando não especificada):

```js
debug.assert(false)  # erro: "assertion failed"
```

Útil para testes e validações:

```js
fn dividir(a, b) {
    debug.assert(b != 0, "Divisor não pode ser zero")
    return a / b
}
```

## debug.panic(mensagem?)

Para a execução imediatamente com um erro:

```js
debug.panic("Algo deu muito errado!")
# Programa termina com a mensagem de erro
```

Mensagem padrão:

```js
debug.panic()  # erro: "panic!"
```

## Referência rápida

| Função                      | Descrição                              |
|----------------------------|----------------------------------------|
| `debug.dump(x)`            | Representação detalhada (com tipos)    |
| `debug.inspect(x)`         | Representação compacta `<tipo: valor>` |
| `debug.type(x)`            | Nome do tipo                           |
| `debug.assert(cond, msg?)` | Verifica condição, erro se falsa       |
| `debug.panic(msg?)`        | Para execução com erro                 |
