# Operadores

## Operadores aritméticos

| Operador | Descrição       | Exemplo       | Resultado |
|----------|-----------------|---------------|-----------|
| `+`      | Adição          | `10 + 3`      | `13`      |
| `-`      | Subtração       | `10 - 3`      | `7`       |
| `*`      | Multiplicação   | `10 * 3`      | `30`      |
| `/`      | Divisão         | `10 / 3`      | `3.333...`|
| `%`      | Módulo (resto)  | `10 % 3`      | `1`       |

### Concatenação com `+`

Quando **pelo menos um** dos operandos é uma string, o `+` concatena:

```js
print("Olá" + " " + "mundo")   # Olá mundo
print("Valor: " + 42)          # Valor: 42
print("Ativo: " + true)        # Ativo: true
```

> **Nota:** Operações como `-`, `*`, `/`, `%` com strings causam **erro de runtime**.

## Operadores de comparação

| Operador | Descrição        | Exemplo    | Resultado |
|----------|------------------|------------|-----------|
| `==`     | Igual            | `5 == 5`   | `true`    |
| `!=`     | Diferente        | `5 != 3`   | `true`    |
| `>`      | Maior que        | `5 > 3`    | `true`    |
| `>=`     | Maior ou igual   | `5 >= 5`   | `true`    |
| `<`      | Menor que        | `3 < 5`    | `true`    |
| `<=`     | Menor ou igual   | `5 <= 5`   | `true`    |

Os operadores `>`, `>=`, `<`, `<=` só funcionam com **números**. Usar com outros tipos gera erro:

```js
print(10 > 5)        # true
# print("a" > "b")   # erro: not a number: string
```

## Operadores lógicos

| Operador | Descrição | Exemplo            | Resultado |
|----------|-----------|--------------------|-----------|
| `&&`     | E lógico  | `true && false`    | `false`   |
| `\|\|`   | OU lógico | `false \|\| true`  | `true`    |
| `!`      | Negação   | `!true`            | `false`   |

### Short-circuit (curto-circuito)

Os operadores `&&` e `||` usam avaliação em curto-circuito:

- `&&`: se o lado esquerdo for **falsy**, retorna `false` sem avaliar o lado direito
- `||`: se o lado esquerdo for **truthy**, retorna `true` sem avaliar o lado direito

```js
# O lado direito não é avaliado nestes casos:
print(false && 10/0)   # false (sem erro)
print(true || 10/0)    # true  (sem erro)
```

## Operadores unários

| Operador  | Descrição              | Exemplo     | Resultado |
|-----------|------------------------|-------------|-----------|
| `-`       | Negação numérica       | `-5`        | `-5`      |
| `!`       | Negação lógica         | `!true`     | `false`   |
| `++`      | Incremento (prefixo)   | `++x`       | `x + 1`   |
| `--`      | Decremento (prefixo)   | `--x`       | `x - 1`   |

### Incremento e decremento

Fig suporta `++` e `--` tanto como prefixo quanto como pós-fixo:

```js
let x = 5

# Pós-fixo: retorna o valor ANTES de incrementar
let a = x++
print(a)  # 5
print(x)  # 6

# Prefixo: incrementa ANTES de retornar
let b = ++x
print(b)  # 8
print(x)  # 8
```

## Precedência de operadores

Do **mais alto** (avaliado primeiro) ao **mais baixo**:

| Nível | Operadores                  | Descrição                  |
|-------|-----------------------------|----------------------------|
| 1     | `()`, `[]`, `.`             | Agrupamento, acesso        |
| 2     | `-x`, `!x`, `++x`, `--x`   | Unários                    |
| 3     | `*`, `/`, `%`               | Multiplicação, divisão     |
| 4     | `+`, `-`                    | Adição, subtração          |
| 5     | `>`, `>=`, `<`, `<=`        | Comparação                 |
| 6     | `==`, `!=`                  | Igualdade                  |
| 7     | `&&`                        | E lógico                   |
| 8     | `\|\|`                      | OU lógico                  |

Use parênteses para alterar a ordem de avaliação:

```js
print(2 + 3 * 4)     # 14 (multiplicação primeiro)
print((2 + 3) * 4)   # 20 (parênteses alteram a ordem)
```
