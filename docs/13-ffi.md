# FFI — Interface com Bibliotecas Nativas

> **AVISO: adição experimental** — as funcionalidades de Sandbox e integração ASAN/Valgrind nesta branch são experimentais. Podem existir bugs ou regressões; use em ambientes de teste. Atualizaremos a documentação conforme as mudanças forem validadas.

> **Documentação detalhada:** Para referência completa do protocolo, tipos, erros, memória, sandbox e exemplos, consulte [`docs/ffi/`](ffi/index.md).

O módulo `ffi` permite que programas Fig chamem funções de bibliotecas nativas escritas em C (ou qualquer linguagem que exporte símbolos com convenção C). Com FFI, é possível reutilizar código nativo existente, acessar APIs do sistema e criar extensões de alta performance.

## Importação

```js
use "ffi"
```

## Configuração

O FFI precisa ser habilitado no `fig.toml` do projeto:

```toml
[ffi]
enabled = true
helper = "/caminho/para/ffi-helper"
call_timeout = 5000   # opcional, timeout em ms (padrão: 3000)
```

| Campo          | Tipo   | Obrigatório | Padrão | Descrição |
|----------------|--------|-------------|--------|-----------|
| `enabled`      | bool   | sim         | false  | Ativa o módulo FFI |
| `helper`       | string | sim         | —      | Caminho para o binário helper |
| `call_timeout` | int    | não         | 3000   | Timeout de chamada (ms) |

Para compilar o helper:

```bash
cd fig
go build -o ffi-helper ./tools/ffi-helper
```

## Verificando se FFI está ativo

```js
use "ffi"

if ffi.enabled() {
    print("FFI está habilitado!")
}
```

## Fluxo básico: load → sym → call

### 1. Carregar biblioteca

```js
let lib = ffi.load("./libminhalib.so")
```

### 2. Resolver símbolo (função)

```js
let add = ffi.sym(lib, "add", "int")
```

O terceiro argumento é o tipo de retorno: `"int"`, `"double"`, `"string"` ou `"void"`.

Se a função possuir tipos mistos, passe o `argTypes` como **array** no 4º argumento:

```js
let add = ffi.sym(lib, "add", "int", ["int", "int"])
let open = ffi.sym(lib, "db_open", "int", ["string"])
```

### 3. Chamar a função

```js
let resultado = ffi.call(add, 10, 20)
print(resultado)   # 30
```

## Exemplo completo

Dada uma biblioteca C:

```c
// mymath.c
int add(int a, int b) {
    return a + b;
}

double circle_area(double radius) {
    return 3.14159265 * radius * radius;
}

const char* greet(const char* name) {
    static char buf[256];
    snprintf(buf, sizeof(buf), "Olá, %s!", name);
    return buf;
}

void log_msg(const char* msg) {
    printf("[LOG] %s\n", msg);
}
```

Compilar:

```bash
gcc -shared -fPIC -o libmymath.so mymath.c
```

Usar em Fig:

```js
use "ffi"

let lib = ffi.load("./libmymath.so")

# int add(int, int)
let add = ffi.sym(lib, "add", "int")
print(ffi.call(add, 10, 20))          # 30

# double circle_area(double)
let area = ffi.sym(lib, "circle_area", "double")
print(ffi.call(area, 5.0))            # 78.5398...

# char* greet(char*)
let greet = ffi.sym(lib, "greet", "string")
print(ffi.call(greet, "Fig"))          # "Olá, Fig!"

# void log_msg(char*)
let log = ffi.sym(lib, "log_msg", "void")
ffi.call(log, "Sistema iniciado")      # imprime no stdout, retorna nil
```

## Tipos de retorno

| Tipo Fig     | C equivalente    | Retorno em Fig |
|-------------|------------------|----------------|
| `"int"`     | `int` (32-bit)   | number         |
| `"double"`  | `double` (64-bit)| number         |
| `"string"`  | `char*`          | string         |
| `"void"`    | `void`           | nil            |

## Quantidade de argumentos suportada

| Retorno    | 0 | 1 | 2 | 3 | 4 |
|-----------|---|---|---|---|---|
| `"int"`   | ✓ | ✓ | ✓ | ✓ | * |
| `"double"`| ✓ | ✓ | ✓ | — | — |
| `"string"`| ✓ | ✓ | ✓ | ✓ | ✓ |
| `"void"`  | ✓ | ✓ | ✓ | — | — |

* `"int"` suporta até 3 argumentos nativos (0..3) e **algumas** assinaturas mistas com `arg_types` (ex.: `"iisi"`).

## Tipos de argumento (`arg_types`)

Quando a função tem argumentos mistos (ex.: `string` e `int`), declare os tipos na chamada `sym`:

```js
let f = ffi.sym(lib, "pessoa_info", "string", ["string", "int", "string", "int"])
# arg_types = ["string", "int", "string", "int"]
```

Tipos de argumento válidos: `"int"`, `"double"`, `"string"`, `"struct:NomeDoStruct"`.

## Structs

### Definindo um struct (API antiga)

```js
ffi.define_struct("Point", [
    { name: "x", type: "int" },
    { name: "y", type: "int" }
])
```

### Nova API de wrapper para structs

A partir da versão 0.1.x existe um construtor de structs de alto nível que
automatiza a definição e a validação. Ele cria um "tipo" Fig com métodos
úteis (`new`, `validate`, `flatten`) e consulta automaticamente o esquema
interno ao registrar símbolos.

```js
let Point = ffi.struct("Point", [
    { name: "x", type: "int" },
    { name: "y", type: "int" }
])

let p1 = Point.new(10, 20)            # positional
let p2 = Point.new({ y: 5, x: 3 })    # named map

// você pode usar o descritor diretamente em argTypes
let addPoint = ffi.sym(lib, "add_point", "int", [Point, Point])
// o runtime converterá para ["struct:Point","struct:Point"]

// validação explícita
Point.validate(p2)   # true or error

// converter em lista de valores/tipos
let rep = Point.flatten(p1)
print(rep.values)     # [10,20]
print(rep.types)      # ["int","int"]
```

Você ainda pode usar `ffi.define_struct` e montar objetos manualmente; a nova
API é compatível e pode ser adotada gradualmente.

### Structs aninhados

```js
ffi.define_struct("Address", [
    { name: "city", type: "string" },
    { name: "zip",  type: "int" }
])

ffi.define_struct("Person", [
    { name: "name", type: "string" },
    { name: "age",  type: "int" },
    { name: "addr", type: "struct:Address" }
])
```

### Structs aninhados

```js
ffi.define_struct("Address", [
    { name: "city", type: "string" },
    { name: "zip",  type: "int" }
])

ffi.define_struct("Person", [
    { name: "name", type: "string" },
    { name: "age",  type: "int" },
    { name: "addr", type: "struct:Address" }
])
```

### Passando structs

Structs são expandidos em campos individuais (flat ABI). Um `Person` com `Address` aninhado resulta em 4 argumentos na chamada C:

```js
let sym = ffi.sym(lib, "person_info", "string", "struct:Person")

let result = ffi.call(sym, {
    __struct__: "Person",
    name: "Carlos",
    age: 30,
    addr: {
        __struct__: "Address",
        city: "Curitiba",
        zip: 80000
    }
})
```

### Erros comuns com structs

- `call: unknown struct schema: X` — o schema `X` não foi definido pelo `ffi.define_struct`.
- `missing struct field: Person.age` — o objeto passado não contém um campo esperado.
- `call: arg N expects TYPE, got <value>` — valor de campo tem tipo incorreto (por exemplo, string em um campo `int`).

Essas mensagens são geradas pelo client em Go antes do JSON ser enviado ao helper,
então elas aparecem imediatamente no runtime Fig.
A função C correspondente recebe os campos expandidos:

```c
// person_info(char* name, int age, char* city, int zip)
const char* person_info(const char* name, int age, const char* city, int zip);
```

## Callbacks (C → Fig)

Callbacks permitem que funções C chamem funções Fig durante a execução.

### Registrando um callback

```js
let cb = ffi.register_callback(fn(arg) {
    return arg + "!"
})
```

### Usando em chamadas

```js
let sym = ffi.sym(lib, "call_with_prefix", "string")
let result = ffi.call(sym, cb, "prefix:")
print(result)   # "prefix:world!"
```

### Removendo callback

```js
ffi.unregister_callback(cb)
```

### Limitações de callbacks

- Callbacks são **serializados** — apenas um executa por vez (limitação do CGo).
- Timeout fixo de **2 segundos** por callback.
- Evite uso intensivo de callbacks em chamadas altamente concorrentes.

## Memória e ownership

### Regra geral

> **Quem aloca, libera.**

### Alocação e liberação

```js
# Alocar 256 bytes (zerados)
let mem = ffi.alloc(256)
# mem = { __mem__: "m-1", size: 256 }

# Liberar
ffi.free(mem)
```

### Strings nativas

`strdup` copia uma string Fig para memória nativa C:

```js
let s = ffi.strdup("hello world")
# s = { __mem__: "m-1", size: 12 }

# usar em chamadas C...

ffi.free_string(s)   # ou ffi.free(s)
```

### Leitura e escrita de memória

```js
let mem = ffi.alloc(64)

# Escrever dados
let dados = ffi.bytes_from_string("conteúdo")
ffi.mem_write(mem, dados)

# Ler dados
let lido = ffi.mem_read(mem, 64)
let texto = ffi.bytes_to_string(lido)

ffi.free(mem)
```

### Tabela de ownership

| Quem alocou        | Quem libera  | Função              |
|---------------------|-------------|---------------------|
| `ffi.alloc()`      | código Fig   | `ffi.free()`        |
| `ffi.strdup()`     | código Fig   | `ffi.free_string()` |
| retorno de `call`   | automático  | helper libera       |

## Bytes e dados binários

Dados binários são transportados como objetos `{ __bytes__: "base64..." }`.

```js
# String → bytes
let b = ffi.bytes_from_string("ABC")

# Bytes → array numérico
let arr = ffi.bytes_to_array(b)    # [65, 66, 67]

# Array → bytes
let b2 = ffi.bytes_from_array([72, 105])

# Bytes → string
let s = ffi.bytes_to_string(b2)    # "Hi"
```

## Timeout e cancelamento

### Timeout de chamada

Configurável via `fig.toml`:

```toml
[ffi]
call_timeout = 5000   # 5 segundos
```

Se uma chamada exceder o timeout:

```
call: timeout waiting for result
```

### Timeout de callback

Callbacks têm timeout fixo de **2 segundos**.

## Chamadas dinâmicas com `call_raw`

`call_raw` aceita os argumentos como **array** e os transmite ao helper sem
nenhum processamento adicional. Isso é útil quando você constrói a lista de
parâmetros em tempo de execução ou quando não sabe de antemão quantos argumentos
serão necessários. O valor retornado é literalmente o que o helper devolver ao
runtime, portanto você pode inspecionar a resposta bruta sem conversão.

```js
let result = ffi.call_raw(sym_id, [10, 20, 30])
```
se o helper apenas ecoar os argumentos, `result` será [10, 20, 30]


## Comandos diretos ao helper

`helper_cmd` envia um comando genérico ao helper process. O segundo parâmetro
`data` é opcional; ele pode ser um número, string, array ou objeto e será incluído
na mensagem JSON. Quando omitido, apenas o nome do comando é enviado. O helper
retorna **exatamente** o JSON fornecido pelo processo auxiliar, sem nenhuma
conversão adicional.

```js
let resp = ffi.helper_cmd("ping")            # { ok: true, resp: "pong" }
let info = ffi.helper_cmd("stats", { })      # depende do que o helper suporta
```

> Dica: `helper_cmd` é especialmente útil em scripts de debug ou para acionar
> funcionalidades específicas do helper que ainda não possuem builtins em Fig.

## Testando a conexão

```js
let r = ffi.ping()
print(r)
```

## Referência de funções

| Função | Descrição |
|--------|-----------|
| `ffi.enabled()` | Retorna se FFI está habilitado |
| `ffi.lib_ext()` | Extensão da plataforma: `".so"`, `".dylib"` ou `".dll"` |
| `ffi.lib_name(base)` | Nome da lib: `"lib<base>.so"` / `".dylib"` / `"<base>.dll"` |
| `ffi.ping()` | Testa conexão com helper, retorna `"pong"` |
| `ffi.load(path)` | Carrega `.so`/`.dylib`/`.dll`, retorna handle |
| `ffi.sym(handle, nome, retorno, [argTypes])` | Resolve símbolo, retorna symbol id |
| `ffi.call(symId, ...args)` | Chama função nativa |
| `ffi.call_raw(symId, argsArray)` | Chama com args em array |
| `ffi.define_struct(nome, campos)` | Registra schema de struct |
| `ffi.register_callback(fn)` | Registra callback Fig invocável por C |
| `ffi.unregister_callback(cbObj)` | Remove callback |
| `ffi.helper_cmd(cmd, data)` | Comando genérico ao helper |
| `ffi.alloc(tamanho)` | Aloca memória zerada no helper |
| `ffi.free(memObj)` | Libera memória |
| `ffi.strdup(string)` | Copia string para memória nativa |
| `ffi.free_string(memObj)` | Libera string nativa (alias de `free`) |
| `ffi.mem_write(memObj, bytesObj)` | Escreve bytes na memória |
| `ffi.mem_read(memObj, tamanho)` | Lê bytes da memória |
| `ffi.bytes_from_string(str)` | String → bytes |
| `ffi.bytes_to_string(bytesObj)` | Bytes → string |
| `ffi.bytes_from_array(arr)` | Array numérico → bytes |
| `ffi.bytes_to_array(bytesObj)` | Bytes → array numérico |
| `ffi.sandbox_status()` | Retorna objeto com config e contadores do sandbox |

## Criando uma biblioteca FFI — passo a passo

### 1. Escrever o código C

```c
// meuplugin.c
#include <string.h>
#include <stdio.h>

int somar(int a, int b) {
    return a + b;
}

const char* saudacao(const char* nome) {
    static char buf[256];
    snprintf(buf, sizeof(buf), "Bem-vindo, %s!", nome);
    return buf;
}
```

### 2. Compilar como shared library

```bash
# Linux
gcc -shared -fPIC -o libmeuplugin.so meuplugin.c

# macOS
gcc -shared -fPIC -o libmeuplugin.dylib meuplugin.c
```

### 3. Configurar `fig.toml`

```toml
[ffi]
enabled = true
helper = "./ffi-helper"
```

### 4. Usar em Fig (portável)

```js
use "ffi"

let lib = ffi.load("./" + ffi.lib_name("meuplugin"))

let somar = ffi.sym(lib, "somar", "int")
print(ffi.call(somar, 5, 3))   # 8

let saudar = ffi.sym(lib, "saudacao", "string")
print(ffi.call(saudar, "Fig"))  # "Bem-vindo, Fig!"
```

## Erros comuns

| Erro | Causa | Solução |
|------|-------|---------|
| `ffi not enabled` | FFI desabilitado | Adicionar `enabled = true` no `fig.toml` |
| `cannot start helper` | Binário não encontrado | Verificar caminho em `helper` |
| `call: timeout` | Função demorou demais | Aumentar `call_timeout` |
| `unknown struct schema` | Struct não definido | Chamar `define_struct` antes |
| `sym: symbol not found` | Função não exportada | Verificar com `nm -D lib.so` |
| `sandbox: max loaded libraries limit` | Excedeu `max_libs` | Aumentar `max_libs` no `[ffi.sandbox]` |
| `sandbox: max live allocations limit` | Excedeu `max_allocs` | Liberar memória com `free`/`free_string` |
| `sandbox: helper restart limit exceeded` | Helper crashou demais | Verificar biblioteca C; aumentar `max_restarts` |

## Dicas

- Use `nm -D libX.so | grep funcao` para verificar se um símbolo está exportado.
- Compile com `-shared -fPIC` para compatibilidade.
- Sempre libere memória alocada com `alloc`/`strdup`.
- O helper process é reiniciado automaticamente se crashar.

## Cross-platform (multi-plataforma)

O FFI funciona em Linux, macOS e Windows. As diferenças de plataforma são abstraídas automaticamente.

### Extensões de biblioteca

| Plataforma | Extensão | Exemplo |
|------------|----------|----------|
| Linux      | `.so`    | `libmymath.so` |
| macOS      | `.dylib` | `libmymath.dylib` |
| Windows    | `.dll`   | `mymath.dll` |

### Helpers de plataforma

Use `ffi.lib_ext()` e `ffi.lib_name()` para escrever código portável:

```js
use "ffi"

# Extensão correta para a plataforma
let ext = ffi.lib_ext()      # ".so", ".dylib" ou ".dll"

# Nome completo da biblioteca
let name = ffi.lib_name("mymath")  # "libmymath.so" ou "libmymath.dylib" ou "mymath.dll"

# Carregar de forma portável
let lib = ffi.load("./" + ffi.lib_name("mymath"))
```

### Comunicação com o helper

| Plataforma | Transporte | Modo |
|------------|-----------|------|
| Linux      | Unix socket | `--socket` (padrão) |
| macOS      | Unix socket | `--socket` (padrão) |
| Windows    | stdin/stdout pipe | `--server` (automático) |

A seleção do transporte é automática — o código Fig não precisa saber qual está sendo usado.

### Dynamic loading interno

O helper usa abstração portável para carregamento dinâmico:

| Plataforma | API nativa |
|------------|----------------------------------|
| Linux      | `dlopen` / `dlsym` / `dlerror` |
| macOS      | `dlopen` / `dlsym` / `dlerror` |
| Windows    | `LoadLibraryA` / `GetProcAddress` / `FormatMessage` |

### Compilação da biblioteca C

```bash
# Linux
gcc -shared -fPIC -o libmymath.so mymath.c

# macOS
gcc -shared -fPIC -o libmymath.dylib mymath.c

# Windows (MinGW)
gcc -shared -o mymath.dll mymath.c
```

## Tooling: IDL, gerador de bindings e templates

O `ffi-gen` é uma ferramenta que lê um arquivo de definição (`ffi.def.toml`) e gera código Fig com os wrappers FFI automaticamente.

### Formato `ffi.def.toml`

```toml
[library]
name = "mymath"         # nome base (extensão adicionada automaticamente)
# path = "./libcustom"  # opcional: path explícito (sem extensão)

[[structs]]
name = "Point"
fields = [
    { name = "x", type = "int" },
    { name = "y", type = "int" }
]

[[functions]]
name = "add"            # nome da função Fig gerada
symbol = "add"          # símbolo C (opcional, default = name)
return = "int"          # tipo de retorno
args = ["int", "int"]   # tipos dos argumentos

[[functions]]
name = "greet"
return = "string"
args = ["string"]

[[functions]]
name = "circle_area"
return = "double"
args = ["double"]

[[functions]]
name = "noop"
return = "void"
args = []
```

### Gerando bindings

```bash
# Gerar para stdout
go run ./tools/ffi-gen -input mylib.ffi.def.toml

# Gerar para arquivo
go run ./tools/ffi-gen -input mylib.ffi.def.toml -output bindings.fig
```

Exemplo de saída gerada:

```js
# Auto-generated FFI bindings for mymath
# Generated by: fig ffi-gen — do not edit manually

use "ffi"

let __lib = ffi.load("./" + ffi.lib_name("mymath"))

let __sym_add = ffi.sym(__lib, "add", "int")
fn add(a, b) {
    return ffi.call(__sym_add, a, b)
}

let __sym_greet = ffi.sym(__lib, "greet", "string")
fn greet(a) {
    return ffi.call(__sym_greet, a)
}
```

### Scaffolding de projeto

Crie um projeto FFI completo com um único comando:

```bash
go run ./tools/ffi-gen -init meuprojeto
```

Isso cria a seguinte estrutura:

```
meuprojeto/
├── fig.toml                      # Configuração do projeto
├── meuprojeto.ffi.def.toml       # Definições FFI (IDL)
├── meuprojeto.c                  # Código C de exemplo
├── Makefile                      # Build da lib + geração de bindings
├── main.fig                      # Código Fig de exemplo
└── README.md                     # Documentação
```

### Tipos suportados na IDL

| Tipo no `ffi.def.toml` | Tipo C equivalente | Tipo Fig |
|------------------------|--------------------|----------|
| `"int"`               | `int`              | number   |
| `"double"`            | `double`           | number   |
| `"string"`            | `char*`            | string   |
| `"void"`              | `void`             | nil      |
| `"struct:Nome"`       | struct expandido   | object   |

## Sandbox e Políticas de Segurança

O FFI inclui um sistema de sandbox que limita recursos e controla a resiliência do helper process. As políticas são configuradas na seção `[ffi.sandbox]` do `fig.toml`.

### Configuração

```toml
[ffi]
enabled = true
helper = "./ffi-helper"

[ffi.sandbox]
max_memory_mb  = 512    # limite de memória do helper (RLIMIT_AS, Linux)
max_cpu_seconds = 60    # limite de CPU do helper (RLIMIT_CPU, Linux)
max_libs       = 10     # máximo de bibliotecas carregadas simultaneamente
max_allocs     = 1000   # máximo de alocações ativas (alloc + strdup)
max_restarts   = 5      # máximo de reinícios do helper antes de abortar
```

| Parâmetro | Default | Descrição |
|-----------|---------|-----------|
| `max_memory_mb` | `0` (ilimitado) | Limite de memória virtual do helper via `prlimit` (Linux) |
| `max_cpu_seconds` | `0` (ilimitado) | Limite de tempo de CPU via `prlimit` (Linux) |
| `max_libs` | `0` (ilimitado) | Número máximo de `ffi.load()` permitidos |
| `max_allocs` | `0` (ilimitado) | Número máximo de alocações ativas (`alloc` + `strdup`) |
| `max_restarts` | `5` | Reinícios do helper antes de retornar erro permanente |

> **Nota:** `max_memory_mb` e `max_cpu_seconds` usam `prlimit` (syscall Linux). Em outras plataformas são ignorados silenciosamente.

### Contadores e status

Use `ffi.sandbox_status()` para inspecionar a configuração e o estado atual:

```js
use "ffi"

let status = ffi.sandbox_status()
print(status.max_libs)       # 10
print(status.loaded_libs)    # 2
print(status.max_allocs)     # 1000
print(status.live_allocs)    # 0
print(status.max_restarts)   # 5
print(status.restarts)       # 0
```

### Resiliência e reinício do helper

Quando o helper process crasha, o FFI o reinicia automaticamente na próxima chamada. O sandbox rastreia quantas vezes isso acontece por projeto:

- Se `max_restarts > 0` e o contador de reinícios excede o limite, todas as chamadas FFI subsequentes retornam erro `"sandbox: helper restart limit exceeded"`.
- O contador é zerado quando `StopAllHelpers()` é chamado (final de execução).
- Isso protege contra loops infinitos de crash+restart em bibliotecas C defeituosas.

### Rastreamento de alocações

Cada `ffi.alloc()` e `ffi.strdup()` incrementam o contador de alocações ativas. Cada `ffi.free()` e `ffi.free_string()` decrementam. Se `max_allocs > 0` e o limite é atingido, novas alocações são bloqueadas com erro até que memória seja liberada.

## ASAN e Valgrind

O projeto inclui infraestrutura para verificação de erros de memória em bibliotecas C usadas via FFI.

### Abordagem: driver ASAN standalone

Como o Go runtime não é compatível com ASAN (o AddressSanitizer exige ser o primeiro na lista de bibliotecas), usamos um **driver C standalone** que é compilado com `-fsanitize=address` e testa as bibliotecas via `dlopen`:

```
tests/ffi_integration/asan_driver.c   # Driver que carrega e testa funções via dlopen
tests/ffi_integration/lib_asan_clean.c # Biblioteca C "limpa" para validação ASAN
```

### Compilando com ASAN

```bash
# Compilar a biblioteca com ASAN
gcc -shared -fPIC -fsanitize=address -g -o libtest.so mylib.c

# Compilar o driver com ASAN
gcc -fsanitize=address -g -o asan_driver tests/ffi_integration/asan_driver.c -ldl

# Executar (detecta leaks, buffer overflows, use-after-free)
ASAN_OPTIONS=detect_leaks=1 ./asan_driver ./libtest.so
```

### Scripts de verificação

Dois scripts prontos para uso estão em `tools/`:

```bash
# Verificar com ASAN (usa socat para comunicar com helper)
./tools/asan-check.sh ./libmeuplugin.so

# Verificar com Valgrind (usa helper em modo --server)
./tools/valgrind-check.sh ./libmeuplugin.so
```

### Testes automatizados

Os testes em `tests/ffi_asan_test.go` compilam automaticamente o driver e as bibliotecas com `-fsanitize=address` e verificam que não há erros de memória:

```bash
go test ./tests/ -run "TestAsan" -v -timeout 120s
```

| Teste | O que valida |
|-------|-------------|
| `TestAsanCleanLibrary` | Biblioteca ASAN-safe não tem erros |
| `TestAsanWithMainLib` | Biblioteca principal (lib.c) não tem erros |
| `TestAsanScriptExists` | Scripts de verificação existem e são executáveis |
| `TestValgrindAvailability` | Valgrind está instalado (skip se ausente) |
| `TestAsanBuildIntegration` | Build ASAN contém símbolos esperados |

