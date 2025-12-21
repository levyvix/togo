# Guia Justfile

Este projeto usa **Justfile** para automatizar comandos frequentes. É similar a Makefile, mas com sintaxe mais simples.

## Instalação

### No Linux/macOS:
```bash
cargo install just
# ou
brew install just    # macOS com Homebrew
```

### No Windows:
```bash
cargo install just
# ou use chocolatey/scoop
```

### Verificar instalação:
```bash
just --version
```

## Como Usar

### Listar todos os comandos:
```bash
just
# ou
just --list
```

### Rodar um comando:
```bash
just <comando>
```

## Comandos Disponíveis

### 🏗️ BUILD

```bash
just build            # Compilar aplicação
just build-release    # Build de release (com versão)
just install          # Instalar globalmente ($GOBIN)
just clean            # Limpar binários e temp
```

### 🧪 TESTES

```bash
just test             # Rodar testes com verbose
just test-coverage    # Testes com cobertura (gera coverage.html)
just test-race        # Testes com race detector
just test-one TestName    # Rodar um teste específico
just test-demo        # Demo interativa: create → list → done → delete
```

### 🎨 FORMATAÇÃO

```bash
just fmt              # Formatar código
just fmt-check        # Verificar formatação (sem modificar)
just vet              # Rodar vet (detecta problemas comuns)
just lint             # Rodar golangci-lint (se instalado)
```

### 🚀 DESENVOLVIMENTO

```bash
just run              # Compilar e rodar
just run create "Minha tarefa"    # Rodar com argumentos
just watch            # Recompilar ao detectar mudanças (requer entr)
```

### 📚 DOCUMENTAÇÃO

```bash
just help             # Ver help da CLI
just help create      # Ver help de um comando específico
just docs             # Gerar documentação (godoc)
just info             # Mostrar info do projeto
```

### 📦 RELEASES

```bash
just release          # Build multi-platform (Linux, macOS, Windows)
```

### 🔄 CI/CD

```bash
just check            # Rodar: fmt-check + vet + test
just pipeline         # Rodar: clean + fmt + vet + test + build
```

### 🛠️ UTILITÁRIOS

```bash
just deps-update      # Atualizar dependências (go get -u)
just deps-clean       # Limpar dependências não usadas (go mod tidy)
just show-justfile    # Mostrar conteúdo do justfile
```

## Exemplos Práticos

### Fluxo de Desenvolvimento Normal:
```bash
just build            # 1. Compilar
just run list         # 2. Testar localmente
just test             # 3. Rodar testes
just fmt              # 4. Formatar código
```

### Antes de fazer commit:
```bash
just pipeline         # Roda: clean → fmt → vet → test → build
```

### Depois que código está pronto:
```bash
just release          # Build para distribuir (Linux, macOS, Windows)
ls -lh release/       # Ver binários prontos
```

### Desenvolvimento iterativo:
```bash
just watch            # Recompilar ao salvar arquivo
```

### Entender o projeto:
```bash
just info             # Ver linhas de código, estrutura, etc
just help             # Ver ajuda da CLI
just help done        # Ver ajuda de um comando específico
```

## Anatomia do Justfile

Cada comando tem:

```justfile
# Comentário
@nome-do-comando PARAMETRO="default":
    echo "O que está acontecendo..."
    comando-real
    echo "Pronto!"
```

- `@` = Não mostra os comandos sendo executados
- `${{ VAR }}` = Variável
- `{{PARAMETRO}}` = Parâmetro passado
- `#!/bin/bash` = Script multi-linha

## Personalizando

Você pode editar `justfile` para:

1. **Adicionar novos comandos:**
```justfile
@meu-comando:
    echo "Fazendo algo..."
    go build
```

2. **Modificar binário ou versão:**
```justfile
BINARY_NAME := "meu-app"
VERSION := "2.0.0"
```

3. **Adicionar mais plataformas em release:**
```justfile
GOOS=linux GOARCH=arm64 go build ...
```

## Comparação: Justfile vs Makefile

| Aspecto | Justfile | Makefile |
|---------|----------|----------|
| Sintaxe | Simples | Complexa |
| Variáveis | `{{VAR}}` | `$(VAR)` |
| Shells | bash | sh padrão |
| Tabs | Não obrigatório | **OBRIGATÓRIO** |
| Legibilidade | Alta | Média |
| Aprendizado | Fácil | Médio |

## Comparação: Justfile vs Python pytest

```python
# Python: pytest
pytest tests/
pytest tests/test_models.py -v
pytest tests/ --cov=src

# Go/Just: testes integrados
just test
just test-coverage
just test-race
```

**Go é diferente:** testes ficam no pacote, não em pasta separada.

## Troubleshooting

### Erro: "just: command not found"
```bash
cargo install just
```

### Erro: "entr: command not found"
```bash
# Não é obrigatório, só para `just watch`
# Instale se quiser: brew install entr
```

### Erro: "golangci-lint not installed"
```bash
# Opcional, para lint mais robusto
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Dicas

1. **Use `just pipeline` antes de fazer commit:**
   ```bash
   just pipeline  # Garante código válido
   ```

2. **Adicione ao seu shell:**
   ```bash
   # ~/.bashrc ou ~/.zshrc
   alias j=just
   ```
   Então use: `j build` em vez de `just build`

3. **Veja o que o comando faz antes de executar:**
   ```bash
   just --dry-run build  # Mostra comandos sem executar
   ```

4. **Edite o justfile para sua workflow:**
   - Adicione comandos frequentes
   - Remove os que não usa
   - Mantenha organizado

## Próximas Melhorias

- [ ] Adicionar script de setup automático
- [ ] Integrar com GitHub Actions
- [ ] Adicionar testes de integração
- [ ] Adicionar Docker build

## Referências

- [Documentação oficial Just](https://just.systems/)
- [Cookbook Just](https://just.systems/man/en/chapter_31.html)
