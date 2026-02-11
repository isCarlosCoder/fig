# Módulo term

```js
use "term"
```

Controle básico de terminal, leitura de teclas em modo raw, buffers em memória e utilitários de estilo. **Experimental** — sujeito a remoção no futuro.

## term.startRaw() / term.stopRaw() / term.isRaw()
Coloca o stdin em modo raw (sem linha-canônica) para leitura imediata de bytes.

```js
term.startRaw()
print(term.isRaw())  # true
term.stopRaw()
print(term.isRaw())  # false
```

## Leitura de teclas

### term.readKey(blocking=true, timeout_ms=0)
Lê bytes brutos do stdin (retorna string) — `null` quando não há dados e não está bloqueando.

```js
term.startRaw()
let k = term.readKey(false, 10)  # tenta por 10ms
if (k != null) { print(k) }
term.stopRaw()
```

### term.readKeyBlocking()
Bloqueia até uma tecla ser lida.

### term.pollKey()
Não bloqueia — retorna `null` se não houver tecla.

### term.keyPressed()
Retorna boolean indicando se há dados disponíveis sem consumir.

## Cursor e tela

### Alternate screen buffer (no scrollback)

Use `term.enterAltScreen()` to switch to the terminal's *alternate screen buffer* (CSI ?1049h). This is useful for full-screen UIs and drawing with `createBuffer`/`drawBuffer` since it prevents the terminal scrollback from containing intermediate frames. When finished, call `term.exitAltScreen()` to restore the main screen (CSI ?1049l).

```js
term.enterAltScreen()
# draw loop
term.exitAltScreen()
```

### term.clear()
Limpa a tela e move o cursor para (1,1).

### term.move(row, col)
Move o cursor para posição absoluta (1-based).

```js
term.clear()
term.move(5, 10)
print("Located")
```

### term.moveUp(n), term.moveDown(n), term.moveLeft(n), term.moveRight(n)
Movimentos relativos do cursor.

### term.hideCursor(), term.showCursor()
Esconde / mostra o cursor.

### term.writeAt(row, col, str)
Escreve `str` na posição absoluta.

### term.size()
Retorna `{rows, cols}`.

## Buffer de tela (in-memory)

### term.createBuffer(w, h)
Retorna um objeto buffer com campos `{w, h, rows}` e métodos:
- `buf.set(row, col, char)` — escreve `char` (first rune)
- `buf.clear()` — preenche com espaços
- `buf.fill(char)` — preenche com `char`
- `buf.copy(otherBuf)` — copia conteúdo (clip ao tamanho destino)

### term.drawBuffer(buf)
Desenha o buffer na tela (posição 1,1).

```js
let b = term.createBuffer(20,4)
b.fill(".")
b.set(2,5, "X")
term.drawBuffer(b)
```

## Estilos e cores (24-bit)

- `term.fg("#rrggbb")` ou `term.fg(r, g, b)`
- `term.bg(...)`
- `term.resetStyle()`
- `term.bold(bool)`, `term.underline(bool)`, `term.invert(bool)`

```js
term.fg("#ff0000")
print("Red")
term.resetStyle()
```

## Tempo e frames

- `term.sleep(ms)` — pausa ms milissegundos
- `term.now()` — timestamp em ms
- `term.frameLimit(fps)` — retorna função que regula a taxa de quadros

## Eventos

- `term.onResize(fn)` — registra callback para SIGWINCH; passe `null` para remover
- `term.onKey(fn)` — registra callback para tecla pressionada; passe `null` para remover

> Observação: callbacks são executados via `task.spawn` quando possível; garanta que o módulo `task` esteja disponível para run-time callbacks.

## Referência rápida

| Função | Descrição |
|---|---|
| `term.startRaw()` | Coloca stdin em modo raw |
| `term.stopRaw()` | Restaura modo anterior |
| `term.isRaw()` | Retorna se está em raw |
| `term.readKey(blocking, timeout_ms)` | Lê bytes brutos |
| `term.readKeyBlocking()` | Lê bloqueando |
| `term.pollKey()` | Lê sem bloquear |
| `term.keyPressed()` | Verifica disponibilidade |
| `term.clear()` | Limpa tela |
| `term.move(r,c)` | Move cursor absolute |
| `term.moveUp/Down/Left/Right(n)` | Move cursor rel.
| `term.hideCursor()`/`term.showCursor()` | Esconder/mostrar cursor |
| `term.writeAt(r,c,s)` | Escrever em posição |
| `term.size()` | Retorna {rows, cols} |
| `term.write/writeln/refresh/flush` | Buffer de saída |
| `term.fg/bg/resetStyle` | Cores 24-bit |
| `term.bold/underline/invert` | Atributos de texto |
| `term.createBuffer(w,h)` | Cria buffer em memória |
| `term.drawBuffer(buf)` | Desenha buffer no terminal |
| `term.sleep(ms)` | Pausa execução |
| `term.now()` | Timestamp em ms |
| `term.frameLimit(fps)` | Limita taxa de quadros |
| `term.onResize(fn)` | Callback em resize |
| `term.onKey(fn)` | Callback em tecla |

---

**Nota:** Use este módulo com cautela em scripts não interativos — muitas funções fazem sentido apenas em TTYs interativos.
