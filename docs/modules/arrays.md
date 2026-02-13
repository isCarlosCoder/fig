# Módulo arrays

```js
use "arrays"
```

Operações para manipulação de arrays (listas).

## Adicionando elementos

### arrays.push(arr, valor)

Adiciona um elemento ao **final** do array:

```js
let frutas = ["maçã", "banana"]
arrays.push(frutas, "uva")
print(frutas)  # [maçã, banana, uva]
```

### arrays.unshift(arr, valor)

Adiciona um elemento ao **início** do array:

```js
let nums = [2, 3]
arrays.unshift(nums, 1)
print(nums)  # [1, 2, 3]
```

### arrays.insert(arr, índice, valor)

Insere um elemento na posição especificada:

```js
let letras = ["a", "c", "d"]
arrays.insert(letras, 1, "b")
print(letras)  # [a, b, c, d]
```

## Removendo elementos

### arrays.pop(arr)

Remove e retorna o **último** elemento:

```js
let nums = [1, 2, 3]
let ultimo = arrays.pop(nums)
print(ultimo)  # 3
print(nums)    # [1, 2]
```

### arrays.shift(arr)

Remove e retorna o **primeiro** elemento:

```js
let nums = [1, 2, 3]
let primeiro = arrays.shift(nums)
print(primeiro)  # 1
print(nums)      # [2, 3]
```

### arrays.remove(arr, índice)

Remove e retorna o elemento no índice especificado:

```js
let cores = ["vermelho", "verde", "azul"]
let removido = arrays.remove(cores, 1)
print(removido)  # verde
print(cores)     # [vermelho, azul]
```

## Consultando

### arrays.len(arr)

Retorna o comprimento do array:

```js
print(arrays.len([1, 2, 3]))  # 3
print(arrays.len([]))          # 0
```

### arrays.index(arr, valor)

Retorna o índice do valor, ou `-1` se não encontrar:

```js
let nums = [10, 20, 30]
print(arrays.index(nums, 20))   # 1
print(arrays.index(nums, 99))   # -1
```

### Indexação direta com índices negativos

Arrays suportam indexação direta com inteiros negativos (comportamento estilo Python):
`arr[-1]` retorna o último elemento, `arr[-2]` o penúltimo, e assim por diante. Índices fora do intervalo continuam gerando erro.

```js
let a = ["x", "y", "z"]
print(a[-1])   # z
print(a[-3])   # x
```

### arrays.contains(arr, valor)

Retorna `true` se o array contém o valor:

```js
let frutas = ["maçã", "banana"]
print(arrays.contains(frutas, "banana"))  # true
print(arrays.contains(frutas, "uva"))     # false
```

## Transformando

### arrays.slice(arr, início, fim)

Retorna um **novo** sub-array de `início` até `fim` (exclusivo):

```js
let nums = [0, 1, 2, 3, 4, 5]
let parte = arrays.slice(nums, 1, 4)
print(parte)  # [1, 2, 3]
print(nums)   # [0, 1, 2, 3, 4, 5] (original intacto)
```

### arrays.concat(a, b)

Retorna um **novo** array combinando dois arrays:

```js
let a = [1, 2]
let b = [3, 4]
let c = arrays.concat(a, b)
print(c)  # [1, 2, 3, 4]
```

### arrays.reverse(arr)

Inverte o array **in-place** (modifica o original):

```js
let nums = [1, 2, 3]
arrays.reverse(nums)
print(nums)  # [3, 2, 1]
```

### arrays.sort(arr)

Ordena o array **in-place** — números em ordem numérica, outros tipos por representação string:

```js
let nums = [3, 1, 4, 1, 5]
arrays.sort(nums)
print(nums)  # [1, 1, 3, 4, 5]

let nomes = ["carlos", "ana", "bruno"]
arrays.sort(nomes)
print(nomes)  # [ana, bruno, carlos]
```

### arrays.unique(arr)

Retorna um **novo** array sem duplicatas:

```js
let nums = [1, 2, 2, 3, 3, 3]
let unicos = arrays.unique(nums)
print(unicos)  # [1, 2, 3]
```

### arrays.shuffle(arr)

Embaralha o array **in-place** (aleatoriamente):

```js
let cartas = [1, 2, 3, 4, 5]
arrays.shuffle(cartas)
print(cartas)  # ex: [3, 1, 5, 2, 4]
```

## Funções de ordem superior

### arrays.map(arr, fn)

Retorna um **novo** array com `fn` aplicada a cada elemento:

```js
let nums = [1, 2, 3, 4, 5]
let dobro = arrays.map(nums, fn(x) { return x * 2 })
print(dobro)  # [2, 4, 6, 8, 10]
```

### arrays.filter(arr, fn)

Retorna um **novo** array com os elementos para os quais `fn` retorna `true`:

```js
let nums = [1, 2, 3, 4, 5]
let pares = arrays.filter(nums, fn(x) { return x % 2 == 0 })
print(pares)  # [2, 4]
```

### arrays.reduce(arr, fn, valorInicial)

Reduz o array a um único valor, chamando `fn(acumulador, elemento)` para cada item:

```js
let nums = [1, 2, 3, 4, 5]
let soma = arrays.reduce(nums, fn(acc, x) { return acc + x }, 0)
print(soma)  # 15
```

### arrays.find(arr, fn)

Retorna o **primeiro** elemento para o qual `fn` retorna `true`, ou `null` se nenhum satisfizer:

```js
let nums = [1, 2, 3, 4, 5]
let achado = arrays.find(nums, fn(x) { return x > 3 })
print(achado)  # 4

let nenhum = arrays.find(nums, fn(x) { return x > 100 })
print(nenhum)  # null
```

### Encadeando funções

```js
# map + filter: triplicar e manter só > 6
let result = arrays.filter(
    arrays.map([1, 2, 3, 4, 5], fn(x) { return x * 3 }),
    fn(x) { return x > 6 }
)
print(result)  # [9, 12, 15]
```

## Referência rápida

| Função                          | Descrição                              | Modifica original? |
|--------------------------------|----------------------------------------|--------------------|
| `arrays.push(arr, v)`          | Adiciona ao final                      | Sim                |
| `arrays.pop(arr)`              | Remove do final                        | Sim                |
| `arrays.shift(arr)`            | Remove do início                       | Sim                |
| `arrays.unshift(arr, v)`       | Adiciona ao início                     | Sim                |
| `arrays.insert(arr, i, v)`     | Insere na posição                      | Sim                |
| `arrays.remove(arr, i)`        | Remove da posição                      | Sim                |
| `arrays.slice(arr, start, end)`| Sub-array                              | Não                |
| `arrays.concat(a, b)`          | Combinar arrays                        | Não                |
| `arrays.reverse(arr)`          | Inverter                               | Sim                |
| `arrays.sort(arr)`             | Ordenar                                | Sim                |
| `arrays.index(arr, v)`         | Índice do valor                        | Não                |
| `arrays.contains(arr, v)`      | Verificar se contém                    | Não                |
| `arrays.unique(arr)`           | Remover duplicatas                     | Não                |
| `arrays.shuffle(arr)`          | Embaralhar                             | Sim                |
| `arrays.map(arr, fn)`          | Mapear cada elemento                   | Não                |
| `arrays.filter(arr, fn)`       | Filtrar elementos                      | Não                |
| `arrays.reduce(arr, fn, init)` | Reduzir array a um valor               | Não                |
| `arrays.find(arr, fn)`         | Encontrar primeiro que satisfaça `fn`  | Não                |
| `arrays.len(arr)`              | Comprimento                            | Não                |
