# Módulo system

```js
use "system"
```

Informações do sistema, tempo e controle de execução.

## Tempo

### system.now()

Retorna o timestamp Unix atual em **milissegundos**:

```js
let ts = system.now()
print(ts)  # ex: 1717234567890
```

### system.clock()

Retorna o tempo em **segundos** com precisão de nanossegundos (ideal para benchmarks):

```js
let inicio = system.clock()
# ... operação ...
let fim = system.clock()
print("Tempo: " + (fim - inicio) + "s")
```

### system.sleep(ms)

Pausa a execução por `ms` milissegundos:

```js
print("Esperando...")
system.sleep(1000)  # pausa 1 segundo
print("Pronto!")
```

## Ambiente

### system.env(nome)

Retorna o valor de uma variável de ambiente, ou `null` se não existir:

```js
let home = system.env("HOME")
print(home)  # ex: /home/carlos

let inexistente = system.env("VAR_QUE_NAO_EXISTE")
print(inexistente)  # null
```

### system.args()

Retorna os argumentos de linha de comando do processo como array (equivalente a `os.Args`):

```js
let argumentos = system.args()
print(argumentos)
```

### system.argv()

Quando você executa um script com `fig run <file> [args]`, o CLI injeta os argumentos do script (aqueles após o nome do arquivo) na função `system.argv()` como um `array` de strings. Exemplo de uso no shell:

```sh
fig run myscript.fig a b
```

No script:

```js
use "system"
print(system.argv()) // ["a", "b"]
```

### system.cwd()

A função `system.cwd()` retorna o diretório de trabalho atual (string) no momento da execução do script:

```js
use "system"
print(system.cwd()) // ex: /home/user/projects
```

### system.platform()

Retorna o nome do sistema operacional:

```js
print(system.platform())  # ex: linux, darwin, windows
```

### system.version()

Retorna a versão da linguagem Fig:

```js
print(system.version())  # 0.1.0
```

## Controle de execução

### system.exit(código?)

Encerra o programa com um código de saída (padrão: 0):

```js
if (erro) {
    print("Erro fatal!")
    system.exit(1)  # encerra com código 1
}
```

## Controle do limite de passos de avaliação (experimental)
Para evitar loops acidentais o interpretador conta "passos" de avaliação e dispara um erro quando o número ultrapassa um limite (protege contra recursão infinita e *spin loops*). Em alguns cenários legítimos (laços pesados, operações de bulk) você pode optar por desabilitar temporariamente essa checagem — use com muito cuidado.

- `system.disableStepLimit()` → null
  - Desabilita globalmente a checagem de passos de avaliação (process-wide).

- `system.enableStepLimit()` → null
  - Restaura a checagem de passos.

- `system.isStepLimitDisabled()` → boolean
  - Retorna `true` se o limite está atualmente desabilitado.

- `system.withoutStepLimit(fn)` → any
  - Executa a função `fn` com a checagem de passos temporariamente desabilitada. **Requer** que o módulo `task` esteja carregado (para executar `fn` na goroutine do interpretador e preservar ambiente/semântica); se `task` não for carregado, a chamada retorna um erro.

**Exemplo (desabilitar temporariamente):**

```js
use "system"
# Opção simples (global):
system.disableStepLimit()
let i = 0
while (i < 1000000) { i = i + 1 }
system.enableStepLimit()
print(i)
```

**Exemplo (wrapper seguro com task):**

```js
use "task"
use "system"
fn heavy() {
    let i = 0
    while (i < 1000000) { i = i + 1 }
    return i
}
let r = system.withoutStepLimit(heavy)
print(r)
```

> ⚠️ Atenção: uso responsável requerido
>
> - `disableStepLimit()` é process-wide: se usado inadvertidamente pode deixar sua aplicação presa em um loop sem proteção. Prefira `system.withoutStepLimit(fn)` quando possível, pois encapsula o escopo.
> - Essas funções não criam timeouts; se `fn` entrar em loop infinito o processo ficará bloqueado — **evite** desabilitar o limite para código que possa não terminar.
> - Documente e revise cuidadosamente o uso em código de produção. Recomenda-se usar apenas para operações determinísticas e bem testadas (bulk fills, migrações, cálculos intensivos que não dependem de I/O assíncrono).

## Exemplo: Benchmark simples

```js
use "system"

let inicio = system.clock()

let soma = 0
for i in range(0, 1000000) {
    soma = soma + i
}

let fim = system.clock()
print("Soma: " + soma)
print("Tempo: " + (fim - inicio) + "s")
```

## Referência rápida

| Função               | Descrição                                    |
|---------------------|-----------------------------------------------|
| `system.now()`            | Timestamp Unix em milissegundos               |
| `system.clock()`         | Tempo em segundos (alta precisão)             |
| `system.sleep(ms)`       | Pausar por N milissegundos                    |
| `system.env(name)`       | Variável de ambiente                          |
| `system.args()`          | Argumentos do processo (`os.Args`)            |
| `system.argv()`          | Argumentos do script (passados ao `fig run`)  |
| `system.cwd()`           | Diretório de trabalho no momento da execução  |
| `system.platform()`      | Nome do sistema operacional                   |
| `system.version()`       | Versão da linguagem Fig                       |
| `system.exit(code?)`     | Encerrar o programa                           |
| `system.disableStepLimit()` | Desabilita globalmente a checagem de passos (experimental) |
| `system.enableStepLimit()`  | Reabilita a checagem de passos (experimental) |
| `system.isStepLimitDisabled()` | Retorna boolean indicando se o limite está desabilitado |
| `system.withoutStepLimit(fn)`  | Executa `fn` com a checagem de passos desabilitada (requer `task`) |
