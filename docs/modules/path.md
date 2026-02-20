# Módulo path

```js
use "path"
```

Funções utilitárias para manipulação de caminhos de arquivos e diretórios. Elas
são uma camada fina sobre o pacote `path/filepath` da biblioteca padrão, com
tipos e conversões compatíveis com Fig.

## path.join(...segments)

Concatena vários segmentos usando o separador correto da plataforma e normaliza
o resultado (remove `.` e `..` quando possível):

```js
print(path.join("/usr", "local", "bin"))   # /usr/local/bin
print(path.join("foo", "../bar"))          # bar
```

## path.base(p)

Retorna o último componente do caminho (nome do arquivo ou última pasta):

```js
print(path.base("/foo/bar/baz.txt"))  # baz.txt
```

## path.dir(p)

Obtém o diretório pai de um caminho. Se `p` não tiver barra, retorna `"."`:

```js
print(path.dir("/foo/bar/baz.txt"))  # /foo/bar
```

## path.ext(p)

Extraí a extensão (incluindo o `.`) de um caminho. Se não houver, retorna a
string vazia:

```js
print(path.ext("/foo/bar/baz.txt"))  # .txt
```

## path.abs(p)

Converte o caminho `p` para sua forma absoluta, usando o diretório de trabalho
atual como base.

```js
print(path.abs(".") )  # ex: /home/user/projeto
```

## path.clean(p)

Normaliza um caminho removendo `.` e resolvendo `..` onde possível. Não acessa
o sistema de arquivos.

```js
print(path.clean("foo/../bar"))  # bar
```

## path.isAbs(p)

Retorna `true` se o caminho for absoluto.

```js
print(path.isAbs("/foo"))  # true
```

## path.rel(base, target)

Calcula o caminho relativo de `target` em relação a `base`. Pode retornar um
caminho com `..`:

```js
print(path.rel("/a/b", "/a/b/c/d"))  # c/d
```

## path.split(p)

Divide um caminho em `[diretório, nomeDoArquivo]`.

```js
let parts = path.split("/foo/bar.txt")
print(parts)  # ["/foo/", "bar.txt"]
```

## path.splitExt(p)

Retorna `[raiz, extensão]` onde `raiz` é o caminho sem a extensão e
`extensão` é o `.` mais tudo depois dele.

```js
print(path.splitExt("foo.tar.gz"))  # ["foo.tar", ".gz"]
```

## path.exists(p)

Verifica se o arquivo ou diretório existe no sistema de arquivos. Equivalente
a `system.exists`, mas colocada aqui por conveniência quando se trabalha com
outros métodos de caminho.

## path.real(p)

Resolve links simbólicos e retorna o caminho absoluto resultante. Útil para
obter a localização "real" de um arquivo.

```js
print(path.real("/usr/bin/python"))  # /usr/bin/python3.9 (por exemplo)
```

## Exemplo rápido

```js
use "path"

let p = path.join(".", "src", "../README.md")
print(path.clean(p))
```

## Referência rápida

| Função                | Descrição                             |
|-----------------------|---------------------------------------|
| `path.join`           | Junta segmentos                       |
| `path.base`           | Último componente do caminho          |
| `path.dir`            | Diretório pai                         |
| `path.ext`            | Extensão do arquivo                   |
| `path.abs`            | Caminho absoluto                      |
| `path.clean`          | Normaliza (remove `.`/`..`)           |
| `path.isAbs`          | Testa se é absoluto                   |
| `path.rel`            | Caminho relativo                      |
| `path.split`          | Divide em `[dir, nome]`               |
| `path.splitExt`       | Divide raiz e extensão                |
| `path.exists`         | Existe no sistema de arquivos?        |
| `path.real`           | Resolve links simbólicos              |
