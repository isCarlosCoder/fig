# Módulo regex

```js
use "regex"
```

Expressões regulares (usa a sintaxe RE2 do Go).

## regex.match(texto, padrão)

Retorna `true` se o padrão encontra correspondência no texto:

```js
print(regex.match("hello123", "[0-9]+"))   # true
print(regex.match("hello", "[0-9]+"))      # false
print(regex.match("abc@def.com", ".+@.+")) # true
```

## regex.findAll(texto, padrão)

Retorna um array com todas as correspondências encontradas:

```js
let nums = regex.findAll("abc 123 def 456 ghi", "[0-9]+")
print(nums)  # [123, 456]

let palavras = regex.findAll("Olá Mundo Fig", "[A-Z][a-záéíóúã]+")
print(palavras)  # [Olá, Mundo, Fig]
```

Se nenhuma correspondência for encontrada, retorna um array vazio:

```js
let resultado = regex.findAll("abc", "[0-9]+")
print(resultado)  # []
```

## regex.replaceRegex(texto, padrão, substituição)

Substitui **todas** as correspondências do padrão pela string de substituição:

```js
let resultado = regex.replaceRegex("abc 123 def 456", "[0-9]+", "***")
print(resultado)  # abc *** def ***
```

```js
# Removendo pontuação
let limpo = regex.replaceRegex("Olá, mundo! Tudo bem?", "[^a-zA-ZáéíóúÁÉÍÓÚãõ ]", "")
print(limpo)  # Olá mundo Tudo bem
```

## regex.splitRegex(texto, padrão)

Divide o texto usando o padrão como separador:

```js
let partes = regex.splitRegex("um1dois2três3quatro", "[0-9]")
print(partes)  # [um, dois, três, quatro]
```

```js
# Dividir por múltiplos espaços
let palavras = regex.splitRegex("a   b  c    d", "\\s+")
print(palavras)  # [a, b, c, d]
```

## Sintaxe de padrões

Fig usa a sintaxe **RE2** (a mesma do Go). Os padrões mais comuns:

| Padrão      | Descrição                          |
|------------|-------------------------------------|
| `.`         | Qualquer caractere                  |
| `[abc]`     | Um dos caracteres a, b ou c        |
| `[a-z]`     | Qualquer letra minúscula           |
| `[0-9]`     | Qualquer dígito                    |
| `[^abc]`    | Qualquer caractere exceto a, b, c  |
| `\d`        | Dígito (equivale a `[0-9]`)       |
| `\w`        | Letra, dígito ou `_`              |
| `\s`        | Espaço em branco                   |
| `+`         | Um ou mais                         |
| `*`         | Zero ou mais                       |
| `?`         | Zero ou um                         |
| `{n}`       | Exatamente n vezes                 |
| `{n,m}`     | Entre n e m vezes                  |
| `^`         | Início da string                   |
| `$`         | Fim da string                      |
| `(grupo)`   | Grupo de captura                   |
| `a\|b`       | a OU b                             |

> **Nota:** Como strings em Fig usam `\` para escapes, use `\\` para representar `\` em padrões regex: `"\\d+"` para representar `\d+`.

## Exemplo: Validação de email

```js
use "regex"

fn validarEmail(email) {
    return regex.match(email, "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
}

print(validarEmail("user@example.com"))    # true
print(validarEmail("sem-arroba.com"))      # false
```

## Referência rápida

| Função                              | Descrição                            |
|------------------------------------|--------------------------------------|
| `regex.match(s, pattern)`           | Testar se padrão casa com texto      |
| `regex.findAll(s, pattern)`         | Encontrar todas as correspondências  |
| `regex.replaceRegex(s, pat, repl)`  | Substituir correspondências          |
| `regex.splitRegex(s, pattern)`      | Dividir texto pelo padrão            |
