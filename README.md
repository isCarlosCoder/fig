# Fig

**Fig** é uma linguagem de programação interpretada, dinâmica e de propósito geral, escrita em Go. Simples de aprender, poderosa o suficiente para scripts do dia a dia, APIs HTTP e automação.

Este projeto foi criado como um passa tempo e pode conter varios erros, bugs e comportamentos inesperados.

```js
use "math"

struct Circulo {
    raio

    fn init(raio) {
        this.raio = raio
    }

    fn area() {
        return math.PI * math.pow(this.raio, 2)
    }
}

let c = Circulo(5)
print("Área: " + c.area())  # Área: 78.53981633974483
```

## Instalação

```bash
git clone https://github.com/iscarloscoder/fig.git
cd fig
go build -o fig .
```

Opcional — instalar globalmente:

```bash
sudo mv fig /usr/local/bin/
```

## Uso

```bash
fig run programa.fig    # executar arquivo
fig run                # executar main via fig.toml
fig init meu-projeto   # criar projeto
fig install owner/repo # instalar modulo externo
fig --version           # versão (0.1.0)
fig --help              # ajuda
```

## Características

- Tipagem dinâmica com 6 tipos primitivos (number, string, boolean, null, array, object)
- Funções de primeira classe, closures e recursão
- Structs com campos, construtores e métodos
- Try/onerror para tratamento de erros
- 15 módulos embutidos com 150+ funções
- Cliente e servidor HTTP integrados
- Importação de arquivos `.fig`, módulos builtin e módulos externos
- Sintaxe limpa — ponto e vírgula opcional, comentários com `#`

## Documentação

### Linguagem

| # | Tópico | Descrição |
|---|--------|-----------|
| 1 | [Instalação e Primeiros Passos](docs/01-instalacao.md) | Instalação, execução, estrutura de um programa |
| 2 | [Variáveis e Tipos](docs/02-variaveis-e-tipos.md) | `let`, tipos primitivos, truthiness, escopo |
| 3 | [Operadores](docs/03-operadores.md) | Aritméticos, comparação, lógicos, precedência |
| 4 | [Controle de Fluxo](docs/04-controle-de-fluxo.md) | if/elif/else, while, do-while, for, for-in, range, enumerate, break, continue |
| 5 | [Funções](docs/05-funcoes.md) | Declaração, retorno, anônimas, closures, recursão |
| 6 | [Structs](docs/06-structs.md) | Structs, campos, init, métodos, `this` |
| 7 | [Imports e Módulos](docs/07-imports-e-modulos.md) | `import` (arquivos), `use` (builtins) |
| 8 | [Try / Onerror](docs/08-try-onerror.md) | Tratamento de erros, fallback, break/continue |
| 9 | [Print](docs/09-print.md) | Saída no terminal |
| 10 | [Criacao de Modulos](docs/10-criacao-de-modulos.md) | Como criar e publicar modulos Fig |
| 11 | [Testes com figtest](docs/11-figtest.md) | Framework de testes integrado |

### Módulos Embutidos

| Módulo | Descrição | Funções |
|--------|-----------|---------|
| [math](docs/modules/math.md) | Funções matemáticas e constantes | 18 funções + PI, E, INF |
| [strings](docs/modules/strings.md) | Manipulação de strings | 15 funções |
| [arrays](docs/modules/arrays.md) | Operações com arrays | 15 funções |
| [objects](docs/modules/objects.md) | Operações com objetos | 9 funções |
| [types](docs/modules/types.md) | Verificação e conversão de tipos | 12 funções |
| [io](docs/modules/io.md) | Entrada/saída e arquivos | 7 funções |
| [json](docs/modules/json.md) | Serialização e parsing JSON | 4 funções |
| [http](docs/modules/http.md) | Cliente e servidor HTTP | 12 funções |
| [system](docs/modules/system.md) | Sistema, tempo e ambiente | 8 funções |
| [crypto](docs/modules/crypto.md) | Hashes e codificação | 7 funções |
| [regex](docs/modules/regex.md) | Expressões regulares | 4 funções |
| [functional](docs/modules/functional.md) | Programação funcional | 5 funções |
| [debug](docs/modules/debug.md) | Depuração e diagnóstico | 5 funções |
| [runtime](docs/modules/runtime.md) | Informações do runtime | 5 funções |
| [task](docs/modules/task.md) | Concorrência com goroutines | 4 funções |
| [figtest](docs/modules/figtest.md) | Framework de testes integrado | 20 funções |

## Exemplos rápidos

### Variáveis e tipos

```js
let nome = "Fig"
let versao = 0.1
let ativo = true
let lista = [1, 2, 3]
let config = {debug: false, port: 8080}
```

### Funções e closures

```js
fn contador() {
    let n = 0
    return fn() {
        n++
        return n
    }
}

let c = contador()
print(c())  # 1
print(c())  # 2
```

### Servidor HTTP

```js
use "http"

http.route("GET", "/", fn(req, res) {
    res.json({mensagem: "Olá, Fig!"})
})

http.listen(3000)
```

### Tratamento de erros

```js
use "io"

let dados = try io.readFile("config.txt") onerror(e) {
    print("Aviso: " + e)
    return "padrão"
}
```

## Licença

MIT
