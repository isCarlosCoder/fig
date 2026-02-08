# Módulo math

```js
use "math"
```

Funções matemáticas e constantes numéricas.

## Constantes

| Constante  | Valor                    | Descrição              |
|-----------|--------------------------|------------------------|
| `math.PI`  | `3.141592653589793`      | Pi (π)                 |
| `math.E`   | `2.718281828459045`      | Número de Euler (e)    |
| `math.INF` | `+Inf`                   | Infinito positivo      |

```js
print(math.PI)   # 3.141592653589793
print(math.E)    # 2.718281828459045
print(math.INF)  # +Inf
```

## Funções básicas

### math.abs(x)

Retorna o valor absoluto de `x`:

```js
print(math.abs(-5))   # 5
print(math.abs(3))    # 3
```

### math.pow(base, expoente)

Retorna `base` elevado a `expoente`:

```js
print(math.pow(2, 10))  # 1024
print(math.pow(3, 2))   # 9
```

### math.sqrt(x)

Retorna a raiz quadrada de `x`:

```js
print(math.sqrt(16))   # 4
print(math.sqrt(2))    # 1.4142135623730951
```

### math.cbrt(x)

Retorna a raiz cúbica de `x`:

```js
print(math.cbrt(27))  # 3
print(math.cbrt(8))   # 2
```

## Arredondamento

### math.floor(x)

Arredonda para baixo:

```js
print(math.floor(3.7))   # 3
print(math.floor(-2.3))  # -3
```

### math.ceil(x)

Arredonda para cima:

```js
print(math.ceil(3.2))   # 4
print(math.ceil(-2.7))  # -2
```

### math.round(x)

Arredonda para o inteiro mais próximo:

```js
print(math.round(3.5))   # 4
print(math.round(3.4))   # 3
```

## Mínimo, máximo e clamp

### math.min(a, b)

Retorna o menor entre `a` e `b`:

```js
print(math.min(3, 7))   # 3
print(math.min(-1, 1))  # -1
```

### math.max(a, b)

Retorna o maior entre `a` e `b`:

```js
print(math.max(3, 7))   # 7
print(math.max(-1, 1))  # 1
```

### math.clamp(x, min, max)

Limita `x` ao intervalo `[min, max]`:

```js
print(math.clamp(5, 0, 10))    # 5
print(math.clamp(-3, 0, 10))   # 0
print(math.clamp(15, 0, 10))   # 10
```

## Números aleatórios

### math.rand()

Retorna um número decimal aleatório entre 0 (inclusive) e 1 (exclusive):

```js
let r = math.rand()
print(r)  # ex: 0.7234...
```

### math.randInt(min, max)

Retorna um inteiro aleatório entre `min` (inclusive) e `max` (exclusive):

```js
let dado = math.randInt(1, 7)  # 1 até 6
print(dado)
```

## Trigonometria

> Todas as funções trigonométricas usam **radianos**.

### math.sin(x), math.cos(x), math.tan(x)

```js
print(math.sin(0))            # 0
print(math.cos(0))            # 1
print(math.sin(math.PI / 2))  # 1
print(math.tan(math.PI / 4))  # 1 (aprox.)
```

## Logaritmos e exponenciação

### math.log(x)

Logaritmo natural (base e):

```js
print(math.log(math.E))  # 1
print(math.log(1))       # 0
```

### math.log10(x)

Logaritmo base 10:

```js
print(math.log10(100))   # 2
print(math.log10(1000))  # 3
```

### math.exp(x)

Retorna `e^x`:

```js
print(math.exp(1))  # 2.718281828459045
print(math.exp(0))  # 1
```

## Referência rápida

| Função               | Descrição                            |
|---------------------|--------------------------------------|
| `math.abs(x)`        | Valor absoluto                       |
| `math.pow(a, b)`     | Potência                             |
| `math.sqrt(x)`       | Raiz quadrada                        |
| `math.cbrt(x)`       | Raiz cúbica                          |
| `math.floor(x)`      | Arredondar para baixo                |
| `math.ceil(x)`       | Arredondar para cima                 |
| `math.round(x)`      | Arredondar para mais próximo         |
| `math.min(a, b)`     | Menor valor                          |
| `math.max(a, b)`     | Maior valor                          |
| `math.clamp(x, m, M)`| Limitar ao intervalo                 |
| `math.rand()`        | Aleatório [0, 1)                     |
| `math.randInt(a, b)` | Inteiro aleatório [a, b)             |
| `math.sin(x)`        | Seno                                 |
| `math.cos(x)`        | Cosseno                              |
| `math.tan(x)`        | Tangente                             |
| `math.log(x)`        | Logaritmo natural                    |
| `math.log10(x)`      | Logaritmo base 10                    |
| `math.exp(x)`        | Exponencial (e^x)                    |
