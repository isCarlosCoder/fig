# Módulo crypto

```js
use "crypto"
```

Funções de hash e codificação.

## Hashes

### crypto.hash(s)

Retorna o hash FNV-1a de 32 bits da string como número:

```js
let h = crypto.hash("hello")
print(h)  # número inteiro (ex: 1335831723)
```

### crypto.sha1(s)

Retorna o hash SHA-1 como string hexadecimal:

```js
print(crypto.sha1("hello"))
# aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
```

### crypto.sha256(s)

Retorna o hash SHA-256 como string hexadecimal:

```js
print(crypto.sha256("hello"))
# 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
```

## Base64

### crypto.base64Encode(s)

Codifica uma string em Base64:

```js
let encoded = crypto.base64Encode("Olá, mundo!")
print(encoded)  # T2zDoSwgbXVuZG8h
```

### crypto.base64Decode(s)

Decodifica uma string Base64:

```js
let decoded = crypto.base64Decode("T2zDoSwgbXVuZG8h")
print(decoded)  # Olá, mundo!
```

## Hexadecimal

### crypto.hexEncode(s)

Codifica uma string para hexadecimal:

```js
let hex = crypto.hexEncode("Fig")
print(hex)  # 466967
```

### crypto.hexDecode(s)

Decodifica uma string hexadecimal:

```js
let original = crypto.hexDecode("466967")
print(original)  # Fig
```

## Exemplo: Verificando integridade

```js
use "crypto"
use "io"

let conteudo = io.readFile("dados.txt")
let hash = crypto.sha256(conteudo)
print("SHA-256: " + hash)
```

## Referência rápida

| Função                     | Descrição                    |
|---------------------------|------------------------------|
| `crypto.hash(s)`           | Hash FNV-1a 32-bit (número) |
| `crypto.sha1(s)`           | Hash SHA-1 (hex string)     |
| `crypto.sha256(s)`         | Hash SHA-256 (hex string)   |
| `crypto.base64Encode(s)`   | Codificar em Base64         |
| `crypto.base64Decode(s)`   | Decodificar Base64          |
| `crypto.hexEncode(s)`      | Codificar em hexadecimal    |
| `crypto.hexDecode(s)`      | Decodificar hexadecimal     |
