# Criacao de Modulos Fig

Este guia mostra como criar um modulo externo para o Fig e publica-lo no GitHub.

## 1) Inicie um projeto

Crie a estrutura base:

```bash
fig init minha-biblioteca
cd minha-biblioteca
```

## 2) Ajuste o fig.toml para biblioteca

Abra o arquivo `fig.toml` e ajuste os campos principais:

- `project.type`: use "library"
- `project.name`: nome do modulo (use o mesmo do repo)
- `project.version`: versao inicial
- `project.main`: ponto de entrada do modulo
- `authors.name`: seu nome

Exemplo:

```toml
[project]
name = "minha-biblioteca"
version = "0.1.0"
description = "Utilitarios para Fig"
type = "library"
main = "src/main.fig"

[authors]
name = "Seu Nome"
```

## 3) Escreva o codigo do modulo

Coloque as funcoes e variaveis exportadas em `src/main.fig`:

```js
fn hello(name) {
    return "ola, " + name
}

let answer = 42
```

Esses nomes ficam disponiveis para quem importar o modulo.

## 4) Garanta o .gitignore

O diretorio `_modules/` deve estar no `.gitignore` do projeto:

```
_modules/
```

> Dependencias transitivas sao resolvidas no `_modules/` do projeto raiz. Modulos externos nao precisam levar `_modules` no repositorio.

## 5) Teste localmente

Crie um projeto de teste em outro diretorio e instale seu modulo:

```bash
fig init ../teste-modulo
cd ../teste-modulo
fig install seu-usuario/minha-biblioteca
```

Teste o import:

```js
import "mod:seu-usuario/minha-biblioteca" lib
print(lib.hello("Fig"))
print(lib.answer)
```

## 6) Publique no GitHub

Inicialize o git e publique:

```bash
git init
git add .
git commit -m "Inicializa modulo Fig"
git branch -M main
git remote add origin https://github.com/seu-usuario/minha-biblioteca.git
git push -u origin main
```

## 7) Como usuarios vao instalar

Depois de publicado, qualquer projeto Fig pode instalar:

```bash
fig install seu-usuario/minha-biblioteca
```

E importar com:

```js
import "mod:seu-usuario/minha-biblioteca" lib
```

## Dicas

- Use nomes simples e consistentes para funcoes e variaveis.
- Mantenha o `main` como API publica do modulo.
- Atualize `project.version` quando publicar novas versoes.
