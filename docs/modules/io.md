# Módulo io

```js
use "io"
```

Entrada do usuário e operações com arquivos.

## Entrada do usuário

### io.input(prompt?)

Exibe um prompt (opcional) e lê uma linha do teclado:

```js
let nome = io.input("Qual é o seu nome? ")
print("Olá, " + nome + "!")
```

Sem prompt:

```js
let valor = io.input()
```

### io.readLine()

Lê uma linha da entrada padrão (sem prompt):

```js
print("Digite algo:")
let linha = io.readLine()
print("Você digitou: " + linha)
```

## Leitura de arquivos

### io.readFile(caminho)

Lê o conteúdo completo de um arquivo e retorna como string:

```js
let conteudo = io.readFile("dados.txt")
print(conteudo)
```

Se o arquivo não existir, gera um erro de runtime (use `try/onerror` para tratar):

```js
let texto = try io.readFile("config.txt") onerror(e) {
    print("Arquivo não encontrado: " + e)
    return ""
}
```

## Escrita de arquivos

### io.writeFile(caminho, dados)

Escreve dados em um arquivo. Se o arquivo já existir, ele é **sobrescrito**. Se não existir, é criado:

```js
io.writeFile("saida.txt", "Olá, mundo!")
```

### io.appendFile(caminho, dados)

Adiciona dados ao **final** de um arquivo existente. Se o arquivo não existir, é criado:

```js
io.appendFile("log.txt", "Linha 1\n")
io.appendFile("log.txt", "Linha 2\n")
```

## Verificação e remoção

### io.exists(caminho)

Retorna `true` se o arquivo ou diretório existe:

```js
if (io.exists("config.txt")) {
    print("Arquivo encontrado!")
} else {
    print("Arquivo não existe")
}
```

### io.deleteFile(caminho)

Remove um arquivo:

```js
io.deleteFile("temporario.txt")
```

Se o arquivo não existir, gera um erro.

## Diretórios (novas funções)

As funções a seguir operam de forma multiplataforma usando as rotinas padrão do Go — elas funcionam tanto em Linux/macOS quanto em Windows.

### io.isDir(caminho)

Retorna `true` se o caminho existe e for um diretório; `false` se não existir ou for um arquivo.

```js
if (io.isDir("/tmp")) {
    print("/tmp é um diretório")
}
```

### io.mkdir(caminho)

Cria um diretório único; falha se o diretório já existir.

```js
io.mkdir("logs")
```

### io.mkdirAll(caminho)

Cria um diretório e todos os pais necessários (comportamento "mkdir -p").

```js
io.mkdirAll("tmp/data/2026")
```

### io.readDir(caminho)

Retorna um `array` com os nomes (strings) das entradas no diretório. Ordem não garantida.

```js
let itens = io.readDir("./src")
print(itens[0])
```

### io.rmdir(caminho)

Remove um diretório **vazio**; falha se não for vazio.

```js
io.rmdir("olddir")
```

### io.rmdirAll(caminho)

Remove um diretório e todo o seu conteúdo recursivamente (cuidado — operação destrutiva).

```js
io.rmdirAll("tmp/build")
```

> Observação: `io.rmdirAll` é equivalente a `rm -rf` — use com cuidado.

## Exemplo completo

```js
use "io"

# Escrevendo um arquivo
io.writeFile("lista.txt", "Item 1\nItem 2\nItem 3\n")

# Verificando se existe
if (io.exists("lista.txt")) {
    # Lendo o conteúdo
    let conteudo = io.readFile("lista.txt")
    print(conteudo)

    # Adicionando mais itens
    io.appendFile("lista.txt", "Item 4\n")

    # Limpando
    io.deleteFile("lista.txt")
    print("Arquivo removido")
}
```

## CSV (em desenvolvimento)

### io.writeCSV(path, data)

```js
use "io"

let tabela = [
    { nome: "Alice", idade: 30 },
    { nome: "Bob", idade: 25 }
]

io.writeCSV("pessoas.csv", tabela)
```

Para ler CSV, use `io.readCSV(path)` que retorna um array de objetos, onde cada objeto representa uma linha com chaves como os nomes das colunas.

```js
let pessoas = io.readCSV("pessoas.csv")

for pessoa in pessoas {
    print(pessoa.nome + " tem " + pessoa.idade + " anos")
}
```

saida:

```bash
Alice tem 30 anos
Bob tem 25 anos
```

## Referência rápida

| Função                       | Descrição                               |
|------------------------------|------------------------------------------|
| `io.input(prompt?)`          | Ler entrada do teclado (com prompt)      |
| `io.readLine()`              | Ler linha da entrada padrão              |
| `io.readFile(path)`          | Ler conteúdo de um arquivo               |
| `io.writeFile(path, data)`   | Escrever/sobrescrever arquivo            |
| `io.appendFile(path, data)`  | Adicionar ao final do arquivo            |
| `io.exists(path)`            | Verificar se arquivo/diretório existe    |
| `io.isDir(path)`             | Verificar se o caminho é um diretório    |
| `io.mkdir(path)`             | Criar diretório (falha se existir)       |
| `io.mkdirAll(path)`          | Criar diretório recursivamente (mkdir -p)|
| `io.readDir(path)`           | Listar entradas de um diretório         |
| `io.rmdir(path)`             | Remover diretório vazio                  |
| `io.rmdirAll(path)`          | Remover diretório recursivamente (rm -rf)|
| `io.deleteFile(path)`        | Remover arquivo                          |
| `io.writeCSV(path, data)`    | Escrever array de objetos em arquivo CSV |
| `io.readCSV(path)`           | Ler arquivo CSV como array de objetos    |
