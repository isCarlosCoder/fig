# Módulo types

```js
use "types"
```

Verificação e conversão de tipos.

## Verificação de tipos

### types.type(x)

Retorna o nome do tipo como string:

```js
print(types.type(42))          # number
print(types.type("hello"))     # string
print(types.type(true))        # boolean
print(types.type(null))        # nil
print(types.type([1, 2]))      # array
print(types.type({a: 1}))      # object
```

### types.isNumber(x)

```js
print(types.isNumber(42))     # true
print(types.isNumber("42"))   # false
```

### types.isString(x)

```js
print(types.isString("hi"))   # true
print(types.isString(42))     # false
```

### types.isBool(x)

```js
print(types.isBool(true))    # true
print(types.isBool(1))       # false
```

### types.isArray(x)

```js
print(types.isArray([1, 2]))  # true
print(types.isArray("abc"))   # false
```

### types.isObject(x)

```js
print(types.isObject({a: 1}))  # true
print(types.isObject([1]))     # false
```

### types.isNil(x)

```js
print(types.isNil(null))   # true
print(types.isNil(0))      # false
print(types.isNil(""))     # false
```

### types.isFunction(x)

Funciona com funções declaradas e builtins:

```js
fn soma(a, b) { return a + b }
print(types.isFunction(soma))   # true
print(types.isFunction(42))     # false
```

## Conversão de tipos

### types.toInt(x)

Converte para número inteiro (trunca decimais):

```js
print(types.toInt(3.7))      # 3
print(types.toInt("42"))     # 42
print(types.toInt(true))     # 1
print(types.toInt(false))    # 0
```

### types.toFloat(x)

Converte para número decimal:

```js
print(types.toFloat("3.14"))  # 3.14
print(types.toFloat(42))      # 42
print(types.toFloat(true))    # 1
```

### types.toString(x)

Converte qualquer valor para string:

```js
print(types.toString(42))     # 42
print(types.toString(true))   # true
print(types.toString(null))   # null
```

### types.toBool(x)

Converte para booleano usando as regras de truthiness:

```js
print(types.toBool(1))       # true
print(types.toBool(0))       # false
print(types.toBool(""))      # false
print(types.toBool(null))    # false
print(types.toBool("abc"))   # true
```

> **Nota:** As regras do `types.toBool` são do Go (0 e "" são falsy), que podem diferir das regras de truthiness do Fig. Em condições `if/while`, o Fig considera 0 e "" como **truthy**.

## Referência rápida

| Função               | Descrição                         |
|----------------------|-----------------------------------|
| `types.type(x)`      | Nome do tipo como string          |
| `types.isNumber(x)`  | É número?                         |
| `types.isString(x)`  | É string?                         |
| `types.isBool(x)`    | É booleano?                       |
| `types.isArray(x)`   | É array?                          |
| `types.isObject(x)`  | É objeto?                         |
| `types.isNil(x)`     | É null?                           |
| `types.isFunction(x)`| É função?                         |
| `types.toInt(x)`     | Converter para inteiro            |
| `types.toFloat(x)`   | Converter para decimal            |
| `types.toString(x)`  | Converter para string             |
| `types.toBool(x)`    | Converter para booleano           |
