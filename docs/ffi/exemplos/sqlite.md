# Exemplo Avançado: Wrapper SQLite

Este exemplo demonstra como usar FFI para interagir com a biblioteca SQLite3 diretamente do Fig.

> **Nota:** Este é um exemplo avançado que envolve ponteiros opacos e gerenciamento cuidadoso de memória.

## Pré-requisitos

```bash
# Instalar SQLite3 dev (Ubuntu/Debian)
sudo apt install libsqlite3-dev

# Verificar se a lib está disponível
ldconfig -p | grep sqlite3
```

## 1. Wrapper C

Como o SQLite usa tipos complexos (`sqlite3*`, `sqlite3_stmt*`), criamos um wrapper C que expõe funções simplificadas:

Crie `sqlite_wrapper.c`:

```c
#include <sqlite3.h>
#include <stdlib.h>
#include <string.h>

static sqlite3* db = NULL;

// Abre (ou cria) um banco de dados
int db_open(const char* path) {
    if (db) sqlite3_close(db);
    return sqlite3_open(path, &db);
}

// Executa SQL sem retorno (CREATE, INSERT, DELETE, etc.)
// Retorna string alocada (caller deve liberar)
char* db_exec(const char* sql) {
    char* errmsg = NULL;
    int rc = sqlite3_exec(db, sql, NULL, NULL, &errmsg);
    if (rc != SQLITE_OK) {
        // Retorna a mensagem de erro (caller não deve liberar — é estática aqui)
        return errmsg;
    }
    return "";  // string vazia = sucesso
}

// Executa um SELECT e retorna a primeira coluna da primeira linha como string
const char* db_query_scalar(const char* sql) {
    sqlite3_stmt* stmt;
    int rc = sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
    if (rc != SQLITE_OK) return strdup(sqlite3_errmsg(db));

    const char* val = "";
    if (sqlite3_step(stmt) == SQLITE_ROW) {
        val = (const char*)sqlite3_column_text(stmt, 0);
    }

    sqlite3_finalize(stmt);
    return strdup(val ? val : "");
}

// Fecha o banco de dados
int db_close() {
    if (!db) return 0;
    int rc = sqlite3_close(db);
    db = NULL;
    return rc;
}
```

## 2. Compilação

```bash
gcc -shared -fPIC -o libsqlite_wrapper.so sqlite_wrapper.c -lsqlite3
```

## 3. Código Fig

```js
use "ffi"

var lib = ffi.load("./libsqlite_wrapper.so")

var db_open = ffi.sym(lib, "db_open", "int", ["string"])
var db_exec = ffi.sym(lib, "db_exec", "string", ["string"])
var db_query = ffi.sym(lib, "db_query_scalar", "string", ["string"])
var db_close = ffi.sym(lib, "db_close", "int", [])

// Abrir banco (em arquivo local)
var rc = ffi.call(db_open, "./data/myproject.db")
if rc != 0 {
    print("Erro ao abrir banco: " + rc)
}

// Criar tabela
ffi.call(db_exec, "CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")

// Inserir dados
ffi.call(db_exec, "INSERT INTO users (name) VALUES ('Alice')")
ffi.call(db_exec, "INSERT INTO users (name) VALUES ('Bob')")

// Consultar (scalar)
var nome = ffi.call(db_query, "SELECT name FROM users WHERE id = 1")
print("Usuário 1: " + nome)  // "Alice"

var count = ffi.call(db_query, "SELECT COUNT(*) FROM users")
print("Total: " + count)  // "2"

// Fechar
ffi.call(db_close)
print("Banco fechado com sucesso")

// --- Prepared statements ---
var db_prepare = ffi.sym(lib, "db_prepare", "int", ["string"])
var db_bind_text = ffi.sym(lib, "db_bind_text", "int", ["int", "int", "string", "int"])
var db_step = ffi.sym(lib, "db_step", "int", ["int"])
var db_column_text = ffi.sym(lib, "db_column_text", "string", ["int", "int"])
var db_finalize = ffi.sym(lib, "db_finalize", "int", ["int"])

var stmt = ffi.call(db_prepare, "SELECT name FROM users WHERE id = ?")
ffi.call(db_bind_text, stmt, 1, "1", -1)
var step = ffi.call(db_step, stmt)
if step == 1 {
    print("Resultado (prepared): " + ffi.call(db_column_text, stmt, 0))
}
ffi.call(db_finalize, stmt)
```

## Notas importantes

1. **Memória:** Funções que retornam `string` devem retornar memória alocada (ex.: `strdup`) porque o helper sempre chama `free()`.
2. **Thread safety:** O SQLite com um ponteiro global `db` não é thread-safe. Use um ponteiro por conexão em código real.
3. **Erros:** Verifique sempre os retornos — `db_open` retorna 0 em sucesso, `db_exec` retorna string vazia em sucesso.

## Ver também

- [Guia rápido](../quickstart.md)
- [Gerenciamento de memória](../memoria.md)
- [Tipos suportados](../tipos.md)
