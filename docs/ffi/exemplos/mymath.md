# Exemplo: Biblioteca Matemática (mymath)

Este exemplo mostra o fluxo completo para criar e usar uma biblioteca C com funções matemáticas via FFI no Fig.

## 1. Código C

Crie o arquivo `mymath.c`:

```c
#include <math.h>

// Soma dois inteiros
int add(int a, int b) {
    return a + b;
}

// Multiplica um double por 2
double mul2(double x) {
    return x * 2.0;
}

// Soma três inteiros
int sum3(int a, int b, int c) {
    return a + b + c;
}

// Calcula potência (wrapper de pow)
double power(double base, double exp) {
    return pow(base, exp);
}
```

## 2. Compilação

```bash
# Linux
gcc -shared -fPIC -o libmymath.so mymath.c -lm

# macOS
gcc -shared -fPIC -o libmymath.dylib mymath.c -lm
```

## 3. Configuração (fig.toml)

```toml
[ffi]
enabled = true
helper = "./ffi-helper"
```

## 4. Código Fig

```js
use "ffi"

// Carregar a biblioteca
var lib = ffi.load("./libmymath.so")

// Obter símbolos com tipos declarados
var add = ffi.sym(lib, "add", "int", ["int", "int"])
var mul2 = ffi.sym(lib, "mul2", "double", ["double"])
var sum3 = ffi.sym(lib, "sum3", "int", ["int", "int", "int"])
var power = ffi.sym(lib, "power", "double", ["double", "double"])

// Chamar funções
print(ffi.call(add, 10, 20))       // 30
print(ffi.call(mul2, 3.14))        // 6.28
print(ffi.call(sum3, 1, 2, 3))     // 6
print(ffi.call(power, 2.0, 10.0))  // 1024.0
```

## 5. Executar

```bash
# Compilar o helper (se ainda não foi feito)
cd tools/ffi-helper && go build -o ../../ffi-helper . && cd ../..

# Rodar o programa
fig run programa.fig
```

## Erros comuns

| Sintoma | Causa provável | Solução |
|---------|---------------|---------|
| `ERR_DLOPEN_FAILED` | Biblioteca não encontrada | Verifique o caminho em `ffi.load()` |
| `ERR_DLSYM_FAILED` | Função não existe na lib | Verifique o nome do símbolo com `nm -D libmymath.so` |
| `ERR_TYPE_ERROR` | Tipo do argumento incompatível | Verifique `argTypes` em `ffi.sym()` |

## Ver também

- [Tipos suportados](../tipos.md)
- [Erros](../erros.md)
- [Protocolo](../protocolo.md)
