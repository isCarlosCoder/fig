# Funções

## Declarando funções

Use `fn` para declarar uma função:

```js
fn saudacao() {
    print("Olá, mundo!")
}

saudacao()  # Olá, mundo!
```

## Parâmetros

Funções podem receber parâmetros:

```js
fn soma(a, b) {
    return a + b
}

let resultado = soma(3, 5)
print(resultado)  # 8
```

## Retorno

Use `return` para retornar um valor. Se não houver `return`, a função retorna `null`:

```js
fn dobro(x) {
    return x * 2
}

fn semRetorno() {
    let x = 10
    # nenhum return explícito
}

print(dobro(5))       # 10
print(semRetorno())   # null
```

`return` sem valor também retorna `null`:

```js
fn verificar(x) {
    if (x < 0) {
        return  # retorna null
    }
    return x * 2
}
```

## Funções anônimas

Funções podem ser criadas sem nome e atribuídas a variáveis:

```js
let dobro = fn(x) {
    return x * 2
}

print(dobro(5))  # 10
```

Funções anônimas podem ser passadas como argumento:

```js
fn aplicar(funcao, valor) {
    return funcao(valor)
}

let resultado = aplicar(fn(x) { return x * x }, 4)
print(resultado)  # 16
```

## Funções como valores

Em Fig, funções são **valores de primeira classe** — podem ser atribuídas a variáveis, passadas como argumento e retornadas de outras funções:

```js
fn criar_multiplicador(fator) {
    return fn(x) {
        return x * fator
    }
}

let triplicar = criar_multiplicador(3)
print(triplicar(5))   # 15
print(triplicar(10))  # 30
```

## Closures (fechamentos)

Funções capturam variáveis do escopo onde foram criadas:

```js
fn contador() {
    let n = 0
    return fn() {
        n = n + 1
        return n
    }
}

let cont = contador()
print(cont())  # 1
print(cont())  # 2
print(cont())  # 3
```

A variável `n` continua existindo enquanto a função retornada existir — isso é uma **closure**.

## Recursão

Funções podem chamar a si mesmas:

```js
fn fatorial(n) {
    if (n <= 1) {
        return 1
    }
    return n * fatorial(n - 1)
}

print(fatorial(5))  # 120
```

```js
fn fibonacci(n) {
    if (n <= 1) {
        return n
    }
    return fibonacci(n - 1) + fibonacci(n - 2)
}

print(fibonacci(10))  # 55
```

## Escopo de funções

Cada chamada de função cria um **novo escopo**. Variáveis declaradas dentro de uma função não existem fora:

```js
let x = "global"

fn teste() {
    let x = "local"
    print(x)  # local
}

teste()
print(x)  # global
```

Funções podem acessar variáveis de escopos externos (mas não o contrário):

```js
let mensagem = "Olá"

fn imprimir() {
    print(mensagem)  # Olá (acessa escopo externo)
}

imprimir()
```

## Resumo

| Conceito                | Sintaxe                                       |
|------------------------|-----------------------------------------------|
| Declaração             | `fn nome(params) { corpo }`                   |
| Chamada                | `nome(args)`                                   |
| Retorno                | `return valor`                                 |
| Função anônima         | `fn(params) { corpo }`                        |
| Closure                | Função que captura variáveis externas          |
| Primeira classe        | Funções como valores, argumentos e retornos    |
| Recursão               | Função que chama a si mesma                    |
