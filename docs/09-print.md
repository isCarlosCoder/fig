# print

## Saída no terminal

A função `print` exibe um valor no terminal:

```js
print("Hello, World!")
print(42)
print(true)
print(null)
print([1, 2, 3])
print({nome: "Fig"})
```

## Sintaxe

```
print(expressão)
print(expressão1, expressão2, ...)
```

- Aceita **qualquer tipo** de valor
- Aceita **múltiplos argumentos** separados por vírgula — são impressos na mesma linha, separados por espaço
- Sempre imprime **uma linha** (automaticamente adiciona quebra de linha)
- Retorna `null`

## Múltiplos argumentos

Quando você passa vários valores separados por vírgula, eles são impressos na mesma linha com espaço entre eles:

```js
print(1, 2, 3, 4, 5)         # 1 2 3 4 5
print("nome:", "Fig")         # nome: Fig
print("x =", 10, "y =", 20)  # x = 10 y = 20
```

Isso é mais prático do que concatenar com `+`:

```js
let nome = "Carlos"
let idade = 30

# Com múltiplos argumentos (mais simples)
print("Nome:", nome, "Idade:", idade)  # Nome: Carlos Idade: 30

# Com concatenação (mais verboso)
print("Nome: " + nome + " Idade: " + idade)  # Nome: Carlos Idade: 30
```

## Representação dos tipos

| Tipo       | Exemplo de saída                |
|------------|----------------------------------|
| Number     | `42`, `3.14`                     |
| String     | `Hello` (sem aspas)              |
| Boolean    | `true`, `false`                  |
| Null       | `null`                           |
| Array      | `[1, 2, 3]`                     |
| Object     | `{nome: Fig, idade: 30}`        |
| Function   | `<fn nome>`                     |
| Struct      | `<struct Nome>`                 |
| Instance   | `<instance of Nome>`            |

## Concatenação dentro do print

O `print` aceita qualquer expressão, incluindo concatenações:

```js
let nome = "Fig"
let versao = 1

print("Linguagem: " + nome)           # Linguagem: Fig
print("Versão: " + versao)            # Versão: 1
print("Resultado: " + (2 + 3))        # Resultado: 5
```

## Múltiplos prints

Cada chamada imprime em uma nova linha:

```js
print("Linha 1")
print("Linha 2")
print("Linha 3")
```

Saída:

```
Linha 1
Linha 2
Linha 3
```

> **Nota:** Não existe `println` ou `printf` — `print` é a única função de saída embutida diretamente na linguagem. Para entrada do usuário, use o módulo `io` com `io.input()` ou `io.readLine()`.
