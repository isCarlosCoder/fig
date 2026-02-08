# Instalação e Primeiros Passos

## Requisitos

- **Go 1.21** ou superior instalado ([golang.org/dl](https://golang.org/dl/))

## Instalação

Clone o repositório e compile o executável:

```bash
git clone https://github.com/isCarlosCoder/fig.git
cd fig
go build -o fig .
```

Isso gera o executável `fig` na pasta do projeto.

### Instalação global (opcional)

Para usar o comando `fig` de qualquer lugar, mova o executável para um diretório no seu `PATH`:

```bash
sudo mv fig /usr/local/bin/
```

## Executando um arquivo Fig

Crie um arquivo com a extensão `.fig`:

```js
# hello.fig
print("Hello, World!")
```

Execute com:

```bash
fig run hello.fig
```

Saída:

```
Hello, World!
```

## Comandos disponíveis

| Comando            | Descrição                        |
|--------------------|----------------------------------|
| `fig run <arquivo>` | Executa um arquivo `.fig`       |
| `fig --help` ou `fig -h`  | Mostra a ajuda            |
| `fig --version` ou `fig -v` | Mostra a versão (ex: `0.1.0`) |

## Estrutura de um programa Fig

Um programa Fig é composto por **instruções** executadas de cima para baixo. Não existe uma função `main` — o código começa a rodar diretamente.

```js
# Isso é um comentário (usa #)
let nome = "Fig"
print("Olá, " + nome + "!")
```

### Ponto e vírgula

O ponto e vírgula (`;`) é **opcional** no final de cada instrução. Ambas as formas são válidas:

```js
let x = 10;
let y = 20
```

### Comentários

Comentários começam com `#` e vão até o final da linha:

```js
# Isso é um comentário
let x = 42  # Comentário ao lado do código
```

## Próximos passos

- [Variáveis e Tipos](02-variaveis-e-tipos.md) — aprenda a declarar variáveis e os tipos disponíveis
- [Operadores](03-operadores.md) — conheça os operadores aritméticos, lógicos e de comparação
