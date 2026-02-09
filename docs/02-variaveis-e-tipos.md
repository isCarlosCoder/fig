# Variáveis e Tipos

## Declarando variáveis

Use a palavra-chave `let` para declarar uma variável:

```js
let nome = "Carlos"
let idade = 30
let ativo = true
```

Você também pode declarar sem valor inicial — a variável começa como `null`:

```js
let resultado
print(resultado)  # null
```

## Atribuição

Após declarar, use `=` para alterar o valor:

```js
let x = 10
x = 20
print(x)  # 20
```

> **Importante:** Você só pode atribuir a variáveis que já foram declaradas com `let`. Usar uma variável não declarada causa erro.

## Tipos primitivos

Fig possui **6 tipos primitivos**:

### Number (número)

Números inteiros e decimais são o mesmo tipo:

```js
let inteiro = 42
let decimal = 3.14
let negativo = -7
```

Operações aritméticas retornam números:

```js
print(10 + 3)    # 13
print(10 / 3)    # 3.3333333333333335
print(10 % 3)    # 1
```

### String (texto)

Strings são delimitadas por aspas duplas (`"`) ou simples (`'`):

```js
let nome = "Fig"
let saudacao = 'Olá, mundo!'
```

**Concatenação** com `+`:

```js
let msg = "Olá, " + "mundo!"
print(msg)  # Olá, mundo!
```

Quando um dos lados do `+` é string, o outro é automaticamente convertido:

```js
print("Idade: " + 30)     # Idade: 30
print("Ativo: " + true)   # Ativo: true
```

**Sequências de escape** suportadas:

| Sequência | Resultado       |
|-----------|-----------------|
| `\n`      | Nova linha      |
| `\t`      | Tabulação       |
| `\\`      | Barra invertida |
| `\"`      | Aspas duplas    |
| `\'`      | Aspas simples   |
| `\xNN`    | Hex (1 byte)    |
| `\ooo`    | Octal (1 byte)  |

### Boolean (booleano)

Os valores `true` e `false`:

```js
let ligado = true
let desligado = false
```

### Null

Representa a ausência de valor:

```js
let vazio = null
print(vazio)  # null
```

### Array (lista)

Arrays são listas ordenadas de valores, podendo conter tipos mistos:

```js
let numeros = [1, 2, 3, 4, 5]
let misturado = [1, "dois", true, null]
let vazio = []
```

**Acesso por índice** (começa em 0):

```js
let frutas = ["maçã", "banana", "uva"]
print(frutas[0])   # maçã
print(frutas[2])   # uva
```

**Atribuição por índice:**

```js
frutas[1] = "manga"
print(frutas)  # ["maçã", "manga", "uva"]
```

**Arrays aninhados:**

```js
let matrix = [[1, 2], [3, 4]]
print(matrix[0][1])  # 2
print(matrix[1][0])  # 3
```

### Object (objeto)

Objetos são coleções de pares chave-valor:

```js
let pessoa = {nome: "Carlos", idade: 30, ativo: true}
```

Chaves podem ser identificadores ou strings:

```js
let config = {
    "host": "localhost",
    port: 5432
}
```

**Acesso por ponto** (`.`) ou **colchetes** (`[]`):

```js
print(pessoa.nome)       # Carlos
print(pessoa["idade"])   # 30
```

**Atribuição em propriedades:**

```js
pessoa.idade = 31
pessoa.email = "carlos@fig.dev"  # adiciona nova propriedade
```

**Propriedade inexistente** retorna `null` (sem erro):

```js
print(pessoa.telefone)  # null
```

**Objetos aninhados:**

```js
let config = {
    db: {host: "localhost", port: 5432},
    debug: false
}
print(config.db.host)  # localhost
```

## Regras de truthiness (valores "verdadeiros")

Em condições (`if`, `while`, etc.), Fig avalia qualquer valor como verdadeiro ou falso:

| Valor                | Truthy/Falsy |
|----------------------|-------------|
| `true`               | truthy      |
| `false`              | **falsy**   |
| `null`               | **falsy**   |
| `0`                  | truthy      |
| `""`                 | truthy      |
| `42`, `"abc"`, etc.  | truthy      |
| `[]` (array vazio)   | **falsy**   |
| `[1, 2]`             | truthy      |
| `{}` (objeto vazio)  | **falsy**   |
| `{a: 1}`             | truthy      |

> **Nota:** Diferente de muitas linguagens, o número `0` e a string vazia `""` são **truthy** em Fig. Apenas `false`, `null`, arrays vazios e objetos vazios são falsy.

```js
if (0) { print("sim") }         # imprime "sim" (0 é truthy)
if ([]) { print("sim") }        # NÃO imprime (array vazio é falsy)
if ({a: 1}) { print("sim") }   # imprime "sim" (objeto não-vazio é truthy)
```

## Escopo de variáveis

Variáveis declaradas dentro de blocos (`{ }`) só existem naquele escopo:

```js
let a = 1
if (true) {
    let a = 2       # variável nova, só existe neste bloco
    print(a)        # 2
}
print(a)            # 1 (a original não mudou)
```

Variáveis de escopos externos são acessíveis em escopos internos:

```js
let x = 10
if (true) {
    print(x)  # 10 (acessa do escopo externo)
}
```

## Igualdade

Fig usa `==` e `!=` para comparações. A igualdade compara por **tipo e valor**:

```js
print(1 == 1)           # true
print("abc" == "abc")   # true
print(null == null)     # true
print(1 == "1")         # false (tipos diferentes)
print([1, 2] == [1, 2]) # true (compara elemento a elemento)
print({a: 1} == {a: 1}) # true (compara chave a chave)
```
