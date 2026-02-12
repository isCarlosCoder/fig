# ffi-gen — Gerador de Bindings FFI

O `ffi-gen` lê um arquivo `.ffi.def.toml` e gera código Fig (`.fig`) que chama `ffi.load`, `ffi.sym` e `ffi.call` automaticamente.

## Uso básico

```bash
# Gerar bindings para stdout
ffi-gen -input mylib.ffi.def.toml

# Gerar para arquivo
ffi-gen -input mylib.ffi.def.toml -output bindings.fig

# Scaffold de projeto novo
ffi-gen -init meu_projeto
```

## Flags

| Flag | Descrição |
|------|-----------|
| `-input <arquivo>` | Caminho para o `.ffi.def.toml` de entrada |
| `-output <arquivo>` | Arquivo `.fig` de saída (stdout se omitido) |
| `-init <nome>` | Cria um projeto FFI completo com exemplos |
| `-validate` | Apenas valida o `.ffi.def.toml` sem gerar código |
| `-check-helper <caminho>` | Verifica se o binário do helper existe e é executável |
| `-fig-toml <caminho>` | Valida configuração FFI no `fig.toml` do projeto |

## Formato do .ffi.def.toml

```toml
[library]
name = "mymath"      # Nome base da biblioteca
# path = "/opt/lib"  # Caminho explícito (opcional, sobrepõe name)

[[structs]]
name = "Point"
fields = [
    { name = "x", type = "int" },
    { name = "y", type = "int" }
]

[[functions]]
name = "add"           # Nome da função Fig gerada
symbol = "add"         # Nome do símbolo C (padrão: igual a name)
return = "int"         # Tipo de retorno
args = ["int", "int"]  # Tipos dos argumentos
```

## Tipos suportados

| Tipo | Descrição |
|------|-----------|
| `int` | Inteiro |
| `double` | Ponto flutuante |
| `string` | String C (`char*`) |
| `void` | Sem retorno |
| `struct:Nome` | Referência a um struct definido |

Tipos não reconhecidos geram erro:

```
❌ function bad_func: unknown return type: pointer. Supported: int, double, string, void, struct:Name
```

## Código gerado

O ffi-gen agora gera:

1. **Comentário com assinatura C** para cada função:
   ```
   # C: int add(int a, int b)
   ```

2. **argTypes sempre incluídos** no `ffi.sym()`:
   ```js
   let __sym_add = ffi.sym(__lib, "add", "int", ["int", "int"])
   ```

3. **Wrapper Fig** com parâmetros nomeados:
   ```js
   fn add(a, b) {
       return ffi.call(__sym_add, a, b)
   }
   ```

## Flag --validate

Valida o `.ffi.def.toml` sem gerar código:

```bash
ffi-gen -input mylib.ffi.def.toml -validate
# ✅ mylib.ffi.def.toml is valid

ffi-gen -input bad.ffi.def.toml -validate
# ❌ function bad: unknown return type: pointer. Supported: int, double, string, void, struct:Name
# exit code 1
```

## Flag --check-helper

Verifica se o binário do helper existe e é executável:

```bash
ffi-gen -check-helper ./ffi-helper
# ✅ Helper binary OK: ./ffi-helper

ffi-gen -check-helper /nonexistent
# ❌ helper binary not found: /nonexistent
# exit code 1
```

## Flag --fig-toml

Valida a configuração FFI no `fig.toml` (emite avisos, não bloqueia geração):

```bash
ffi-gen -input mylib.ffi.def.toml -fig-toml fig.toml
# ⚠️  fig.toml: ffi.enabled is not true
# ⚠️  fig.toml: ffi.helper binary not found: ./ffi-helper
# (geração prossegue normalmente)
```

## Ver também

- [Tipos suportados](tipos.md)
- [Helper — referência](helper.md)
- [Guia rápido](quickstart.md)
