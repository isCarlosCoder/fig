# FFI — Instalação do `ffi-helper`

Este documento descreve como instalar e tornar acessível o binário `ffi-helper` — o processo auxiliar usado pelo sistema FFI do Fig.

Por padrão o comando `fig setup-ffi` tenta detectar um `ffi-helper` global (no `PATH`). Se não encontrar, ele configura um helper local em `.fig/ffi/ffi-helper` e explica como compilar/instalar o binário.

## 1) Verificar se já existe (automático)

Quando você executa:

```
fig setup-ffi
```

O `fig` fará `which`/`exec.LookPath("ffi-helper")`. Se encontrar um binário global, o `fig.toml` será atualizado para apontar para esse caminho absoluto.

## 2) Compilar e instalar globalmente (recomendado)

```bash
cd tools/ffi-helper
go build -o ffi-helper .
# mover para /usr/local/bin (requer sudo)
sudo mv ffi-helper /usr/local/bin/
# ou instalar para ~/.local/bin e garantir que esteja no PATH
mv ffi-helper ~/.local/bin/
```

Após isso `fig setup-ffi` detectará automaticamente `ffi-helper` e gravará o caminho no `fig.toml`.

## 3) Compilar para o projeto (opção local)

Se preferir manter o helper dentro do repositório do projeto (sem instalar globalmente):

```bash
cd tools/ffi-helper
go build -o ../../.fig/ffi/ffi-helper .
```

O `fig setup-ffi` usa `./.fig/ffi/ffi-helper` como caminho **local por projeto** quando não existe um binário no PATH.

## 4) Se algo deu errado — verificação manual

- Verifique se `ffi-helper` está no PATH:
  - `command -v ffi-helper` ou `which ffi-helper`
- Verifique permissões de execução: `ls -l $(command -v ffi-helper)`
- Se instalou em `~/.local/bin`, assegure `export PATH="$HOME/.local/bin:$PATH"` no seu shell.

## 5) Referência rápida

- Comando de configuração: `fig setup-ffi`
- Local padrão quando não existe binário global: `./.fig/ffi/ffi-helper`
- Documentação online: https://github.com/isCarlosCoder/fig/docs/ffi/setup.md

---

Se quiser, posso adicionar instruções para instalação automática no `fig setup-ffi` (ex.: compilar e copiar para `~/.local/bin` quando detectado que usuário deseja). Quer que eu implemente isso também?