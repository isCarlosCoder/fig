# Módulo objects

```js
use "objects"
```

Operações para manipulação de objetos (pares chave-valor).

## objects.keys(obj)

Retorna um array com todas as chaves do objeto (na ordem de inserção):

```js
let pessoa = {nome: "Carlos", idade: 30}
print(objects.keys(pessoa))  # [nome, idade]
```

## objects.values(obj)

Retorna um array com todos os valores do objeto:

```js
let pessoa = {nome: "Carlos", idade: 30}
print(objects.values(pessoa))  # [Carlos, 30]
```

## objects.entries(obj)

Retorna um array de pares `[chave, valor]`:

```js
let pessoa = {nome: "Carlos", idade: 30}
let entries = objects.entries(pessoa)
print(entries)  # [[nome, Carlos], [idade, 30]]
```

Útil para iterar sobre chave e valor:

```js
for par in objects.entries(pessoa) {
    print(par[0] + ": " + par[1])
}
# nome: Carlos
# idade: 30
```

## objects.hasKey(obj, chave)

Retorna `true` se o objeto possui a chave especificada:

```js
let config = {host: "localhost", port: 5432}
print(objects.hasKey(config, "host"))    # true
print(objects.hasKey(config, "user"))    # false
```

## objects.deleteKey(obj, chave)

Remove a chave do objeto e retorna o valor removido:

```js
let obj = {a: 1, b: 2, c: 3}
let removido = objects.deleteKey(obj, "b")
print(removido)  # 2
print(obj)       # {a: 1, c: 3}
```

## objects.merge(a, b)

Retorna um **novo** objeto combinando `a` e `b`. Se houver chaves duplicadas, os valores de `b` prevalecem:

```js
let base = {cor: "azul", tamanho: "M"}
let custom = {tamanho: "G", peso: 100}
let resultado = objects.merge(base, custom)
print(resultado)  # {cor: azul, tamanho: G, peso: 100}
```

## objects.clone(obj)

Retorna uma **cópia rasa** do objeto:

```js
let original = {a: 1, b: 2}
let copia = objects.clone(original)
copia.a = 99
print(original.a)  # 1 (não afetou o original)
print(copia.a)     # 99
```

## objects.size(obj)

Retorna o número de chaves no objeto:

```js
print(objects.size({a: 1, b: 2, c: 3}))  # 3
print(objects.size({}))                    # 0
```

## objects.clear(obj)

Remove todas as chaves do objeto, retornando-o vazio:

```js
let obj = {x: 1, y: 2}
objects.clear(obj)
print(obj)             # {}
print(objects.size(obj))  # 0
```

## Referência rápida

| Função                       | Descrição                            | Modifica original? |
|-----------------------------|--------------------------------------|--------------------|
| `objects.keys(obj)`          | Array de chaves                      | Não                |
| `objects.values(obj)`        | Array de valores                     | Não                |
| `objects.entries(obj)`       | Array de `[chave, valor]`            | Não                |
| `objects.hasKey(obj, key)`   | Verificar se chave existe            | Não                |
| `objects.deleteKey(obj, key)`| Remover chave                        | Sim                |
| `objects.merge(a, b)`        | Combinar objetos                     | Não                |
| `objects.clone(obj)`         | Cópia rasa                           | Não                |
| `objects.size(obj)`          | Número de chaves                     | Não                |
| `objects.clear(obj)`         | Remover todas as chaves              | Sim                |
