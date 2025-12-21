# Justfile - Automatize comandos frequentes
# Instale com: cargo install just
# Use com: just <comando>

set shell := ["bash", "-c"]
set positional-arguments

# Variáveis
BINARY_NAME := "go-todo-list"
VERSION := "1.1.0"

# Mostrar ajuda (padrão)
@default:
    just --list

# ═══════════════════════════════════════════════════════════
# 🏗️  BUILD
# ═══════════════════════════════════════════════════════════

# Build da aplicação
@build:
    echo "🔨 Compilando {{BINARY_NAME}}..."
    go build -o {{BINARY_NAME}}
    echo "✅ Build concluído!"

# Build com informações de versão
@build-release:
    echo "🔨 Compilando release {{VERSION}}..."
    go build \
        -ldflags="-s -w -X main.Version={{VERSION}}" \
        -o {{BINARY_NAME}}
    echo "✅ Release build concluído!"

# Install da aplicação (em $GOBIN ou $GOPATH/bin)
@install:
    echo "📦 Instalando {{BINARY_NAME}}..."
    go install
    echo "✅ Instalação concluída!"

# Clean - remove binários
@clean:
    echo "🧹 Limpando..."
    rm -f {{BINARY_NAME}}
    rm -f tasks.json
    rm -f coverage.out coverage.html
    go clean
    echo "✅ Limpeza concluída!"

# ═══════════════════════════════════════════════════════════
# 🧪 TESTES
# ═══════════════════════════════════════════════════════════

# Rodar testes com verbose
@test:
    echo "🧪 Rodando testes..."
    go test ./... -v

# Rodar testes com coverage
@test-coverage:
    echo "🧪 Rodando testes com cobertura..."
    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out -o coverage.html
    echo "✅ Relatório gerado: coverage.html"

# Rodar testes com race detector
@test-race:
    echo "🧪 Rodando testes com race detector..."
    go test -race ./...

# Rodar apenas um teste
@test-one TEST="":
    @if [ -z "{{TEST}}" ]; then \
        echo "❌ Use: just test-one TestName"; \
        exit 1; \
    fi
    echo "🧪 Rodando teste: {{TEST}}"
    go test -run {{TEST}} ./... -v

# ═══════════════════════════════════════════════════════════
# 🎨 FORMATAÇÃO E QUALIDADE
# ═══════════════════════════════════════════════════════════

# Formatar código
@fmt:
    echo "🎨 Formatando código..."
    go fmt ./...
    echo "✅ Formatação concluída!"

# Verificar formatação (sem modificar)
@fmt-check:
    echo "🔍 Verificando formatação..."
    @if ! go fmt ./... > /dev/null 2>&1; then \
        echo "❌ Código não está formatado!"; \
        echo "Execute: just fmt"; \
        exit 1; \
    fi
    echo "✅ Código está formatado!"

# Lint (se tiver golangci-lint instalado)
@lint:
    @if ! command -v golangci-lint &> /dev/null; then \
        echo "⚠️  golangci-lint não instalado"; \
        echo "Instale com: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
        exit 1; \
    fi
    echo "🔍 Rodando linter..."
    golangci-lint run ./...

# vet - Verificar problemas comuns
@vet:
    echo "🔍 Rodando vet..."
    go vet ./...
    echo "✅ Verificação concluída!"

# ═══════════════════════════════════════════════════════════
# 🚀 DESENVOLVIMENTO
# ═══════════════════════════════════════════════════════════

# Rodar aplicação
@run *ARGS:
    echo "▶️  Rodando {{BINARY_NAME}}..."
    just build
    ./{{BINARY_NAME}} {{ARGS}}

# Rodar teste de funcionalidade completo
@test-demo:
    #!/bin/bash
    echo "🎬 DEMO - Teste Completo"
    just clean
    just build

    echo ""
    echo "📝 Criando tarefas..."
    ./{{BINARY_NAME}} create "Aprender Go"
    ./{{BINARY_NAME}} create "Escrever testes"
    ./{{BINARY_NAME}} create "Deploy"

    echo ""
    echo "📋 Listando tarefas..."
    ./{{BINARY_NAME}} list

    echo ""
    echo "✓ Marcando tarefa 1 como concluída..."
    ./{{BINARY_NAME}} done 1

    echo ""
    echo "📋 Listando após marcar como done..."
    ./{{BINARY_NAME}} list

    echo ""
    echo "🗑️  Deletando tarefa 2..."
    ./{{BINARY_NAME}} delete 2

    echo ""
    echo "📋 Lista final..."
    ./{{BINARY_NAME}} list

# Watch - Recompilar ao detectar mudanças (requer entr)
@watch:
    @if ! command -v entr &> /dev/null; then \
        echo "⚠️  entr não instalado"; \
        echo "Instale com: cargo install entr (ou package manager)"; \
        exit 1; \
    fi
    echo "👀 Monitorando mudanças..."
    find . -name '*.go' | entr -r just build

# ═══════════════════════════════════════════════════════════
# 📚 DOCUMENTAÇÃO
# ═══════════════════════════════════════════════════════════

# Gerar documentação (godoc)
@docs:
    echo "📚 Gerando documentação..."
    @if command -v godoc &> /dev/null; then \
        echo "Abra: http://localhost:6060/pkg/levyvix/go-todo-list/"; \
        godoc -http=:6060; \
    else \
        echo "godoc não disponível, mostrando README:"; \
        cat README.md; \
    fi

# Ver help de um comando
@help COMMAND="":
    @if [ -z "{{COMMAND}}" ]; then \
        echo "Use: just <comando> --help"; \
        just build > /dev/null 2>&1; \
        ./{{BINARY_NAME}} --help; \
    else \
        just build > /dev/null 2>&1; \
        ./{{BINARY_NAME}} {{COMMAND}} --help; \
    fi

# ═══════════════════════════════════════════════════════════
# 📦 RELEASES
# ═══════════════════════════════════════════════════════════

# Preparar release (build multi-platform)
@release:
    #!/bin/bash
    echo "📦 Preparando release {{VERSION}}..."

    mkdir -p release

    # Linux
    echo "Building for Linux..."
    GOOS=linux GOARCH=amd64 go build -o release/{{BINARY_NAME}}-linux-amd64

    # macOS
    echo "Building for macOS..."
    GOOS=darwin GOARCH=amd64 go build -o release/{{BINARY_NAME}}-macos-amd64

    # Windows
    echo "Building for Windows..."
    GOOS=windows GOARCH=amd64 go build -o release/{{BINARY_NAME}}-windows-amd64.exe

    echo "✅ Releases criados em ./release/"
    ls -lh release/

# ═══════════════════════════════════════════════════════════
# 🔄 CI/CD
# ═══════════════════════════════════════════════════════════

# Rodar checks (formato, vet, testes)
@check: fmt-check vet test
    echo "✅ Todos os checks passaram!"

# Rodar pipeline completo
@pipeline: clean fmt vet test build
    echo "✅ Pipeline completo executado com sucesso!"

# ═══════════════════════════════════════════════════════════
# 🛠️  UTILITÁRIOS
# ═══════════════════════════════════════════════════════════

# Atualizar dependências
@deps-update:
    echo "📦 Atualizando dependências..."
    go get -u ./...
    go mod tidy
    echo "✅ Dependências atualizadas!"

# Limpar módulos não usados
@deps-clean:
    echo "🧹 Limpando dependências..."
    go mod tidy
    echo "✅ Dependências limpas!"

# Verificar dependências com vulnerabilidades
@deps-check:
    echo "🔍 Verificando vulnerabilidades..."
    go list -json -m all | nancy sleuth
    echo "✅ Verificação concluída!"

# Mostrar informações do projeto
@info:
    @echo "📊 Informações do Projeto"
    @echo "========================"
    @echo "Nome: {{BINARY_NAME}}"
    @echo "Versão: {{VERSION}}"
    @echo "Go Version: $(go version | awk '{print $3}')"
    @echo ""
    @echo "📁 Estrutura:"
    @find . -maxdepth 2 -type d -not -path '*/.*' | sort
    @echo ""
    @echo "📊 Linhas de código:"
    @find . -name '*.go' -not -path './.*' | xargs wc -l | tail -1

# ═══════════════════════════════════════════════════════════
# 🎓 HELP
# ═══════════════════════════════════════════════════════════

# Mostrar todos os comandos disponíveis
@list-commands:
    just --list

# Mostrar este arquivo
@show-justfile:
    cat justfile
