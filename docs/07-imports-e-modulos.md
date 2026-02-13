# Imports e Módulos

Fig possui dois mecanismos de importação:

- **`import`** — importa outro arquivo `.fig` (código do usuário)
- **`use`** — importa um módulo embutido (builtin) da linguagem
- **`import "mod:..."`** — importa um módulo externo instalado

---

## import — importando arquivos .fig

Use `import` para executar outro arquivo `.fig`. O arquivo importado é executado em um *ambiente de módulo* isolado e o conjunto de funções/variáveis exportadas fica disponível como um **objeto de módulo**. Acessar membros do módulo é feito via `modulo.membro()`.

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
# fornecemos um alias local (opcional) para o módulo
import "matematica.fig" mat

print(mat.somar(3, 5))     # 8
print(mat.subtrair(10, 4)) # 6
```

Se nenhum alias for informado, o interpretador derivará um nome de módulo a partir do nome do arquivo (por exemplo, `matematica.fig` → `matematica`).

### Regras do import

- O caminho é uma **string** com o caminho relativo do arquivo
- O arquivo importado é **executado** em seu próprio ambiente de módulo
- As funções e variáveis declaradas ficam disponíveis como propriedades do objeto do módulo (ex.: `matematica.somar()`)
- Você pode fornecer um alias opcional: `import "arquivo.fig" alias`
- A extensão `.fig` é opcional
- O import é resolvido em tempo de execução
- Se um import relativo não for encontrado no diretório do arquivo atual, o interpretador tentará resolver o caminho relativo ao diretório do projeto (quando aplicável), tornando imports como `src/utils.fig` mais tolerantes em diferentes locais do projeto
- Novo: você pode usar `*` como alias para `import "arquivo.fig" *` — isso "despeja" (inject) os símbolos top-level do arquivo importado no escopo do arquivo que importa (veja seção abaixo)

### Import com `*` — despejar símbolos (apenas import local)

Você pode usar o caractere `*` como alias em `import` para despejar (injetar) todos os símbolos top-level de um arquivo `.fig` importado diretamente no escopo do arquivo que realizou o import.

Exemplo:

```js
# arquivo: utilidades.fig
fn coisa() { return "ok" }
let VALOR = 42

# arquivo: main.fig
import "utilidades.fig" *
print(coisa())    # agora acessível diretamente
print(VALOR)      # 42
```

Regras e observações importantes:

- O `import "..." *` só funciona para **imports locais de arquivos `.fig`** (não funciona para `use` nem para `import "mod:..."`).
- O arquivo importado continua sendo executado em seu próprio ambiente, mas suas definições top-level são copiadas (definidas) no escopo do importador.
- Se um símbolo já existir no escopo do importador, a operação falhará com um erro de tempo de execução (conflito de nomes).
- Preferir `import "arquivo.fig" alias` quando quiser manter o namespace do módulo e evitar colisões.

---

### Importando módulos externos

Para módulos instalados em `_modules`, use o prefixo `mod:`. Recomendamos **importar pelo alias** registrado no FigRepo:

```js
import "mod:myfigtestdependency"
print(myfigtestdependency.magic)
```

Você também pode nomear o import com um identificador local:

```js
import "mod:myfigtestdependency" lib
print(lib.hello("Fig"))
```

Nota: a forma preferida para importar módulos é `mod:<alias>` (onde `<alias>` é o nome registrado no FigRepo).
> Para instalar um módulo registrado no FigRepo, use `fig install <alias>` no diretório do projeto. O comando resolve o alias no registry (`FIGREPO_BASE`, por padrão `https://figrepo.vercel.app`) e instala o repositório correspondente.
>
> Exemplo: `fig install logger`
>
> Para instalar vários de uma vez: `fig install logger color-tools`.
>
> Para sincronizar dependências do `fig.toml` (útil após clonar): `fig install` (sem argumentos).

> Para remover um módulo instalado, use `fig remove <alias>`. Por exemplo: `fig remove logger`.
> Para remover vários de uma vez: `fig remove a b`.
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
import "utils/helpers.fig" helpers
import "utils/validadores.fig" validators
print(helpers.someHelper())
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
| `date`       | Manipulação de datas e horários                     |

> Veja a documentação detalhada de cada módulo na pasta [modules/](modules/).

---

## Diferença entre `import` e `use`

| Aspecto     | `import`                                  | `use`                          |
|------------|-------------------------------------------|--------------------------------|
| Carrega     | Arquivo `.fig` do usuário                 | Módulo embutido (builtin)      |
| Caminho     | Relativo ao arquivo atual                 | Nome do módulo (string)        |
| Acesso      | Via objeto de módulo (`modulo.membro()`)  | Via prefixo `módulo.função()`  |
| Sintaxe     | `import "caminho/arquivo.fig"`           | `use "nomeDoModulo"`         |

### Import externo (mod)

| Aspecto     | `import "mod:..."`            |
|------------|--------------------------------|
| Carrega     | Módulo externo instalado        |
| Caminho     | `mod:package`            |
| Acesso      | Via nome do módulo ou alias     |
| Sintaxe     | `import "mod:package" alias` |

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
