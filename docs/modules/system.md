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
| `system.now()`       | Timestamp Unix em milissegundos               |
| `system.clock()`     | Tempo em segundos (alta precisão)             |
| `system.sleep(ms)`   | Pausar por N milissegundos                    |
| `system.env(name)`   | Variável de ambiente                          |
| `system.args()`      | Argumentos do processo (`os.Args`)            |
| `system.argv()`      | Argumentos do script (passados ao `fig run`)  |
| `system.cwd()`       | Diretório de trabalho no momento da execução  |
| `system.platform()`  | Nome do sistema operacional                   |
| `system.version()`   | Versão da linguagem Fig                       |
| `system.exit(code?)` | Encerrar o programa                           |
