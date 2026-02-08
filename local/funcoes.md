# Funções

## Função simples

```js
fn hello() {
    print("Hello")
}
```

## Função com parâmetros

```js
fn add(a, b) {
    return a + b
}
```

## Função com variáveis locais

```js
fn square(x) {
    let y = x * x
    return y
}

print(square(5));
```

## Função com condicional

```js
fn sign(n) {
    if (n > 0) {
        return 1;
    } elif (n < 0) {
        return -1;
    } else {
        return 0;
    }
}

print(sign(-3));
```

## Função com loop

```js
fn sumTo(n) {
    let s = 0;
    let i = 1;
    while (i <= n) {
        s = s + i;
        i++;
    }
    return s;
}

print(sumTo(5));
```

## Função recursiva

```js
fn fact(n) {
    if (n <= 1) {
        return 1;
    }
    return n * fact(n - 1);
}

print(fact(5));
```

## return antecipado
```js
fn findEven(n) {
    let i = 0;
    while (i <= n) {
        if (i % 2 == 0) {
            return i;
        }
        i++;
    }
    return -1;
}

print(findEven(7));
```

## Função sem return (retorna nil)

```js
fn log(x) {
    print(x);
}

log("teste");
```
