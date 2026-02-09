# figtest

Framework de testes integrado do FigLang.

```js
use "figtest"
```

## Funções Básicas

| Função | Descrição |
|--------|-----------|
| `test(name, fn)` | Define e executa um teste |
| `describe(name, fn)` | Agrupa testes |
| `skip(name, fn)` | Marca teste como ignorado |
| `only(name, fn)` | Executa apenas este teste |

## Asserts

| Função | Descrição |
|--------|-----------|
| `assert(cond, msg?)` | Falha se `cond` for falso |
| `assertEq(a, b, msg?)` | Falha se `a ≠ b` |
| `assertNeq(a, b, msg?)` | Falha se `a = b` |
| `assertError(fn, msg?)` | Passa se `fn` gerar erro |
| `assertNoError(fn, msg?)` | Falha se `fn` gerar erro |
| `assertNear(a, b, eps, msg?)` | Falha se `|a-b| > eps` |
| `assertContains(container, item, msg?)` | Verifica se contém item |
| `assertType(val, type, msg?)` | Verifica tipo |
| `assertLength(val, len, msg?)` | Verifica tamanho |

## Hooks

| Função | Descrição |
|--------|-----------|
| `beforeEach(fn)` | Executa antes de cada teste no `describe` |
| `afterEach(fn)` | Executa depois de cada teste no `describe` |
| `beforeAll(fn)` | Executa uma vez antes dos testes do `describe` |
| `afterAll(fn)` | Executa uma vez depois dos testes do `describe` |

## Relatórios

| Função | Descrição |
|--------|-----------|
| `summary()` | Retorna texto com resumo |
| `reset()` | Limpa resultados |
| `count()` | Total de testes |
| `passed()` | Testes que passaram |
| `failed()` | Testes que falharam |
| `skipped()` | Testes ignorados |

## Exemplo

```js
use "figtest"

figtest.describe("Operações", fn() {
    figtest.beforeEach(fn() {
        # setup
    })

    figtest.test("soma", fn() {
        figtest.assertEq(2 + 2, 4)
    })

    figtest.test("tipo", fn() {
        figtest.assertType(42, "number")
    })

    figtest.skip("wip", fn() {
        figtest.assert(false)
    })
})
```

## CLI

```bash
# Descoberta automática
fig test

# Arquivo específico
fig test tests/math_test.fig

# Com detalhes
fig test --verbose
```

Documentação completa: [docs/11-figtest.md](../../docs/11-figtest.md)
