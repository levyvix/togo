# Just vs Makefile - Comparação Prática

Você pode usar **Justfile** em vez de Makefile. Aqui está por que Just é melhor:

## Sintaxe Lado a Lado

### Makefile (antigo, complicado):
```makefile
.PHONY: build test clean

BINARY=go-todo-list
VERSION=1.1.0

build:
	@echo "🔨 Compilando..."
	go build -o $(BINARY)
	@echo "✅ Pronto!"

test:
	@echo "🧪 Testando..."
	go test ./... -v

clean:
	rm -f $(BINARY)
	rm -rf coverage.*
```

**Problemas:**
- ❌ Tabs são **obrigatórios** (causa erros misteriosos!)
- ❌ Sintaxe `$(VAR)` é estranha
- ❌ Menos legível
- ❌ Variáveis complexas

### Justfile (novo, simples):
```justfile
BINARY := "go-todo-list"
VERSION := "1.1.0"

@build:
    echo "🔨 Compilando..."
    go build -o {{BINARY}}
    echo "✅ Pronto!"

@test:
    echo "🧪 Testando..."
    go test ./... -v

@clean:
    rm -f {{BINARY}}
    rm -rf coverage.*
```

**Vantagens:**
- ✅ Espaços normais (sem tabs obsessivos)
- ✅ Sintaxe `{{VAR}}` intuitiva
- ✅ Mais legível
- ✅ Fácil de aprender

---

## Comparação Detalhada

| Aspecto | Makefile | Justfile |
|---------|----------|----------|
| **Sintaxe** | `$(VAR)` | `{{VAR}}` |
| **Indentação** | **Tabs obrigatórios** | Espaços normais |
| **Comentários** | `# comentário` | `# comentário` |
| **Parâmetros** | `make build VAR=valor` | `just build valor` |
| **Condicionais** | `ifdef VAR` | `if [ -z "{{VAR}}" ]` |
| **Loops** | Possível mas complexo | Shell bash integrado |
| **Help automático** | Não | Sim (`just --list`) |
| **Dry-run** | Não padrão | `just --dry-run` |
| **Multi-linha** | Complexo com `\` | Simples com `#!/bin/bash` |

---

## Exemplos Práticos

### Exemplo 1: Build com parâmetros

**Makefile:**
```makefile
BUILD_FLAGS ?= -v
build:
	@echo "Compilando com flags: $(BUILD_FLAGS)"
	go build $(BUILD_FLAGS)
```

Usar: `make build BUILD_FLAGS="-race"`

**Justfile:**
```justfile
@build FLAGS="-v":
    echo "Compilando com flags: {{FLAGS}}"
    go build {{FLAGS}}
```

Usar: `just build "-race"`

✅ Justfile é mais intuitivo!

---

### Exemplo 2: Script multi-linha

**Makefile:**
```makefile
demo:
	@echo "Criando tarefa..."
	./app create "Task"
	@echo "Listando..."
	./app list
	@echo "Pronto!"
```

**Justfile:**
```justfile
@demo:
    #!/bin/bash
    echo "Criando tarefa..."
    ./app create "Task"
    echo "Listando..."
    ./app list
    echo "Pronto!"
```

✅ Justfile usa bash nativo, mais claro!

---

### Exemplo 3: Verificações condicionais

**Makefile:**
```makefile
check-tool:
	@command -v golangci-lint >/dev/null 2>&1 || \
		(echo "golangci-lint não instalado"; exit 1)
	golangci-lint run
```

**Justfile:**
```justfile
@check-tool:
    @if ! command -v golangci-lint &> /dev/null; then \
        echo "golangci-lint não instalado"; \
        exit 1; \
    fi
    golangci-lint run
```

✅ Ambos funcionam, mas Justfile é mais bash-like!

---

## Quando Usar Cada Um

### Use **Makefile** se:
- ❌ Projeto exige compatibilidade com Make (C/C++)
- ❌ Infraestrutura legacy só conhece Make
- ❌ Você gosta de tradição

### Use **Justfile** se:
- ✅ Projeto moderno (Go, Rust, Node.js)
- ✅ Quer sintaxe clara e simples
- ✅ Quer melhor experiência para desenvolvedores
- ✅ Quer `--list` automático (help)

---

## Migrando de Make para Just

### Passo 1: Converter sintaxe

```makefile
# Makefile
BINARY=$(APP)
TESTS=$(shell find . -name '*_test.go')

$(BINARY):
	go build -o $(BINARY)
```

```justfile
# Justfile
BINARY := "app"
TESTS := `find . -name '*_test.go'`

build:
    go build -o {{BINARY}}
```

Mudanças:
- `$(VAR)` → `{{VAR}}`
- `$()` shells → `` `comando` ``
- Remover targets desnecessários

### Passo 2: Testar
```bash
just --dry-run build    # Ver sem executar
just build              # Executar
```

### Passo 3: Deletar Makefile
```bash
rm Makefile
```

---

## Justfile no seu Projeto

Seu projeto tem um Justfile bem completo!

```bash
just              # Ver todos os comandos
just build        # Compilar
just test         # Testar
just pipeline     # Build + test + format + vet
just release      # Multi-platform build
```

Explore:
```bash
just --list       # Todos os comandos
just --dry-run test   # Ver o que faria
```

---

## Setup Recomendado

### 1. Instalar Just:
```bash
cargo install just
```

### 2. Colocar alias em `.bashrc`/.`zshrc`:
```bash
alias j=just
```

### 3. Usar no seu workflow:
```bash
j build      # em vez de: just build
j test       # em vez de: go test ./...
j pipeline   # em vez de: limpar, formatar, testar, compilar...
```

---

## FAQ

**P: Makefile ainda é relevante?**
R: Sim, para projetos C/C++. Para Go/Rust/Node, Just é melhor.

**P: Preciso desinstalar Just depois?**
R: Não, Just é portável. Instale uma vez, use para sempre.

**P: Posso ter Makefile E Justfile?**
R: Sim! Mas é redundante. Escolha um.

**P: Just funciona no Windows?**
R: Sim, totalmente compatível.

**P: Como distribuo o projeto?**
R: Inclua o Justfile no repositório. Usuários instalam Just uma vez.

---

## Exemplo Real: Seu Projeto

Veja seu `justfile`:

```bash
just --list
```

Alguns comandos úteis:

```bash
just build              # Compilar
just test               # Testes
just fmt                # Formatar
just check              # Todos os checks
just pipeline           # Ciclo completo
just test-coverage      # Com cobertura
just release            # Multi-platform
```

---

## Conclusão

| Critério | Make | Just |
|----------|------|------|
| Simplicidade | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| Legibilidade | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Aprendizado | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Compatibilidade | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| Para Go | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

**Recomendação para Go:** Use Just! 🚀

---

## Links

- [Documentação Just](https://just.systems/)
- [GitHub Just](https://github.com/casey/just)
- [Cookbook](https://just.systems/man/en/chapter_31.html)
