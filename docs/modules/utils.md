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
