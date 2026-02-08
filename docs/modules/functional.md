# Módulo functional

```js
use "functional"
```

Ferramentas de programação funcional para funções builtin.

> **Nota:** As funções `call`, `apply` e `partial` funcionam com funções **builtin** (dos módulos da linguagem). Para funções declaradas pelo usuário (`fn`), use chamadas diretas.

## functional.call(fn, ...args)

Chama uma função builtin com os argumentos fornecidos:

```js
use "functional"
use "math"

let resultado = functional.call(math.abs, -5)
print(resultado)  # 5
```

## functional.apply(fn, args)

Chama uma função builtin passando um array como argumentos (spread):

```js
use "functional"
use "math"

let args = [2, 10]
let resultado = functional.apply(math.pow, args)
print(resultado)  # 1024
```

## functional.partial(fn, ...bound)

Cria uma nova função com os primeiros argumentos já preenchidos:

```js
use "functional"
use "math"

let potenciaDe2 = functional.partial(math.pow, 2)
print(potenciaDe2(8))   # 256  (2^8)
print(potenciaDe2(10))  # 1024 (2^10)
```

## functional.once(fn)

Retorna uma função que executa `fn` **apenas na primeira chamada**. Chamadas subsequentes retornam o resultado cacheado:

```js
use "functional"

let contador = 0
let inicializar = functional.once(fn() {
    contador = contador + 1
    return "inicializado"
})

print(inicializar())  # inicializado
print(inicializar())  # inicializado (resultado cacheado, não executa de novo)
print(contador)       # 1 (executou só uma vez)
```

## functional.memo(fn)

Retorna uma versão **memoizada** da função — resultados são cacheados por argumentos:

```js
use "functional"

let chamadas = 0
let fatorial = functional.memo(fn(n) {
    chamadas++
    if (n <= 1) { return 1 }
    return n * fatorial(n - 1)
})

print(fatorial(5))   # 120
print(fatorial(5))   # 120 (resultado do cache, sem recalcular)
print(chamadas)      # chamou apenas na primeira vez
```

## Referência rápida

| Função                          | Descrição                                   |
|--------------------------------|---------------------------------------------|
| `functional.call(fn, ...args)` | Chamar builtin com argumentos               |
| `functional.apply(fn, arr)`    | Chamar builtin com array de argumentos      |
| `functional.partial(fn, ...)`  | Criar função com argumentos parciais        |
| `functional.once(fn)`          | Função que executa apenas uma vez           |
| `functional.memo(fn)`          | Função com cache de resultados              |
