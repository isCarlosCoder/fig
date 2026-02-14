# Módulo mathx

```js
use "mathx"
```

Extensões numéricas estilo NumPy para operações vetoriais, álgebra linear, transformadas e utilitários para arrays.
As funções do `mathx` normalmente aceitam escalares ou arrays (opera elemento-a-elemento e reconstroi shapes; broadcasting suportado em casos simples).

## Operações vetoriais (elementwise)

Funções aceitam `number` ou `array` e retornam um `number` (quando todos os argumentos forem escalares) ou `array` com a mesma forma de entrada.

```js
print(mathx.add([1,2,3], 1))    # [2, 3, 4]
print(mathx.multiply([1,2], [3,4]))  # [3, 8]
print(mathx.clip([0,-1,10], 0, 5))    # [0, 0, 5]
```

Principais funções: `add`, `subtract`, `multiply`, `divide`, `floor_divide`, `power`, `mod`, `remainder`, `negative`, `absolute`, `sqrt`, `square`, `reciprocal`, `clip`, `maximum`, `minimum`, `sign`.

---

## Comparações e operadores lógicos

Retornam arrays/bools elementwise; aceitam scalars/arrays.

```js
print(mathx.equal([1,2],[1,3]))   # [true, false]
print(mathx.logical_and([true,false],[true,true]))  # [true, false]
print(mathx.isclose([1.0, 1.000001], 1.0))  # [true, true]
```

Funções: `equal`, `not_equal`, `greater`, `greater_equal`, `less`, `less_equal`, `logical_and`, `logical_or`, `logical_not`, `logical_xor`, `all`, `any`, `isfinite`, `isinf`, `isnan`, `isclose`.

---

## Ordenação e operações sobre conjuntos

- `sort`, `argsort`, `lexsort`, `partition`, `argpartition`
- `unique`, `setdiff1d`, `intersect1d`, `union1d`, `in1d`

Exemplo:

```js
print(mathx.sort([3,1,2]))      # [1,2,3]
print(mathx.unique([1,2,2,3]))   # [1,2,3]
```

---

## Amostragem / aleatoriedade (np.random-like, limitada)

APIs para gerar números/arrays aleatórios. Se `size`/shape for informado, retorna `array` com essa forma.

```js
print(mathx.rand())              # scalar ~ U(0,1)
print(mathx.randn(3))            # array de 3 draws ~ N(0,1)
print(mathx.randint(0, 10))      # inteiro em [0,10)
print(mathx.choice(["a","b","c"]))
```

Distribuições: `rand`, `random` (alias), `randn`, `randint`, `choice`, `shuffle`, `permutation`, `normal`, `uniform`, `binomial`, `poisson`, `exponential`, `gamma`, `beta`.

---

## Tipos e conversões (dtype)

- `dtype(value)` — retorna `"number"|"string"|"boolean"|"array"|"object"|"mixed"|"null"`
- `astype(value, dtype)` — converte valores/arrays para `number|string|boolean`
- `issubdtype(a, b)`, `can_cast(a, b)`, `result_type(...)`

```js
print(mathx.dtype([1, "x"]))   # "mixed"
print(mathx.astype(["1", null], "number"))  # [1, 0]
print(mathx.can_cast("123", "number"))      # true
```

---

## Álgebra linear (limitada)

Operações matriciais e decomposições (implementação reduzida — útil para tamanhos pequenos/tests):

- `matmul`, `dot`, `inner`, `tensordot`
- `inv`, `pinv`, `solve`
- `eig`, `eigvals` (2×2), `svd` (1×1 e 2×2), `qr`, `cholesky`, `matrix_rank`

Exemplo:

```js
let A = [[1,2],[3,4]]
let B = [[2,0],[1,2]]
print(mathx.matmul(A, B))
```

---

## FFT / sinais

Transformadas discretas 1D/2D (implementação simples):

- `fft`, `ifft`, `rfft`, `irfft`
- `fft2`, `ifft2`, `fftn`, `ifftn`
- `fftshift`, `ifftshift`

```js
print(mathx.fft([1,0,0,0]))  # espectro (pares [re,im])
```

---

## Entrada / saída (I/O)

Serialização rápida e utilitários para arquivos:

- `save(path, value)` — grava JSON legível
- `load(path)` — tenta decodificar JSON; caso contrário retorna o conteúdo como string
- `savez(path, object | name, val, ...)` — grava múltiplos arrays/valores em JSON
- `savetxt` / `loadtxt` / `genfromtxt` — CSV / texto simplificado
- `tofile` / `fromfile` — grava/ler binário float64 little-endian

```js
mathx.save("arr.json", [1,2,3])
let obj = mathx.load("arr.json")
```

---

## Polinômios

Coeficientes em ordem *highest-first* (como NumPy).

- `polyval(coeffs, x)` — avalia polinômio (x pode ser número ou array)
- `polyfit(x, y, deg)` — ajuste por mínimos quadrados (retorna coeficientes)
- `polyadd`, `polysub`, `polymul`, `polyder`, `polyint`

```js
print(mathx.polyval([1, 2, 3], 2))   # 1*2^2 + 2*2 + 3 = 11
print(mathx.polyfit([0,1,2], [3,6,11], 2))  # aproxima [1,2,3]
```

---

## Utilitários

- `ndim(value)` — número de dimensões (0 para scalars)
- `size(value)` — número total de elementos (scalars→1)
- `itemsize(value)` — tamanho aproximado em bytes por item (number=8, boolean=1, string=len)
- `copyto(dest, src)` — copia `src` para `dest` (in-place)
- `view(array)` — retorna *view* (compartilha backing array)
- `get_printoptions()` / `set_printoptions(obj)` — ajustar `precision`, `linewidth`, `threshold` para representação de arrays

```js
print(mathx.ndim([[1],[2]]))   # 2
print(mathx.size([1,2,3]))      # 3
mathx.set_printoptions({precision: 3})
print(mathx.get_printoptions())
```

---

## Referência rápida (completa)

Abaixo está uma referência agrupada das funções mais utilizadas no `mathx`. Use `use "arrays"` para regras de criação/indexação de arrays.

### Criação & formas

| Função | Descrição |
|---|---|
| `array`, `asarray`, `asanyarray` | Construção/normalização de arrays |
| `copy`, `view`, `copyto` | Cópia / view / escrita in-place |
| `zeros`, `ones`, `empty`, `full` (+ `_like`) | Inicializadores |
| `arange`, `linspace`, `logspace`, `geomspace` | Espaçamentos e sequências |
| `eye`, `identity`, `diag`, `diagflat` | Matrizes identidade/diagonais |
| `fromfunction`, `fromiter`, `frombuffer` | Construção a partir de funções/iteráveis/buffer |
| `shape`, `reshape`, `ravel`, `flatten`, `transpose`, `swapaxes`, `moveaxis`, `expand_dims`, `squeeze` | Manipulação de forma |
| `broadcast_to`, `broadcast_arrays`, `tile`, `repeat`, `concatenate`, `stack`, `vstack`, `hstack`, `dstack`, `column_stack`, `row_stack` | Operações de empilhamento/concatenação |
| `split`, `array_split`, `hsplit`, `vsplit`, `dsplit` | Divisão de arrays |

### Indexação / seleção

| Função | Descrição |
|---|---|
| `take`, `put` | Acesso por índice / atribuição |
| `where`, `nonzero`, `argwhere`, `extract`, `select`, `choose`, `compress` | Seleção condicional / máscaras |

### Operações elementwise (aritmética)

| Função | Descrição |
|---|---|
| `add`, `subtract`, `multiply`, `divide`, `true_divide`, `floor_divide` | Operações aritméticas básicas |
| `power`, `mod`, `remainder` | Potência / módulo |
| `negative`, `absolute`, `fabs`, `sqrt`, `square`, `reciprocal`, `clip` | Unários e clamp |
| `maximum`, `minimum`, `sign` | Reduções elementwise auxiliares |

### Comparações e lógicas

| Função | Descrição |
|---|---|
| `equal`, `not_equal`, `greater`, `greater_equal`, `less`, `less_equal` | Comparadores elementwise |
| `logical_and`, `logical_or`, `logical_xor`, `logical_not` | Operadores lógicos |
| `all`, `any`, `isfinite`, `isinf`, `isnan`, `isclose` | Predicados e agregações lógicas |

### Ordenação e operações com conjuntos

| Função | Descrição |
|---|---|
| `sort`, `argsort`, `lexsort`, `partition`, `argpartition` | Ordenação / índices |
| `unique`, `setdiff1d`, `intersect1d`, `union1d`, `in1d` | Operações de conjunto / unicidade |

### Estatística / agregação

| Função | Descrição |
|---|---|
| `sum`, `mean`, `average`, `var`, `std`, `median` | Estatísticas básicas |
| `percentile`, `quantile`, `min`, `max`, `ptp` | Estatísticas descritivas |
| `argmin`, `argmax`, `cumsum`, `cumprod`, `histogram`, `bincount` | Indexação cumulativa / histogramas |
| `cov`, `corrcoef` | Covariância / correlação |

### Aleatoriedade / amostragem

| Função | Descrição |
|---|---|
| `rand`, `random`, `randn`, `randint`, `choice` | Geração aleatória básica |
| `shuffle`, `permutation` | Permutações / embaralhar |
| `normal`, `uniform`, `binomial`, `poisson`, `exponential`, `gamma`, `beta` | Distribuições estatísticas |

### Tipos & conversão

| Função | Descrição |
|---|---|
| `dtype`, `astype`, `issubdtype`, `can_cast`, `result_type` | Inspeção e coerção de tipos |

### Álgebra linear (limitado)

| Função | Descrição |
|---|---|
| `dot`, `inner`, `outer`, `vdot`, `tensordot` | Operações vetoriais/matriciais |
| `matmul` | Produto matricial |
| `trace`, `norm`, `det`, `inv`, `pinv`, `solve` | Operações matriciais |
| `eig`, `eigvals`, `svd`, `qr`, `cholesky`, `matrix_rank` | Decomposições (casos limitados) |

### FFT / sinal

| Função | Descrição |
|---|---|
| `fft`, `ifft`, `rfft`, `irfft` | Transformadas 1D |
| `fft2`, `ifft2`, `fftn`, `ifftn` | Transformadas 2D / N-D (limitado) |
| `fftshift`, `ifftshift` | Rearranjo de espectro |

### I/O (arquivos)

| Função | Descrição |
|---|---|
| `save`, `load`, `savez` | JSON / múltiplos itens |
| `savetxt`, `loadtxt`, `genfromtxt` | CSV / texto simples |
| `tofile`, `fromfile` | Binário float64 (little-endian) |

### Polinômios

| Função | Descrição |
|---|---|
| `polyval`, `polyfit` | Avaliar / ajustar polinômios |
| `polyadd`, `polysub`, `polymul` | Soma / subtração / multiplicação |
| `polyder`, `polyint` | Derivada / integral simbólica (coeficientes) |

### Utilitários

| Função | Descrição |
|---|---|
| `ndim`, `size`, `itemsize` | Informações de forma/tamanho |
| `get_printoptions`, `set_printoptions` | Configurar representação de arrays |


---

Se quiser, eu adiciono links diretos para exemplos em outras seções do documento ou crio uma tabela de referência exportável (CSV/HTML) para a documentação do site.