# Módulo task

```js
use "task"
```

Concorrência com goroutines. Permite executar funções em paralelo, aguardar resultados, definir timeouts e corridas (race).

## task.spawn(fn)

Cria uma nova tarefa que executa a função `fn` em paralelo. Retorna um **handle** (referência) da tarefa:

```js
use "task"

let t = task.spawn(fn() {
    return 42
})

print(task.await(t))  # 42
```

### Fire and forget

Se não precisar do resultado, basta chamar `spawn` sem guardar o handle:

```js
use "task"
use "system"

task.spawn(fn() {
    system.sleep(1000)
    print("acabou")
})

print("continuou")  # impresso imediatamente
```

## task.await(handle)

Bloqueia até a tarefa terminar e retorna o valor de retorno:

```js
use "task"

let t = task.spawn(fn() { return 21 * 2 })
let resultado = task.await(t)
print(resultado)  # 42
```

### Closures e funções

A função passada ao `spawn` pode usar closures e chamar outras funções:

```js
use "task"

fn work(x) {
    return x * 10
}

let t = task.spawn(fn() { return work(5) })
print(task.await(t))  # 50
```

## task.awaitTimeout(handle, ms)

Como `await`, mas com limite de tempo. Se a tarefa não terminar em `ms` milissegundos, gera um erro:

```js
use "task"
use "system"

let t = task.spawn(fn() {
    system.sleep(2000)
    return 99
})

let v = try task.awaitTimeout(t, 1000) onerror {
    return -1
}

print(v)  # -1 (timeout)
```

## task.race(handles)

Recebe um array de handles e retorna o resultado da **primeira** tarefa que terminar:

```js
use "task"
use "system"

let t1 = task.spawn(fn() { system.sleep(500); return "A" })
let t2 = task.spawn(fn() { system.sleep(300); return "B" })

print(task.race([t1, t2]))  # "B" (terminou primeiro)
```

## Execução paralela

Múltiplas tarefas executam simultaneamente. Três tarefas com 100ms cada levam ~100ms no total, não 300ms:

```js
use "task"
use "system"

fn work(x) {
    system.sleep(100)
    return x * 10
}

let t1 = task.spawn(fn() { return work(1) })
let t2 = task.spawn(fn() { return work(2) })
let t3 = task.spawn(fn() { return work(3) })

print(task.await(t1))  # 10
print(task.await(t2))  # 20
print(task.await(t3))  # 30
```

## Pipeline

Uma tarefa pode aguardar o resultado de outra:

```js
use "task"

let t1 = task.spawn(fn() { return 10 })
let t2 = task.spawn(fn() { return task.await(t1) * 3 })

print(task.await(t2))  # 30
```

## Tratamento de erros

Erros dentro de uma tarefa são propagados ao chamar `await`. Use `try/onerror` para capturá-los:

```js
use "task"
use "debug"

let t = task.spawn(fn() {
    debug.panic("falhou")
})

let v = try task.await(t) onerror(e) {
    print("erro: " + e)
    return 0
}

print(v)  # 0
```

## Uso com structs

### Métodos assíncronos

Métodos de structs podem retornar task handles. O padrão é copiar `this.campo` para uma variável local antes do `spawn` (closures não capturam `this`):

```js
use "task"
use "system"

struct Downloader {
    url

    fn init(url) {
        this.url = url
    }

    fn fetchAsync() {
        let u = this.url
        return task.spawn(fn() {
            system.sleep(100)
            return "dados de " + u
        })
    }
}

let d1 = Downloader("fig.dev/api")
let d2 = Downloader("fig.dev/docs")

let t1 = d1.fetchAsync()
let t2 = d2.fetchAsync()

print(task.await(t1))  # dados de fig.dev/api
print(task.await(t2))  # dados de fig.dev/docs
```

### Worker pool com structs

```js
use "task"
use "system"

struct Worker {
    name

    fn init(name) {
        this.name = name
    }

    fn runAsync(value) {
        let n = this.name
        let v = value
        return task.spawn(fn() {
            system.sleep(50)
            return n + " processou " + v
        })
    }
}

fn spawnWork(w, val) {
    return w.runAsync(val)
}

let w1 = Worker("alpha")
let w2 = Worker("beta")

let j1 = spawnWork(w1, "A")
let j2 = spawnWork(w2, "B")
let j3 = spawnWork(w1, "C")

print(task.await(j1))  # alpha processou A
print(task.await(j2))  # beta processou B
print(task.await(j3))  # alpha processou C
```

### Timeout e race com structs

```js
use "task"
use "system"

struct SlowService {
    delay

    fn init(delay) {
        this.delay = delay
    }

    fn callAsync() {
        let d = this.delay
        return task.spawn(fn() {
            system.sleep(d)
            return "pronto"
        })
    }
}

# timeout
let fast = SlowService(50)
let slow = SlowService(2000)

let r1 = try task.awaitTimeout(fast.callAsync(), 500) onerror {
    return "timeout"
}
print(r1)  # pronto

let r2 = try task.awaitTimeout(slow.callAsync(), 100) onerror {
    return "timeout"
}
print(r2)  # timeout

# race entre métodos
let s1 = SlowService(300)
let s2 = SlowService(50)
let winner = task.race([s1.callAsync(), s2.callAsync()])
print(winner)  # pronto (s2 termina primeiro)
```

## Resumo da API

| Função | Descrição |
|---|---|
| `task.spawn(fn)` | Executa `fn` em paralelo, retorna handle |
| `task.await(handle)` | Bloqueia até terminar, retorna resultado |
| `task.awaitTimeout(handle, ms)` | Await com limite de tempo |
| `task.race([handles])` | Retorna o primeiro resultado |
