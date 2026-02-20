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

### Ferramentas adicionais (FFI)

O projeto inclui utilitários usados para trabalhar com FFI: `ffi-gen` (gerador de bindings) e `ffi-helper` (processo helper que carrega bibliotecas C). Para compilar todos os binários de uma vez, use o script de instalação multiplataforma na raiz do projeto:

- Linux / macOS:

```bash
# Torne o script executável se necessário
chmod +x ./install
./install
```

- Windows (PowerShell):

```powershell
# Execute no PowerShell com permissões necessárias
.\\install.ps1
```

O script tenta compilar `fig`, `ffi-gen` e `ffi-helper` para a plataforma atual. `ffi-helper` exige um toolchain C (gcc/clang) por causa do código C usado internamente; o script avisa se o compilador C não for encontrado.

Para builds manuais:

```bash
# Build local (binários na pasta bin/)
go build -o bin/ffi-gen ./tools/ffi-gen
# Precisa de C toolchain para ffi-helper
go build -o bin/ffi-helper ./tools/ffi-helper
go build -o bin/fig .
```

Para cross-compilar, ajuste `GOOS`/`GOARCH` e garanta toolchain C para a plataforma alvo (cross-compile de código com cgo requer toolchain correspondente).

## Uso

```bash
fig run programa.fig    # executar arquivo
fig run                # executar main via fig.toml
fig init meu-projeto   # criar projeto
fig install <alias>      # instalar um módulo registrado no FigRepo (ex.: logger)
fig install a b          # instalar vários aliases de uma vez
fig install              # sincronizar deps do fig.toml
fig remove <alias>       # remover módulo (p.ex.: logger)
fig remove a b           # remover vários de uma vez
fig --version           # versão (0.1.0)
fig --help              # ajuda
```

## Convenções de Estilo

Recomendações para manter o código Fig organizado e consistente:

### Ordem do arquivo

Sempre declare importações no topo do arquivo, seguidas de `use`, e depois o código:

```js
# 1. imports de arquivos e módulos externos
# Import local: cria um objeto de módulo (alias opcional)
import "utils/helpers.fig" utils
# Import externo (mod:) mantém o alias
import "mod:logger" log

# 2. módulos embutidos
use "arrays"
use "strings"

# 3. código
let dados = [3, 1, 2]
arrays.sort(dados)
```

### Organização de projeto

- Um arquivo por função ou struct — evite arquivos com muitas responsabilidades
- Agrupe arquivos relacionados em pastas (`utils/`, `models/`, `handlers/`)
- Use `src/main.fig` como ponto de entrada e importe o restante

```
meu-projeto/
├── src/
│   ├── main.fig
│   ├── models/
│   │   ├── usuario.fig
│   │   └── produto.fig
│   ├── utils/
│   │   ├── validacao.fig
│   │   └── formatacao.fig
│   └── handlers/
│       └── api.fig
├── tests/
│    └── validacao_test.fig
└── fig.toml
```

### Boas práticas

- Nomes de variáveis e funções em **camelCase**: `minhaVariavel`, `calcularTotal()`
- Nomes de structs em **PascalCase**: `Usuario`, `ItemPedido`
- Comentários com `#` — prefira explicar o *porquê*, não o *quê*
- Ponto e vírgula é opcional — omita para código mais limpo

## Características

- Tipagem dinâmica com 6 tipos primitivos (number, string, boolean, null, array, object)
- Funções de primeira classe, closures e recursão
- Structs com campos, construtores e métodos
- Try/onerror para tratamento de erros
- 16 módulos embutidos com 150+ funções
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
| 12 | [Enums](docs/12-enums.md) | Definição e uso de enums |
| 13 | [FFI — Bibliotecas Nativas](docs/13-ffi.md) | Interface com código C/nativo via FFI |
| 14 | [Funções nativas (@native)](docs/14-native.md) | Marcar funções para execução via fast‑path nativa (validação/compilação) |
| — | [FFI — Referência Detalhada](docs/ffi/index.md) | Protocolo, tipos, erros, memória, sandbox, exemplos |

### Módulos Embutidos

| Módulo | Descrição | Funções |
|--------|-----------|---------|
| [math](docs/modules/math.md) | Funções matemáticas e constantes | 18 funções + PI, E, INF |
| [strings](docs/modules/strings.md) | Manipulação de strings | 15 funções |
| [arrays](docs/modules/arrays.md) | Operações com arrays | 15 funções |
| [objects](docs/modules/objects.md) | Operações com objetos | 9 funções |
| [types](docs/modules/types.md) | Verificação e conversão de tipos | 12 funções |
| [io](docs/modules/io.md) | Entrada/saída e arquivos | 15 funções |
| [json](docs/modules/json.md) | Serialização e parsing JSON | 4 funções |
| [http](docs/modules/http.md) | Cliente e servidor HTTP | 12 funções |
| [system](docs/modules/system.md) | Sistema, tempo e ambiente | 16 funções |
| [term](docs/modules/term.md) | Terminal control and utilities (experimental) | 30+ funções |
| [crypto](docs/modules/crypto.md) | Hashes e codificação | 7 funções |
| [regex](docs/modules/regex.md) | Expressões regulares | 4 funções |
| [functional](docs/modules/functional.md) | Programação funcional | 5 funções |
| [debug](docs/modules/debug.md) | Depuração e diagnóstico | 5 funções |
| [utils](docs/modules/utils.md) | Utilitários | 17 funções |
| [runtime](docs/modules/runtime.md) | Informações do runtime | 5 funções |
| [task](docs/modules/task.md) | Concorrência com goroutines | 4 funções |
| [date](docs/modules/date.md) | Manipulação de datas e horários | 6 funções |
| [figtest](docs/modules/figtest.md) | Framework de testes integrado | 23 funções |

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
