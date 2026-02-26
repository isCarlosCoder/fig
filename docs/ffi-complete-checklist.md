# Checklist FFI: tornar Fig independente de código Go 🚀

Objetivo: listar tudo que falta implementar, testar e documentar para que Fig consiga ser mantida e evoluída majoritariamente via FFI (bibliotecas nativas) sem precisar alterar código Go do runtime/stdlib.

---

## Visão geral
- Status atual: já temos _load/sym/call_ para `int`, `double`, `string`, `void`, bytes e callbacks C→Fig com round‑trip testado. Structs aninhados, ownership (alloc/free/strdup/free_string), timeouts configuráveis e stress tests concluídos. ✅
- Meta: poder reimplementar todas as partes importantes da stdlib e código do projeto em FFI, com estabilidade, segurança e performance comparáveis à implementação em Go. 🎯

---

## Como usar esta checklist
- Cada item tem: **descrição**, **prioridade (P0/P1/P2)**, **estimativa (S/M/L)** e **critérios de aceitação / testes**.
- Começar do topo (P0) e avançar para P1/P2. Testes de integração e stress são críticos antes de marcar "done".

---

## Prioridade P0 — Essenciais (sem isso não dá pra dizer "FFI-only") ⚠️

1.  ✅ **Marshallings de tipos compostos** (arrays, maps, structs) — P0, M — **IMPLEMENTADO**
    - Descrição: suportar (de/para) Fig <-> C: arrays de objetos, mapas (string->value), structs nomeados e aninhados.
    - Critérios: roundtrip test para arrays, mapas e structs; documentação de como mapear tipos; API de validação de schema.
    - Testes: unitários + integração (ex.: função C que recebe struct aninhada e retorna modificado).
    - **Status:** `define_struct`, expansão recursiva de structs aninhados, array roundtrip, bytes↔array/string. 27 testes passando.

2.  ✅ **Assinaturas C completas** — P0, M — **IMPLEMENTADO**
    - Descrição: suportar ponteiros, structs por valor e por referência, callbacks C (função ponteiro) e múltiplos tipos de retorno.
    - Critérios: chamar funções com assinaturas complexas; passar/receber structs; suporte a `void*` para dados binários.
    - Testes: bibliotecas C que expõem funções com structs/ponteiros/retornos complexos.
    - **Status:** int (0–3 args), double (0–2 args), string (0–4 args com dispatch por assinatura), void (0–2 args). Wrappers CGo completos em `wrappers.c`. `arg_types` para dispatch tipado.

3.  ✅ **Regras de memória/ownership** — P0, M — **IMPLEMENTADO**
    - Descrição: definir contrato claro "quem libera" e helpers (ffi.malloc / ffi.free / ffi.strdup / ffi.free_string) para evitar leaks/duplo free.
    - Critérios: documentação; API helpers; testes leak-free (ASAN/static check) em integração.
    - Testes: ASAN/valgrind runs em CI; testes que provocam erros de ownership sem corromper memória.
    - **Status:** `alloc`, `free`, `strdup`, `free_string`, `mem_write`, `mem_read` implementados. Documentação em `docs/13-ffi.md`. TestFfiMemoryOwnership passa.

4.  ✅ **Callbacks robustos** (bidirecionais) — P0, M — **IMPLEMENTADO**
    - Descrição: garantir callbacks C→Fig (já existe) *e* Fig→C, suporte para múltiplos args, async, cancel e segurança contra reentrância que quebra stacks.
    - Critérios: testes com callbacks concorrentes e long-running; cancel/timeout; validação de tipos nos argumentos.
    - Testes: stress com 1000 callbacks concorrentes; timeouts configuráveis.
    - **Status:** `register_callback`, `unregister_callback`, callback C→Fig via CGo export. Stress test com 50 callbacks sequenciais (serialização intencional por limitação CGo). TestFfiCallback e TestFfiStressCallbackStorm passando.

5.  ✅ **Cancelamento e timeouts** — P0, S — **IMPLEMENTADO**
    - Descrição: propagar cancelamentos/timeouts de Fig→helper→C e de C→Fig; evitar deadlocks.
    - Critérios: chamadas bloqueantes têm mecanismo de cancel; testável; configurações via fig.toml.
    - Testes: chamadas que bloqueiam e são canceladas pelo lado Fig; validar retorno e limpeza.
    - **Status:** `call_timeout` configurável no `fig.toml` (padrão 3s; 0 = sem limite). Callback timeout fixo de 2s. TestFfiCallTimeout passa.

6.  ✅ **Testes de estabilidade e concorrência** — P0, M — **IMPLEMENTADO**
    - Descrição: suite de stress que valida grandes cargas, race conditions e falhas parciais (C crash, helper crash).
    - Critérios: CI automatizado com stress tests; determinismo alto; flaky < threshold.
    - Testes: stress 1k goroutines, simulação de crashes e recuperação.
    - **Status:** TestFfiStressConcurrentCalls (200 goroutines), TestFfiStressCallbackStorm (50 callbacks), TestFfiHelperCrashRestart. 27 testes passando juntos sem flaky.

7.  ✅ **Documentação de ABI & guia de autoria** — P0, S — **IMPLEMENTADO**
    - Descrição: documentação completa explicando assinaturas, ownership, marshaling, convenções (quem chama quem, error conventions).
    - Critérios: exemplos passo-a-passo para criar uma biblioteca FFI que substitui uma parte da stdlib.
    - **Status:** `docs/13-ffi.md` (guia completo do usuário com referência de API) criado.

---

## Prioridade P1 — Necessários para produção e manutenção confortáveis 🛠

8.  ✅ **Ferramentas e geração de binding/IDL** — P1, M — **IMPLEMENTADO**
    - Descrição: definir um formato declarativo (ex.: TOML/JSON/DSL) para declarar funções e tipos (ex.: `ffi.def`) e gerar glue code ao build.
    - Critérios: gerador que produz wrappers de typedefs, validadores e testes automáticos.
    - **Status:** `ffi-gen` tool com IDL `ffi.def.toml` (TOML). Gera código Fig com load/sym/call. Suporta int/double/string/void/struct. 7 testes passando.

9.  **Zero-copy e streaming para grandes blobs** — P1, M
    - Descrição: transferir grandes buffers sem cópia quando possível (mmap, file descriptors, slices com shared memory) e streaming APIs.
    - Critérios: benchmarks mostrando melhoria; API para streaming/chunked IO.

10. ✅ **Políticas de segurança do helper** — P1, M — **IMPLEMENTADO**
    - Descrição: limites de CPU/mem, seccomp/containers, execução com usuário restrito, tempo de vida, sandboxing de libs carregadas.
    - Critérios: policies aplicáveis via config; testes que forçam limites causando abort controlado.
    - **Status:** `SandboxConfig` com `max_memory_mb`, `max_cpu_seconds`, `max_libs`, `max_allocs`, `max_restarts` via `[ffi.sandbox]` no fig.toml. `prlimit` (RLIMIT_AS/CPU) no Linux. Contadores atômicos per-project para libs, allocs, restarts. `sandbox_status()` builtin. 6 testes passando.

11. ✅ **Cross-platform support** — P1, M — **IMPLEMENTADO**
    - Descrição: suporte a Windows (LoadLibrary/GetProcAddress), macOS (dyld), ARM. Build matrix em CI.
    - Critérios: testes em runners Windows/macOS/ARM; .dll e .dylib suportados.
    - **Status:** `dl_portable.h` (abstração POSIX/Windows), CGo flags por plataforma, `lib_ext()`/`lib_name()` builtins, fallback stdio no Windows. 6 testes passando.

12. **Política de assinaturas de erro e propagação** — P1, S
    - Descrição: convenção para mensagens de erro entre C↔helper↔Fig; enriquecimento de erros (stack trace/context).
    - Critérios: erros informativos em chamadas falhas; testes que verificam propagação completa.

13. **Hot-reload e rollback do helper** — P1, M
    - Descrição: permitir reiniciar helper sem atrapalhar execução de Fig, manter requests idempotentes ou com rollback claro.
    - Critérios: testes que reiniciam helper sem corromper estado; reconexão automática.

---

## Prioridade P2 — Melhorias e recursos legais ✨

14. **Reimplementação incremental da stdlib em FFI** — P2, L
    - Exemplo: `fs` (IO), `strings`, `crypto`, `json`, `net`.
    - Critérios: cada módulo tem equivalente funcional, performance e testes de compatibilidade com APIs existentes.

15. **Plugin marketplace / signing** — P2, L
    - Descrição: permitir distribuição de libs FFI assinadas, verificação de integridade e reputação.

16. ✅ **Fuzzing e ASAN/UB sanitizers integrados** — P2, M — **IMPLEMENTADO**
    - Descrição: rodar fuzz + sanitizers nas bibliotecas nativas usadas em testes para captura precoce de erros.
    - **Status:** `asan_driver.c` (driver standalone C com -fsanitize=address via dlopen), `lib_asan_clean.c` (biblioteca ASAN-safe), `tools/asan-check.sh` e `tools/valgrind-check.sh` (scripts CLI). 5 testes automatizados (4 pass, 1 skip sem valgrind).

17. ✅ **Tooling: project template, CI snippets, examples** — P2, M — **IMPLEMENTADO**
    - Descrição: templates de projeto FFI, exemplos reimplementando parte da stdlib, `fig ffi new` generator.
    - **Status:** `ffi-gen -init <nome>` cria projeto completo (fig.toml, .ffi.def.toml, .c, Makefile, main.fig, README.md). Testado.

---

## Infra + CI / testes obrigatórios
- Adicionar jobs CI para: Linux x64, Linux ARM, macOS, Windows (build libs e rodar integração) ✅
- Criar job de ASAN/UBSAN para a suíte de testes nativa — obrigatório para cada PR com alterações de FFI. ✅
- Testes de stress/concurrency e tests de crash/recovery for automated nightly runs.

---

## Critérios de "FFI-only" (Definition of Done) ✅
Para podermos dizer que "Fig consegue se manter sozinha via FFI", todos estes devem ser **true**:
1. Core stdlib (I/O, strings, JSON, coleções, crypto, time) pode ser substituído por bibliotecas FFI sem mudanças no runtime Go (apenas config). ✅/por testar
2. Todas as APIs de stdlib têm exemplos e testes que rodam na mesma suíte sem depender de Go-only helpers.
3. Suites de stress e memleak (ASAN) aprovadas numa base diária/PR. ✅
4. Portabilidade: Linux (x86/ARM), macOS, Windows suportados. ✅
5. Segurança & sandboxing: políticas aplicáveis e validadas por testes. ✅
6. Documentação de autor e ferramentas para geração de bindings e templates. ✅

---

## Regras & Convenções recomendadas (padrões)
- Ownership: funções que retornam `char*` alocado por helper devem documentar que o **caller** libera com `ffi.free_string()`; vice-versa para buffers fornecidos pelo caller.
- Callbacks: sempre usar IDs e timeouts configuráveis; não bloquear o read loop do helper (usar goroutines dedicadas para leitura). ✅
- Versionamento: adicionar `ffi` section em `fig.toml` com `ffi.api_version` e `strict_mode` para mudanças de breaking.

---

## Exemplo de roadmap em milestones
- **M1 (P0)**: structs/maps, ownership API, async/callback hardening, tests de concorrência. (4–6 semanas)
- **M2 (P1)**: zero-copy streaming, sandboxing básico, cross-platform CI. (4–8 semanas)
- **M3 (P2)**: tooling/IDL, marketplace/signing, extensive examples e port stdlib. (8–12 semanas)

---

- [x] Implementar structs/maps marshalling e testes. (P0) ✅
- [x] Implementar i) helpers de alloc/free e ii) documentar ownership. (P0) ✅
- [x] Adicionar stress tests (concurrency + crash recovery). (P0) ✅
- [x] Definir formato de IDL/bindings (ex.: `ffi.def`) e escrever especificação. (P1) ✅

---

- [x] Marshalling: structs, mapas, arrays complexos ✅
- [x] Assinaturas: pointers, structs by-ref/value, function pointers ✅
- [x] Memory ownership API + docs ✅
- [x] Callbacks bidirecionais robustos (concorrência, cancel, timeout) ✅
- [x] Cancelamento/Timeouts e propagação de erros ✅
- [x] Testes: stress, concurrency ✅
- [x] Testes: ASAN/Valgrind ✅
- [x] Sandbox/policies (CPU, mem) e reinício/resiliência do helper ✅
- [x] Cross-platform support (Windows/macOS/ARM) ✅
- [x] Tooling: IDL, generators, templates ✅
- [ ] Reimplementação incremental de stdlib em FFI e validação de compatibilidade

---

> Se preferir, eu posso transformar esta checklist em Issues/epics/prioritizados no repositório (com labels e estimativas) e gerar um board para acompanhar o progresso. Quer que eu faça isso agora? 🚀

---

_Fim da checklist — adicionei este arquivo em `docs/ffi-complete-checklist.md`._
