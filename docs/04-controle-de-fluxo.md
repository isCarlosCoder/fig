# Controle de Fluxo

## if / elif / else

A estrutura condicional básica:

```js
let idade = 18

if (idade >= 18) {
    print("Maior de idade")
} elif (idade >= 12) {
    print("Adolescente")
} else {
    print("Criança")
}
```

### Regras

- A **condição** deve estar entre parênteses `()`
- O **corpo** deve estar entre chaves `{ }`
- `elif` e `else` são **opcionais**
- Pode ter quantos `elif` quiser

### Condições aninhadas

```js
let x = 10

if (x > 0) {
    if (x > 5) {
        print("Maior que 5")
    } else {
        print("Entre 1 e 5")
    }
}
```

### Escopo dentro de blocos

Variáveis declaradas dentro de um bloco `if` não existem fora dele:

```js
if (true) {
    let dentro = "só aqui"
    print(dentro)  # funciona
}
# print(dentro)  # erro: variável não encontrada
```

---

## while

Repete enquanto a condição for verdadeira:

```js
let i = 0
while (i < 5) {
    print(i)
    i++
}
# Saída: 0, 1, 2, 3, 4
```

---

## do-while

Executa o bloco **pelo menos uma vez**, depois verifica a condição:

```js
let i = 0
do {
    print(i)
    i++
} while (i < 3)
# Saída: 0, 1, 2
```

Mesmo se a condição for falsa desde o início, o bloco executa uma vez:

```js
let x = 100
do {
    print("Executou!")  # imprime uma vez
} while (x < 0)
```

---

## for (loop clássico)

O loop `for` tem três partes separadas por `;`:

```
for (inicialização; condição; passo) { corpo }
```

```js
for (let i = 0; i < 5; i++) {
    print(i)
}
# Saída: 0, 1, 2, 3, 4
```

Todas as partes são opcionais:

```js
let i = 0
for (; i < 3;) {
    print(i)
    i++
}
```

---

## for-in

Itera sobre os elementos de um **array**:

```js
let frutas = ["maçã", "banana", "uva"]
for fruta in frutas {
    print(fruta)
}
# Saída: maçã, banana, uva
```

Funciona com qualquer expressão que retorne um array:

```js
for item in [1, 2, 3] {
    print(item * 10)
}
# Saída: 10, 20, 30
```

---

## range

Gera uma sequência de números para iterar:

```
range(início, fim)
range(início, fim, passo)
```

O valor `fim` **não é incluído** (exclusivo).

```js
# De 0 até 4
for i in range(0, 5) {
    print(i)
}
# Saída: 0, 1, 2, 3, 4

# De 0 até 8, de 2 em 2
for i in range(0, 10, 2) {
    print(i)
}
# Saída: 0, 2, 4, 6, 8

# Contagem regressiva
for i in range(5, 0, -1) {
    print(i)
}
# Saída: 5, 4, 3, 2, 1
```

---

## enumerate

Itera sobre um array com **índice e valor**:

```js
let cores = ["vermelho", "verde", "azul"]
for i, cor in enumerate(cores) {
    print(i + ": " + cor)
}
# Saída:
# 0: vermelho
# 1: verde
# 2: azul
```

---

## break

Interrompe o loop atual:

```js
for i in range(0, 100) {
    if (i == 5) {
        break
    }
    print(i)
}
# Saída: 0, 1, 2, 3, 4
```

---

## continue

Pula para a próxima iteração do loop:

```js
for i in range(0, 10) {
    if (i % 2 == 0) {
        continue  # pula os pares
    }
    print(i)
}
# Saída: 1, 3, 5, 7, 9
```

---

## match

O `match` é uma estrutura de controle que compara um valor com vários padrões e executa o código correspondente ao primeiro padrão que coincidir. É uma alternativa mais limpa a cadeias de `if/elif`.

### Sintaxe básica

```js
let x = 2
match x {
    1 => { print("um") }
    2 => { print("dois") }
    3 => { print("três") }
    _ => { print("outro") }  # padrão coringa
}
# Saída: dois
```

### Padrão coringa `_`

Use `_` para capturar qualquer valor que não corresponda a nenhum padrão anterior (equivalente ao `else`):

```js
let cor = "roxo"
match cor {
    "vermelho" => { print("Red") }
    "azul" => { print("Blue") }
    _ => { print("Outra cor") }
}
# Saída: Outra cor
```

### Match como expressão (com retorno)

O `match` pode ser usado como expressão para atribuir o resultado a uma variável. Neste caso, os braços usam valores inline ao invés de blocos:

```js
let valor = 1
let resultado = match valor {
    1 => "um"
    2 => "dois"
    _ => "outro"
}
print(resultado)  # "um"
```

### Múltiplos valores por braço

É possível listar vários valores separados por vírgula em um mesmo braço:

```js
let n = 5
match n {
    1, 3, 5 => { print("ímpar") }
    2, 4, 6 => { print("par") }
    _ => { print("outro") }
}
# Saída: ímpar
```

### Match com expressões

Tanto o valor a ser comparado quanto os padrões podem ser expressões:

```js
let a = 3
let b = 2
match a + b {
    4 => { print("quatro") }
    5 => { print("cinco") }
    6 => { print("seis") }
}
# Saída: cinco
```

### Match com tipos

Combine com `types.type()` para verificar tipos:

```js
use "types"

let x = 42
match types.type(x) {
    "number" => { print("é número") }
    "string" => { print("é string") }
    _ => { print("outro tipo") }
}
# Saída: é número
```

### Match em funções

O `match` pode ser usado com `return` em funções:

```js
fn classificar(n) {
    return match n {
        1 => "um"
        2 => "dois"
        _ => "muitos"
    }
}

print(classificar(1))   # "um"
print(classificar(99))  # "muitos"
```

### Regras

- **O primeiro padrão que coincidir é executado** (não há fall-through).
- **É obrigatório ter um braço padrão** usando o `_` (coringa). Se não houver `_`, o `match` gera um **erro em tempo de execução**.
- **Padrões duplicados não são permitidos.** Se dois padrões avaliarem para o mesmo valor, o `match` falha com erro em tempo de execução.
- **O coringa `_` deve ficar sozinho no seu braço** — não é permitido escrever `_ , 1` no mesmo padrão.
- **Os padrões são avaliados antes da correspondência** e duplicatas são detectadas com base nos valores avaliados (portanto, *efeitos colaterais dos padrões ocorrem durante essa avaliação*).
- Os braços podem usar blocos `{ }` ou expressões inline.
- O `_` (coringa) deve ser o último braço.

---

## Resumo

| Estrutura    | Uso                                          |
|-------------|----------------------------------------------|
| `if/elif/else` | Execução condicional                      |
| `match`     | Compara valor com padrões (alternativa a if/elif) |
| `while`     | Repete enquanto condição for verdadeira       |
| `do-while`  | Executa pelo menos uma vez, depois repete     |
| `for`       | Loop clássico com inicialização/condição/passo|
| `for-in`    | Itera sobre elementos de um array            |
| `range`     | Gera sequência numérica para loop             |
| `enumerate` | Itera com índice e valor                      |
| `break`     | Sai do loop                                   |
| `continue`  | Pula para próxima iteração                    |
