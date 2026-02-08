# Structs

## Definição

```js
struct Point {
    x
    y
}
```

## Com métodos

```js
struct Point {
    x
    y

    fn move(dx, dy) {
        this.x = this.x + dx
        this.y = this.y + dy
    }
}
```

## Instanciação

```js
let p = Point()
p.x = 20
p.y = 30
p.move(10, 20)

print(p.x)
print(p.y)
```

## Construtor 'init'

se existir o init ele é chamado automaticamente, se não existir os campos são inicializados com null

```js
struct Point {
    x
    y

    fn init(x, y) {
        this.x = x
        this.y = y
    }

    fn move(dx, dy) {
        this.x = this.x + dx
        this.y = this.y + dy
    }
}

let p = Point(20, 30)

p.move(10, 20)

print(p.x)
print(p.y)
```

## Valores padrão

```js
struct User {
    name = "anon"
    age = 0
}
```

