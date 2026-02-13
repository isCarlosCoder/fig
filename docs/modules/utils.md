# utils module

Utilities with small helpers that don't fit other modules.

## ord(char)

Returns the Unicode code point (number) of a single-character string.

Example:

```js
use "utils";
print(utils.ord("A"));  # 65
print(utils.ord("é"));  # 233
```

Errors:
- `ord()` expects exactly one argument, which must be a single-character string.

## chr(code)

Returns the character string for a Unicode code point integer.

Example:

```js
use "utils";
print(utils.chr(65));   # A
print(utils.chr(233));  # é
```

Errors:
- `chr()` expects one integer argument in range 0..0x10FFFF.

## fromCodePoint(code)

Alias for `chr()` (same behavior).

## codePointAt(s, index)

Returns the Unicode code point (number) of the rune at `index` (0-based) in string `s`.

Example:

```js
use "utils";
print(utils.codePointAt("A", 0));   # 65
print(utils.codePointAt("😀", 0));  # 128512
```

Errors:
- index must be an integer and within range.

## codePoints(s)

Returns an array of code points (numbers) for each rune in `s`.

Example:

```js
use "utils";
print(utils.codePoints("A😀"));  # [65, 128512]
```

## normalize(form, s)

Normalize string `s` to the form `NFC`, `NFD`, `NFKC`, or `NFKD` using Unicode normalization.

Example:

```js
use "utils";
let comb = "e" + utils.fromCodePoint(769);
print(utils.normalize("NFC", comb));  # é
```

Errors:
- `form` must be one of `NFC`, `NFD`, `NFKC`, `NFKD`.

## isLetter(s) / isDigit(s)

Check whether the first rune of `s` is a Unicode letter or digit, respectively.

Example:

```js
use "utils";
print(utils.isLetter("é"));  # true
print(utils.isDigit("3"));   # true
```

## isWhitespace(s) / isUpper(s) / isLower(s) / isAlphaNum(s)

Unicode-aware helpers that inspect the *first rune* of a string.

- `isWhitespace(s)` — returns `true` for Unicode space characters (includes NBSP and other space separators).
- `isUpper(s)` / `isLower(s)` — check case of the first rune.
- `isAlphaNum(s)` — true if first rune is a letter or a digit.

Example:

```js
use "utils";
print(utils.isWhitespace(" "));
print(utils.isWhitespace(utils.fromCodePoint(160)));  # NBSP
print(utils.isUpper("A"));
print(utils.isLower("a"));
print(utils.isAlphaNum("3"));
```

## fromCodePoints(arr) / toCodePoints(s)

Batch conversions between arrays of Unicode code points and strings.

```js
use "utils";
print(utils.fromCodePoints([72,101,108,108,111])); # Hello
print(utils.toCodePoints("Hi"));                 # [72, 105]
```

## runeCount(s) / byteLength(s)

- `runeCount(s)` — number of runes (user-perceived code points) in `s`.
- `byteLength(s)` — number of bytes in UTF-8 encoding of `s`.

Example:

```js
use "utils";
print(utils.runeCount("A😀"));  # 2
print(utils.byteLength("A😀")); # 5
```

## zip(...arrays)

Itera múltiplos arrays em paralelo, agrupando os elementos de cada índice em uma tupla (array). A iteração para no menor array (comportamento igual ao `zip` do Python).

- Assinatura: `zip(a1, a2, ...)` — recebe dois ou mais arrays e retorna um `Array` de `Array` (cada elemento é uma tupla com os valores correspondentes por índice).
- Corta na menor sequência — o tamanho do resultado é igual ao menor comprimento entre os arrays de entrada.

Exemplos:

```js
use "utils";
let a = [1,2,3]
let b = ["x","y","z"]
print(utils.zip(a, b))      # [[1, "x"], [2, "y"], [3, "z"]]

let a2 = [1,2]
let b2 = [10,20,30]
print(utils.zip(a2, b2))    # [[1,10], [2,20]]  (corta no menor array)
```

Erros:
- Todos os argumentos devem ser arrays — caso contrário, `zip()` lança um erro.

Resumo (trechos relevantes):

| Função                         | Descrição                                           |
|-------------------------------:|:---------------------------------------------------|
| `ord(char)`                    | Retorna o code point Unicode do caractere `char`.   |
| `chr(code)`                    | Retorna o caractere para o code point `code`.       |
| `fromCodePoint(code)`          | Alias de `chr(code)`.                               |
| `codePointAt(s, index)`        | Code point da rune em `index` na string `s`.        |
| `codePoints(s)`                | Array de code points (runes) de `s`.                |
| `fromCodePoints(arr)`          | Constrói string a partir de `arr` de code points.   |
| `toCodePoints(s)`              | Converte `s` em um array de code points.            |
| `normalize(form, s)`           | Normalização Unicode (`NFC`, `NFD`, `NFKC`, `NFKD`).|
| `isLetter(s)`                  | True se a primeira rune de `s` for letra.           |
| `isDigit(s)`                   | True se a primeira rune de `s` for dígito.          |
| `isWhitespace(s)`              | True se a primeira rune for espaço Unicode.         |
| `isUpper(s)` / `isLower(s)`    | Teste de caixa (primeira rune).                     |
| `isAlphaNum(s)`                | True se a primeira rune for letra ou dígito.        |
| `runeCount(s)`                 | Número de runes (code points) em `s`.               |
| `byteLength(s)`                | Número de bytes UTF-8 em `s`.                       |
| `zip(...arrays)`               | Itera arrays em paralelo e agrupa por índice.       |
