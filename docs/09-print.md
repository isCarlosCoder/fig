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
```

- Aceita **qualquer tipo** de valor
- Sempre imprime **uma linha** (automaticamente adiciona quebra de linha)
- Retorna `null`

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
