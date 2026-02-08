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

## Referência rápida

| Função                       | Descrição                               |
|-----------------------------|------------------------------------------|
| `io.input(prompt?)`          | Ler entrada do teclado (com prompt)      |
| `io.readLine()`              | Ler linha da entrada padrão              |
| `io.readFile(path)`          | Ler conteúdo de um arquivo               |
| `io.writeFile(path, data)`   | Escrever/sobrescrever arquivo            |
| `io.appendFile(path, data)`  | Adicionar ao final do arquivo            |
| `io.exists(path)`            | Verificar se arquivo/diretório existe    |
| `io.deleteFile(path)`        | Remover arquivo                          |
