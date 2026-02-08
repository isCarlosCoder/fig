# Módulo strings

```js
use "strings"
```

Manipulação de strings (textos).

## strings.len(s)

Retorna o comprimento da string em caracteres (Unicode):

```js
print(strings.len("Fig"))     # 3
print(strings.len("olá"))     # 3
print(strings.len(""))        # 0
```

## strings.upper(s)

Converte para maiúsculas:

```js
print(strings.upper("olá mundo"))  # OLÁ MUNDO
```

## strings.lower(s)

Converte para minúsculas:

```js
print(strings.lower("FIG LANG"))  # fig lang
```

## strings.trim(s)

Remove espaços em branco do início e do fim:

```js
print(strings.trim("  hello  "))  # "hello"
```

## strings.split(s, separador)

Divide a string pelo separador e retorna um array:

```js
let partes = strings.split("a,b,c", ",")
print(partes)  # [a, b, c]

let palavras = strings.split("olá mundo fig", " ")
print(palavras)  # [olá, mundo, fig]
```

## strings.join(array, separador)

Junta os elementos de um array em uma string:

```js
let arr = ["Fig", "é", "legal"]
print(strings.join(arr, " "))   # Fig é legal
print(strings.join(arr, "-"))   # Fig-é-legal
```

## strings.replace(s, antigo, novo)

Substitui **todas** as ocorrências de `antigo` por `novo`:

```js
print(strings.replace("aabbcc", "b", "X"))  # aaXXcc
print(strings.replace("hello", "l", "r"))   # herro
```

## strings.contains(s, sub)

Retorna `true` se a string contém a substring:

```js
print(strings.contains("FigLang", "Lang"))  # true
print(strings.contains("FigLang", "xyz"))   # false
```

## strings.startsWith(s, prefixo)

Retorna `true` se a string começa com o prefixo:

```js
print(strings.startsWith("FigLang", "Fig"))    # true
print(strings.startsWith("FigLang", "Lang"))   # false
```

## strings.endsWith(s, sufixo)

Retorna `true` se a string termina com o sufixo:

```js
print(strings.endsWith("FigLang", "Lang"))  # true
print(strings.endsWith("FigLang", "Fig"))   # false
```

## strings.indexOf(s, sub)

Retorna o índice da **primeira** ocorrência da substring, ou `-1` se não encontrar:

```js
print(strings.indexOf("abcabc", "bc"))  # 1
print(strings.indexOf("hello", "xyz"))  # -1
```

## strings.lastIndexOf(s, sub)

Retorna o índice da **última** ocorrência da substring, ou `-1`:

```js
print(strings.lastIndexOf("abcabc", "bc"))  # 4
print(strings.lastIndexOf("hello", "xyz"))  # -1
```

## strings.substring(s, início, fim)

Retorna uma substring do índice `início` até `fim` (exclusivo):

```js
print(strings.substring("FigLang", 0, 3))  # Fig
print(strings.substring("FigLang", 3, 7))  # Lang
```

## strings.charAt(s, índice)

Retorna o caractere no índice especificado:

```js
print(strings.charAt("Fig", 0))  # F
print(strings.charAt("Fig", 2))  # g
```

## strings.repeat(s, n)

Repete a string `n` vezes:

```js
print(strings.repeat("ab", 3))   # ababab
print(strings.repeat("-", 10))   # ----------
```

## Referência rápida

| Função                             | Descrição                        |
|-----------------------------------|----------------------------------|
| `strings.len(s)`                   | Comprimento da string            |
| `strings.upper(s)`                 | Converter para maiúsculas        |
| `strings.lower(s)`                 | Converter para minúsculas        |
| `strings.trim(s)`                  | Remover espaços das bordas       |
| `strings.split(s, sep)`            | Dividir em array                 |
| `strings.join(arr, sep)`           | Juntar array em string           |
| `strings.replace(s, old, new)`     | Substituir todas ocorrências     |
| `strings.contains(s, sub)`         | Verificar se contém              |
| `strings.startsWith(s, prefix)`    | Verificar início                 |
| `strings.endsWith(s, suffix)`      | Verificar final                  |
| `strings.indexOf(s, sub)`          | Índice da primeira ocorrência    |
| `strings.lastIndexOf(s, sub)`      | Índice da última ocorrência      |
| `strings.substring(s, start, end)` | Extrair substring                |
| `strings.charAt(s, i)`             | Caractere no índice              |
| `strings.repeat(s, n)`             | Repetir n vezes                  |
