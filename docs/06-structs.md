# Structs (Estruturas)

Structs permitem criar tipos personalizados com **campos** e **métodos**.

## Declarando uma struct

```js
struct Ponto {
    x = 0
    y = 0
}
```

Os campos podem ter **valores padrão**. Se nenhum valor for dado, o campo começa como `null`:

```js
struct Usuario {
    nome       # null por padrão
    idade      # null por padrão
    ativo = true
}
```

## Criando instâncias

Chame o nome da struct como uma função:

```js
let p = Ponto()
print(p.x)  # 0
print(p.y)  # 0
```

## Acessando e alterando campos

Use o operador `.` para ler e escrever campos:

```js
let u = Usuario()
u.nome = "Carlos"
u.idade = 30

print(u.nome)   # Carlos
print(u.idade)  # 30
print(u.ativo)  # true
```

## Método `init` (construtor)

O método `init` é chamado automaticamente ao criar uma instância. Use-o para inicializar campos com argumentos:

```js
struct Pessoa {
    nome
    idade

    fn init(nome, idade) {
        this.nome = nome
        this.idade = idade
    }
}

let p = Pessoa("Maria", 25)
print(p.nome)   # Maria
print(p.idade)  # 25
```

> **Nota:** Os argumentos passados na criação (ex: `Pessoa("Maria", 25)`) são repassados ao `init`.

## `this`

Dentro de métodos, `this` se refere à instância atual:

```js
struct Contador {
    valor = 0

    fn incrementar() {
        this.valor = this.valor + 1
    }

    fn obter() {
        return this.valor
    }
}

let c = Contador()
c.incrementar()
c.incrementar()
print(c.obter())  # 2
```

## Métodos

Métodos são funções declaradas dentro da struct com `fn`:

```js
struct Retangulo {
    largura = 0
    altura = 0

    fn init(l, a) {
        this.largura = l
        this.altura = a
    }

    fn area() {
        return this.largura * this.altura
    }

    fn perimetro() {
        return 2 * (this.largura + this.altura)
    }
}

let r = Retangulo(5, 3)
print(r.area())       # 15
print(r.perimetro())  # 16
```

## Métodos com parâmetros

```js
struct Ponto {
    x = 0
    y = 0

    fn init(x, y) {
        this.x = x
        this.y = y
    }

    fn mover(dx, dy) {
        this.x = this.x + dx
        this.y = this.y + dy
    }

    fn distancia(outro) {
        let dx = this.x - outro.x
        let dy = this.y - outro.y
        return math.sqrt(dx * dx + dy * dy)
    }
}

let a = Ponto(0, 0)
let b = Ponto(3, 4)

a.mover(1, 1)
print(a.x)  # 1
print(a.y)  # 1
```

> **Nota:** Para usar `math.sqrt`, é preciso importar o módulo com `use "math"`.

## Exemplo completo

```js
struct Tarefa {
    titulo
    concluida = false

    fn init(titulo) {
        this.titulo = titulo
    }

    fn concluir() {
        this.concluida = true
    }

    fn status() {
        if (this.concluida) {
            return "[x] " + this.titulo
        }
        return "[ ] " + this.titulo
    }
}

let t1 = Tarefa("Estudar Fig")
let t2 = Tarefa("Fazer café")

t1.concluir()

print(t1.status())  # [x] Estudar Fig
print(t2.status())  # [ ] Fazer café
```

## Resumo

| Conceito         | Sintaxe                                     |
|-----------------|---------------------------------------------|
| Declaração       | `struct Nome { campos e métodos }`          |
| Campo com padrão | `campo = valor`                             |
| Campo sem padrão | `campo` (inicia como `null`)                |
| Instanciar       | `Nome()` ou `Nome(args)`                    |
| Construtor       | `fn init(params) { ... }`                   |
| Método           | `fn nomeMetodo(params) { ... }`             |
| Acesso           | `instancia.campo`, `instancia.metodo()`     |
| Referência       | `this` (dentro de métodos)                  |
