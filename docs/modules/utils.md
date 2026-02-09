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
