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

- Não disponivel no momento.


## 6) Publique no GitHub

Inicialize o git e publique:

```bash
git init
git add .
git commit -m "Inicializa modulo Fig"
git branch -M main
git remote add origin https://github.com/<seu-usuario>/<seu-repo>.git
git push -u origin main
```

## 7) Como usuários vão instalar

Depois de publicado e registrado no FigRepo com um alias, qualquer projeto Fig pode instalar usando o alias:

```bash
fig install minha-biblioteca
```

E importar com o alias registrado no `mod:`:

```js
import "mod:minha-biblioteca" lib
```

> Nota: o `fig install` depende da existência do alias no registry. Se o alias não estiver registrado, o CLI retornará erro — não há fallback automático.

### Importando pelo alias

Ao registrar um alias no FigRepo, usuários poderão instalar seu módulo usando `fig install <alias>` e importar no código com `import "mod:<alias>"`. Esta é a forma recomendada porque deixa explícito o nome público do pacote independente do owner/repo do GitHub.
## Dicas

- Use nomes simples e consistentes para funcoes e variaveis.
- Mantenha o `main` como API publica do modulo.
- Atualize `project.version` quando publicar novas versoes.
