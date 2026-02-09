# Módulo date

```js
use "date"
```

Funções para manipulação de datas e horários, baseadas em timestamps Unix em milissegundos.

## Funções

### date.now()

Retorna o timestamp atual em milissegundos (Unix epoch):

```js
let agora = date.now()
print(agora)  # ex: 1770620395048
```

### date.format(timestamp, formato)

Formata um timestamp em string usando tokens de formato:

```js
let agora = date.now()
print(date.format(agora, "YYYY-MM-DD"))         # 2026-02-09
print(date.format(agora, "DD/MM/YYYY"))          # 09/02/2026
print(date.format(agora, "YYYY-MM-DD HH:mm:ss")) # 2026-02-09 03:59:55
print(date.format(agora, "HH:mm"))               # 03:59
```

**Tokens de formato:**

| Token | Descrição       | Exemplo |
|-------|----------------|---------|
| `YYYY`| Ano (4 dígitos)| 2026    |
| `MM`  | Mês (01–12)    | 02      |
| `DD`  | Dia (01–31)    | 09      |
| `HH`  | Hora (00–23)   | 15      |
| `mm`  | Minuto (00–59) | 30      |
| `ss`  | Segundo (00–59)| 45      |

### date.parse(string, formato)

Converte uma string de data em timestamp (milissegundos):

```js
let natal = date.parse("2024-12-25", "YYYY-MM-DD")
print(natal)  # 1735095600000

let preciso = date.parse("2024-06-15 10:30:45", "YYYY-MM-DD HH:mm:ss")
print(date.format(preciso, "DD/MM/YYYY"))  # 15/06/2024
```

### date.add(timestamp, quantidade, unidade)

Adiciona uma quantidade de tempo ao timestamp:

```js
let agora = date.now()

let amanha = date.add(agora, 1, "day")
print(date.format(amanha, "YYYY-MM-DD"))

let daqui_2h = date.add(agora, 2, "hour")
print(date.format(daqui_2h, "HH:mm:ss"))

let semana_que_vem = date.add(agora, 1, "week")
print(date.format(semana_que_vem, "YYYY-MM-DD"))
```

**Unidades aceitas:**

| Unidade               | Aliases                  |
|----------------------|--------------------------|
| Milissegundos        | `ms`, `millisecond`, `milliseconds` |
| Segundos             | `s`, `second`, `seconds` |
| Minutos              | `min`, `minute`, `minutes` |
| Horas                | `h`, `hour`, `hours`     |
| Dias                 | `day`, `days`            |
| Semanas              | `week`, `weeks`          |

### date.diff(ts1, ts2, unidade)

Calcula a diferença entre dois timestamps na unidade especificada:

```js
let inicio = date.parse("2024-01-01", "YYYY-MM-DD")
let fim = date.parse("2024-12-31", "YYYY-MM-DD")

print(date.diff(inicio, fim, "day"))   # 365
print(date.diff(inicio, fim, "hour"))  # 8760

let ts1 = date.parse("2024-01-01 10:00:00", "YYYY-MM-DD HH:mm:ss")
let ts2 = date.parse("2024-01-01 15:30:00", "YYYY-MM-DD HH:mm:ss")
print(date.diff(ts1, ts2, "hour"))     # 5.5
```

> O resultado é negativo quando `ts2` é anterior a `ts1`.

### date.from_timestamp(timestamp)

Decompõe um timestamp em um objeto com os campos individuais:

```js
let agora = date.now()
let dt = date.from_timestamp(agora)

print(dt.ano)      # 2026
print(dt.mes)      # 2
print(dt.dia)      # 9
print(dt.hora)     # 3
print(dt.minuto)   # 59
print(dt.segundo)  # 55
```

**Campos do objeto retornado:**

| Campo     | Tipo   | Descrição             |
|-----------|--------|-----------------------|
| `ano`     | number | Ano (ex: 2026)        |
| `mes`     | number | Mês (1–12)            |
| `dia`     | number | Dia (1–31)            |
| `hora`    | number | Hora (0–23)           |
| `minuto`  | number | Minuto (0–59)         |
| `segundo` | number | Segundo (0–59)        |

## Exemplo completo

```js
use "date"

let agora = date.now()
let formatado = date.format(agora, "DD/MM/YYYY HH:mm:ss")
print("Agora: " + formatado)

let natal = date.parse("2024-12-25 00:00:00", "YYYY-MM-DD HH:mm:ss")
let dias_ate_natal = date.diff(agora, natal, "day")
print("Dias ate o natal: " + dias_ate_natal)

let proximo_mes = date.add(agora, 30, "day")
print("Daqui 30 dias: " + date.format(proximo_mes, "YYYY-MM-DD"))

let dt = date.from_timestamp(agora)
if (dt.mes == 12) {
    print("Estamos em dezembro!")
}
```

## Referência rápida

| Função                          | Descrição                                    |
|--------------------------------|----------------------------------------------|
| `date.now()`                    | Timestamp atual em milissegundos              |
| `date.format(ts, fmt)`          | Formatar timestamp para string                |
| `date.parse(str, fmt)`          | Converter string para timestamp               |
| `date.add(ts, qtd, unidade)`    | Adicionar tempo ao timestamp                  |
| `date.diff(ts1, ts2, unidade)`  | Diferença entre dois timestamps               |
| `date.from_timestamp(ts)`       | Decompor timestamp em objeto com campos       |
