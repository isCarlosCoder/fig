# Funções nativas — @native

O atributo `@native` marca funções simples que serão **compiladas para uma fast‑path nativa** (micro‑executor sobre float64) pelo runtime do Fig. O objetivo é reduzir o overhead de chamada/boxing em trechos numéricos críticos (por exemplo, `sigmoid` em loops de treino).

Sintaxe

- Formas aceitas:
  - `@native fn nome(...) { ... }`
  - `@native() fn nome(...) { ... }`  (parênteses opcionais)
  - `@[native()] fn nome(...) { ... }` (compatibilidade)

Exemplo:

```js
@native fn sigmoid(x) {
  return 1 / (1 + math.exp(-x))
}

print(sigmoid(0))  # 0.5
```

Semântica e comportamento

- Ao encontrar `@native` a linguagem tenta validar e compilar o corpo da função para uma implementação nativa (_native fast‑path_). Se a validação falhar, por padrão a definição gera erro em tempo de definição.
- A implementação nativa substitui a execução via visitor/AST para chamadas dessa função, reduzindo alocação e custo por chamada.
- Os metadados (argumentos opcionais em parênteses) são aceitos e armazenados; ex.: `@native(fallback=true)` (veja opções abaixo).

Restrições iniciais (fase 1)

- Assinatura: parâmetros e retorno devem ser `number` (não são aceitos arrays/objects como parâmetros de fast‑path).
- Corpo: obrigatoriamente **uma única instrução `return <expr>`**.
- Expressões permitidas (inicialmente): literais numéricos, referências a parâmetros, operadores aritméticos (`+ - * / %`), unário `-`, parênteses, e chamadas a um conjunto seguro de funções `math.*` (ex.: `math.exp`, `math.sin`, `math.cos`, `math.sqrt`, `math.log`, `math.abs`, `math.pow`).
- É proibido: declarações locais (`let`), loops, closures, acesso a objetos/arrays, `this`, chamadas de I/O, ou qualquer construção que cause efeitos colaterais.

Opções (meta‑argumentos)

- fallback=true — tenta compilar; se a compilação falhar, mantém a função interpretada (fallback automático). Por padrão (sem `fallback`) falha em tempo de definição.
- (planejado) vectorized=true — sinalizar que a função aceita arrays e será aplicada element‑wise (futuro).

Mensagens de erro comuns

- "native function must contain exactly one statement (a return)" — corpo inválido
- "native function validation failed: unsupported ..." — expressão contém construção não suportada
- "native: argument N is not a number" — chamada passando tipos não numéricos

Boas práticas

- Use `@native` para funções pequenas, puras e frequentemente chamadas (ex.: funções matemáticas usadas em loops). Evite para código que precisa acessar variáveis externas ou manipular estruturas complexas.
- Prefira `@native(fallback=true)` durante migração para evitar quebra imediata quando a função não for compilável.
- Meça desempenho antes/depois com benchmarks — `@native` reduz overhead de chamada, mas o maior ganho vem junto de operações vetoriais nativas ou tipos numéricos especializados.

Exemplos

1) Função pura simples — compilada com sucesso:

```js
@native fn dsigmoid(x) { return x * (1 - x) }
```

2) Migração segura (uso de fallback):

```js
@native(fallback=true) fn maybe_native(x) { return 2 * x }
```

Quando usar

- Cenários compute‑bound com milhões de chamadas curtas (ex.: treino, simulações numéricas).  
- Funções puras que realizam apenas matemática sobre números.

Limitações e roadmap

- Fase inicial suporta apenas expressões escalares numéricas. Próximas etapas planejadas: suporte a vetorização, tipos `FloatArray`/`Matrix` e integração com BLAS para operações matriciais grandes.

Segurança

- O código compilado para a fast‑path roda no runtime controlado do Fig (sem execuções arbitrárias de sistema). O `@native` não habilita acesso a syscalls nem contorna sandboxing.

Referências

- Ver também: [Funções — seção @native](docs/05-funcoes.md)
- Implementação: interpreter/native_compile.go
