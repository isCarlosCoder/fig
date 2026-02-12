# Exemplo: Manipulação de Strings C

Este exemplo demonstra como trabalhar com strings C via FFI, incluindo alocação, duplicação e leitura de memória.

## 1. Código C

Crie o arquivo `mystrings.c`:

```c
#include <string.h>
#include <stdlib.h>
#include <ctype.h>

// Duplica uma string (retorna ponteiro alocado com malloc)
char* dupstr(const char* s) {
    return strdup(s);
}

// Concatena duas strings (caller deve liberar o resultado)
char* concat(const char* a, const char* b) {
    size_t la = strlen(a);
    size_t lb = strlen(b);
    char* result = malloc(la + lb + 1);
    if (!result) return NULL;
    memcpy(result, a, la);
    memcpy(result + la, b, lb + 1);
    return result;
}

// Retorna o comprimento de uma string
int str_len(const char* s) {
    return (int)strlen(s);
}

// Converte para maiúsculas (modifica in-place, retorna o mesmo ponteiro)
char* to_upper(char* s) {
    for (char* p = s; *p; p++) {
        *p = toupper((unsigned char)*p);
    }
    return s;
}
```

## 2. Compilação

```bash
gcc -shared -fPIC -o libmystrings.so mystrings.c
```

## 3. Código Fig

```js
use "ffi"

var lib = ffi.load("./libmystrings.so")

// === Duplicar string ===
var dupstr = ffi.sym(lib, "dupstr", "string", ["string"])
var copia = ffi.call(dupstr, "Olá, Fig!")
print(copia)  // "Olá, Fig!"

// === Concatenar ===
var concat_sym = ffi.sym(lib, "concat", "string", ["string", "string"])
var resultado = ffi.call(concat_sym, "Hello, ", "World!")
print(resultado)  // "Hello, World!"

// === Comprimento ===
var str_len = ffi.sym(lib, "str_len", "int", ["string"])
print(ffi.call(str_len, "Fig"))  // 3

// === Memória manual ===
// Para strings que precisam ser modificadas in-place:
var ptr = ffi.strdup("hello fig")
var to_upper = ffi.sym(lib, "to_upper", "string", ["string"])
var upper = ffi.call(to_upper, ptr)
print(upper)  // "HELLO FIG"
ffi.free(ptr)  // Sempre liberar memória alocada!
```

## Gerenciamento de memória

> **Importante:** Strings retornadas por funções C que alocam memória (como `strdup`, `concat`) precisam ser liberadas com `ffi.free()` quando não forem mais necessárias, para evitar vazamentos.

```js
var ptr = ffi.call(dupstr, "temporário")
// ... usar ptr ...
ffi.free(ptr)  // Liberar quando terminar
```

## Ver também

- [Tipos suportados](../tipos.md) — como strings são convertidas entre Fig e C
- [Protocolo — strdup](../protocolo.md) — detalhes do comando strdup
- [Erros](../erros.md) — códigos de erro relacionados a memória
