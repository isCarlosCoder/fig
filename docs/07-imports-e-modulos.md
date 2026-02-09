# Imports e Módulos

Fig possui dois mecanismos de importação:

- **`import`** — importa outro arquivo `.fig` (código do usuário)
- **`use`** — importa um módulo embutido (builtin) da linguagem
- **`import "mod:..."`** — importa um módulo externo instalado

---

## import — importando arquivos .fig

Use `import` para executar outro arquivo `.fig`. O caminho é relativo ao arquivo atual:

```js
# arquivo: matematica.fig
fn somar(a, b) {
    return a + b
}

fn subtrair(a, b) {
    return a - b
}
```

```js
# arquivo: main.fig
import "matematica.fig"

print(somar(3, 5))     # 8
print(subtrair(10, 4))  # 6
```

### Regras do import

- O caminho é uma **string** com o caminho relativo do arquivo
- O arquivo importado é **executado** no escopo do importador
- Todas as funções e variáveis declaradas ficam disponíveis
- A extensão `.fig` é opcional
- O import é resolvido em tempo de execução

### Importando módulos externos

Para módulos instalados em `_modules`, use o prefixo `mod:`:

```js
import "mod:isCarlosCoder/myfigtestdependency"
print(myfigtestdependency.magic)
```

Você pode definir um alias:

```js
import "mod:isCarlosCoder/myfigtestdependency" lib
print(lib.hello("Fig"))
```

> Para instalar um módulo, use `fig install <owner>/<repo>` no diretório do projeto.
> Para instalar vários de uma vez: `fig install owner/x owner/y`.
> Para sincronizar dependências do `fig.toml` (útil após clonar): `fig install` (sem argumentos).

> Para remover um módulo: `fig remove owner/repo`.
> Para remover vários de uma vez: `fig remove owner/x owner/y`.
> Se outro módulo depender do que está sendo removido, o comando será bloqueado. Use `--force` para forçar.

> Dependências transitivas são resolvidas no `_modules/` do projeto raiz. Módulos externos não carregam um `_modules` próprio.

### Organização em pastas

Você pode organizar seus arquivos em pastas:

```
meu-projeto/
├── main.fig
├── utils/
│   ├── helpers.fig
│   └── validadores.fig
```

```js
# main.fig
import "utils/helpers.fig"
import "utils/validadores.fig"
```

---

## use — importando módulos embutidos

Use `use` para carregar módulos embutidos da linguagem:

```js
use "math"
use "strings"
use "arrays"
```

Após o `use`, as funções do módulo ficam disponíveis como `módulo.função()`:

```js
use "math"

print(math.PI)          # 3.141592653589793
print(math.sqrt(16))    # 4
print(math.abs(-7))     # 7
```

```js
use "strings"

print(strings.upper("olá"))     # OLÁ
print(strings.len("Fig"))       # 3
print(strings.contains("abc", "b"))  # true
```

### Módulos disponíveis

| Módulo       | Descrição                                           |
|-------------|-----------------------------------------------------|
| `math`       | Funções matemáticas e constantes (PI, E, sqrt, etc.)|
| `strings`    | Manipulação de strings (upper, lower, trim, etc.)   |
| `arrays`     | Operações com arrays (push, pop, sort, etc.)   |
| `objects`    | Operações com objetos (keys, values, merge, etc.)   |
| `types`      | Verificação e conversão de tipos                    |
| `system`     | Data/hora, sleep, variáveis de ambiente             |
| `functional` | Programação funcional (partial, once, memo, etc.)   |
| `crypto`     | Hashes e codificação (sha256, base64, etc.)         |
| `debug`      | Ferramentas de depuração (dump, inspect, assert)    |
| `json`       | Serialização e parsing de JSON                      |
| `regex`      | Expressões regulares                                |
| `runtime`    | Informações do runtime (memória, versão, etc.)      |
| `io`         | Leitura/escrita de arquivos e entrada do usuário    |
| `http`       | Cliente e servidor HTTP                             |

> Veja a documentação detalhada de cada módulo na pasta [modules/](modules/).

---

## Diferença entre `import` e `use`

| Aspecto     | `import`                       | `use`                          |
|------------|--------------------------------|--------------------------------|
| Carrega     | Arquivo `.fig` do usuário      | Módulo embutido (builtin)      |
| Caminho     | Relativo ao arquivo atual      | Nome do módulo (string)        |
| Acesso      | Direto (funções ficam no escopo)| Via prefixo `módulo.função()`  |
| Sintaxe     | `import "caminho/arquivo.fig"` | `use "nomeDoModulo"`           |

### Import externo (mod)

| Aspecto     | `import "mod:..."`            |
|------------|--------------------------------|
| Carrega     | Módulo externo instalado        |
| Caminho     | `mod:<owner>/<repo>`            |
| Acesso      | Via nome do módulo ou alias     |
| Sintaxe     | `import "mod:owner/repo" alias` |

---

## Exemplo completo

```js
# Importando módulos embutidos
use "math"
use "strings"

# Importando arquivo do usuário
import "utils.fig"

let raio = 5
let area = math.PI * math.pow(raio, 2)
let msg = "Área do círculo: " + area
print(strings.upper(msg))
```
