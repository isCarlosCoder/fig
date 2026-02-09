# figtest — Framework de Testes Integrado

O módulo `figtest` é a biblioteca de testes integrada do FigLang. Permite escrever testes automatizados, agrupar testes, capturar erros e fazer asserts avançados — sem dependências externas.

## Importação

```js
use "figtest"
```

## Funções Básicas

### `figtest.test(name, fn)`

Define e executa um teste simples.

```js
figtest.test("soma de inteiros", fn() {
    let x = 2 + 2
    figtest.assertEq(x, 4)
})
```

### `figtest.describe(name, fn)`

Agrupa testes para organizar a saída.

```js
figtest.describe("Operações matemáticas", fn() {
    figtest.test("soma", fn() { figtest.assertEq(2 + 2, 4) })
    figtest.test("mult", fn() { figtest.assertEq(3 * 3, 9) })
})
```

### `figtest.skip(name, fn)`

Marca um teste como ignorado (não executa).

```js
figtest.skip("wip - em desenvolvimento", fn() {
    figtest.assert(false)
})
```

### `figtest.only(name, fn)`

Executa apenas este teste, útil para depuração isolada.

```js
figtest.only("debugando este", fn() {
    figtest.assert(true)
})
```

## Asserts

### `figtest.assert(cond, msg?)`

Falha se `cond` for falso.

```js
figtest.assert(x > 0, "x deve ser positivo")
```

### `figtest.assertEq(a, b, msg?)`

Falha se `a` não for igual a `b`.

```js
figtest.assertEq(soma(2, 3), 5)
```

### `figtest.assertNeq(a, b, msg?)`

Falha se `a` for igual a `b`.

```js
figtest.assertNeq(1, 2)
```

### `figtest.assertError(fn, msg?)`

Passa se a função gerar um erro.

```js
figtest.assertError(fn() { figtest.assert(false, "boom") })
```

### `figtest.assertNoError(fn, msg?)`

Falha se a função gerar um erro.

```js
figtest.assertNoError(fn() { let x = 42 })
```

### `figtest.assertNear(a, b, epsilon, msg?)`

Falha se `|a - b| > epsilon`. Útil para comparações de ponto flutuante.

```js
figtest.assertNear(3.14159, 3.14, 0.01)
```

### `figtest.assertContains(container, item, msg?)`

Verifica se uma string ou array contém o item.

```js
figtest.assertContains("hello world", "world")
figtest.assertContains([1, 2, 3], 2)
```

### `figtest.assertType(val, typeName, msg?)`

Verifica se `val` tem o tipo especificado.

```js
figtest.assertType(42, "number")
figtest.assertType("hi", "string")
figtest.assertType(true, "boolean")
figtest.assertType([1], "array")
```

### `figtest.assertLength(val, len, msg?)`

Verifica se string ou array tem o tamanho esperado.

```js
figtest.assertLength("abc", 3)
figtest.assertLength([1, 2], 2)
```

## Hooks

### `figtest.beforeEach(fn)`

Executa antes de cada teste dentro de um `describe`.

```js
figtest.describe("com setup", fn() {
    let counter = 0
    figtest.beforeEach(fn() { counter = counter + 1 })
    figtest.test("primeiro", fn() { figtest.assertEq(counter, 1) })
    figtest.test("segundo", fn() { figtest.assertEq(counter, 2) })
})
```

### `figtest.afterEach(fn)`

Executa depois de cada teste dentro de um `describe`.

### `figtest.beforeAll(fn)`

Executa uma vez antes de todos os testes do `describe`.

```js
figtest.describe("com beforeAll", fn() {
    let ready = false
    figtest.beforeAll(fn() { ready = true })
    figtest.test("check", fn() { figtest.assert(ready) })
})
```

### `figtest.afterAll(fn)`

Executa uma vez depois de todos os testes do `describe`.

## Relatórios e Utilitários

| Função | Descrição |
|--------|-----------|
| `figtest.summary()` | Retorna texto com resumo dos testes |
| `figtest.reset()` | Limpa todos os resultados |
| `figtest.count()` | Número total de testes executados |
| `figtest.passed()` | Número de testes que passaram |
| `figtest.failed()` | Número de testes que falharam |
| `figtest.skipped()` | Número de testes ignorados |

## CLI — `fig test`

O comando `fig test` descobre e executa arquivos de teste automaticamente.

### Descoberta automática

```bash
fig test
```

Procura:
- `tests/*.fig` — todos os arquivos na pasta `tests/`
- `**/*_test.fig` — qualquer arquivo com sufixo `_test.fig` no projeto

### Execução de arquivos específicos

```bash
fig test tests/math_test.fig
fig test tests/*.fig
```

### Opções

| Flag | Descrição |
|------|-----------|
| `--verbose`, `-V` | Mostra detalhes adicionais |

### Saída

```
tests/math_test.fig

Operações matemáticas
  ✓ Operações matemáticas > soma
  ✓ Operações matemáticas > subtração

tests/types_test.fig

Strings
  ✓ Strings > contém substring
  ○ teste wip (skipped)

4 passed, 0 failed, 1 skipped (total: 5)
```

### Isolamento

Cada arquivo de teste é executado em seu próprio ambiente isolado:
- Variáveis de um arquivo não afetam outro
- Hooks (`beforeEach`, `beforeAll`) são por arquivo/describe
- O CLI agrega apenas os contadores totais

### Código de saída

- `0` — todos os testes passaram
- `1` — um ou mais testes falharam
