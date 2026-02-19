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

### Parâmetros opcionais e valores padrão ⚙️

- `name = expr` — parâmetro com **valor padrão**; a expressão é avaliada no momento da chamada (call-time).
- `name?` — parâmetro **opcional** que recebe `null` quando omitido.

Exemplos:

```js
fn add(a, b = 2) { return a + b }
print(add(3))      # 5
print(add(3, 4))   # 7

fn greet(name?) {
    if (name == null) { print("hi") } else { print("hi " + name) }
}

greet()            # hi
greet("Joao")     # hi Joao

fn f(a, b = a + 1) { print(b) }  # default avaliado no momento da chamada
f(5)                # 6
```

> Observação: chamadas continuam posicionais — não há suporte a argumentos nomeados nesta etapa.

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

## Funções nativas (@native)

Use `@native` para marcar funções simples que devem ser executadas pelo "backend" da linguagem (fast‑path numérica). A forma curta `@native` e a forma com parênteses `@native()` são aceitas; a versão com colchetes `@[native()]` também é permitida.

Regras iniciais:

- Apenas expressões numéricas simples são suportadas (literais, parâmetros, `+ - * / %`, `-x`, parênteses)
- Chamadas a `math.<fn>` estão permitidas (ex.: `math.exp`)
- Sem variáveis locais, loops, closures ou arrays no corpo
- Corpo deve ser uma única instrução `return <expr>`
- Se a anotação não puder ser compilada, a definição falha (erro em tempo de definição)

Exemplo:

```js
@native fn sigmoid(x) {
    return 1 / (1 + math.exp(-x))
}

print(sigmoid(0))  # 0.5
```

Use `@native()` com argumentos no futuro (ex.: `@native(fallback=true)`) — argumentos são aceitos e armazenados como metadados.

> Para documentação completa e exemplos avançados veja: [Funções nativas — @native](docs/14-native.md)

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
