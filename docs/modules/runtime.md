# Módulo runtime

```js
use "runtime"
```

Informações sobre o runtime Go da linguagem Fig.

## runtime.gc()

Força uma coleta de lixo (garbage collection):

```js
runtime.gc()
```

## runtime.memUsage()

Retorna um objeto com estatísticas de uso de memória (em bytes):

```js
let mem = runtime.memUsage()
print(mem.alloc)       # memória alocada (em uso)
print(mem.totalAlloc)  # total já alocado
print(mem.sys)         # memória obtida do OS
print(mem.numGC)       # número de coletas de lixo
```

## runtime.version()

Retorna a versão do Go usada para compilar o Fig:

```js
print(runtime.version())  # ex: go1.22.0
```

## runtime.platform()

Retorna a plataforma no formato `"SO/arquitetura"`:

```js
print(runtime.platform())  # ex: linux/amd64
```

## runtime.numCPU()

Retorna o número de CPUs lógicas disponíveis:

```js
print(runtime.numCPU())  # ex: 8
```

## Exemplo: Monitoramento de memória

```js
use "runtime"

let antes = runtime.memUsage()

# ... alguma operação que usa memória ...
let dados = []
for i in range(0, 10000) {
    arrays.push(dados, i)
}

let depois = runtime.memUsage()
print("Memória usada: " + (depois.alloc - antes.alloc) + " bytes")
print("Coletas de lixo: " + depois.numGC)
```

## Referência rápida

| Função              | Descrição                            |
|--------------------|--------------------------------------|
| `runtime.gc()`      | Forçar coleta de lixo                |
| `runtime.memUsage()`| Estatísticas de memória              |
| `runtime.version()` | Versão do Go                         |
| `runtime.platform()`| SO e arquitetura                     |
| `runtime.numCPU()`  | Número de CPUs lógicas               |
