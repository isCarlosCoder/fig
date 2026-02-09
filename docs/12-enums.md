# Enums (Enumerações) ⚑

Enums permitem declarar um conjunto nomeado de variantes com nome fixo. Elas são úteis para representar um grupo de valores mutuamente exclusivos (por exemplo, cores, estados, sinais).

## Sintaxe

```js
enum Color {
  Red
  Green
  Blue
}
```

- `enum` seguido do nome e um bloco com os membros (cada membro é um identificador em sua própria linha ou separados por espaço).
- Os membros são valores nomeados acessíveis como `Color.Red`, `Color.Green`, etc.

## Acesso

- Membros são acessados por ponto: `Color.Red`.
- Como `enum` é internamente representado como um objeto com chaves (membros), você também pode usar indexing por string: `Color["Red"]`.
- A ordem das chaves é preservada e `objects.keys(Color)` retorna a lista de nomes dos membros.

## Comportamento

- Cada membro é um valor distinto que pode ser comparado por igualdade com `==` e `!=`.
- Ao imprimir um membro, ele aparece como `EnumName.MemberName` (ex.: `Color.Red`).
- Enums funcionam bem com `match`/`switch` (i.e., você pode casar contra `Color.Red`).

## Exemplos

```js
enum Color { Red Green Blue }
print(Color.Red)              # Color.Red
let c = Color.Green
print(c == Color.Green)       # true
print(objects.keys(Color))    # ["Red", "Green", "Blue"]

match c {
  Color.Red => print("red")
  Color.Green => print("green")
  Color.Blue => print("blue")
}
```

## Regras e Observações

- Enum members são somente nomes; não aceitam inicializadores (por enquanto).
- Membros são únicos por enum (duplas com o mesmo nome não são permitidas na declaração).
- Enum names coexistem no mesmo espaço de nomes que structs, funções e variáveis (defina com cuidado para evitar colisões).

