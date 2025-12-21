# Go Todo List CLI

Uma aplicação de linha de comando simples e eficiente para gerenciar sua lista de tarefas, construída em Go com o framework Cobra.

## Features

- ✅ Criar novas tarefas com ID sequencial único
- 📋 Listar todas as tarefas com detalhes
- 💾 Persistência de dados em JSON
- ✓ Marcar tarefas como concluídas
- 🗑️ Deletar tarefas
- 🕐 Registro automático de datas de criação e conclusão
- 🔒 Proteção contra race conditions com Mutex
- 🎯 Interface CLI intuitiva
- 📊 Formatação clara com emojis

## Quick Start

### Pré-requisitos
- Go 1.25.4 ou superior
- `just` (opcional, mas recomendado para desenvolvimento)

### Instalação

```bash
git clone https://github.com/levyvix/togo.git
cd togo
just build
```

Ou sem `just`:

```bash
go build -o togo
```

### Usar globalmente

```bash
just install
```

Ou sem `just`:

```bash
go install ./...
```

## Uso

### Comandos Disponíveis

#### 1. Criar uma tarefa

```bash
./togo create "Descrição da tarefa"
```

**Exemplo:**
```bash
./togo create "Estudar Go"
./togo create "Fazer compras"
./togo create "Revisar código"
```

**Saída esperada:**
```
✓ Tarefa criada! ID: 1 | 'Estudar Go'
✓ Tarefa criada! ID: 2 | 'Fazer compras'
```

#### 2. Listar todas as tarefas

```bash
./togo list
```

**Saída esperada:**
```
📋 Lista de Tarefas:
==================================================
[1] ⏳ Estudar Go
    Criada em: 21 Dec 2025 14:30
--------------------------------------------------
[2] ✓ Fazer compras
    Criada em: 21 Dec 2025 14:31
    Concluída em: 21 Dec 2025 15:45
--------------------------------------------------
[3] ⏳ Revisar código
    Criada em: 21 Dec 2025 14:32
--------------------------------------------------
```

**Legenda:**
- `✓` = Tarefa concluída
- `⏳` = Tarefa pendente

#### 3. Marcar uma tarefa como concluída

```bash
./togo done <id>
```

**Exemplo:**
```bash
./togo done 1
```

**Saída esperada:**
```
✓ Tarefa 1 marcada como concluída!
```

#### 4. Deletar uma tarefa

```bash
./togo delete <id>
```

**Exemplo:**
```bash
./togo delete 2
```

**Saída esperada:**
```
✓ Tarefa 2 deletada!
```

### Ajuda

Para ver a ajuda dos comandos:

```bash
./togo --help
./togo create --help
./togo list --help
```

## Estrutura de Dados

As tarefas são armazenadas em um arquivo JSON (`tasks.json`) com a seguinte estrutura:

```json
[
  {
    "id": 1,
    "description": "Estudar Go",
    "done": false,
    "createdAt": "2025-12-21T14:30:00Z",
    "doneAt": null
  },
  {
    "id": 2,
    "description": "Fazer compras",
    "done": true,
    "createdAt": "2025-12-21T14:31:00Z",
    "doneAt": "2025-12-21T15:45:00Z"
  }
]
```

**Campos:**
- `id`: Identificador único sequencial (começando em 1)
- `description`: Descrição da tarefa
- `done`: Status de conclusão (true/false)
- `createdAt`: Data e hora de criação
- `doneAt`: Data e hora de conclusão (null se pendente)

## Testes

### Executar testes

```bash
just test
```

Ou sem `just`:

```bash
go test ./...
```

### Testes com cobertura

```bash
just test-coverage
```

Isso gera um relatório HTML em `coverage.html`.

### Testes com race detector

```bash
just test-race
```

### Executar um teste específico

```bash
just test-one TestName
```

## Desenvolvimento

Para informações detalhadas sobre a arquitetura, estrutura do projeto, padrões de código e como adicionar novos comandos, veja [DEVELOPMENT.md](./DEVELOPMENT.md).

### Comandos de desenvolvimento com `just`

O projeto usa `just` para automatizar tarefas frequentes. Principais comandos:

**Build:**
```bash
just build           # Build básico
just build-release   # Build com informações de versão
just clean          # Limpar binários e arquivos temporários
```

**Testes:**
```bash
just test           # Rodar testes com verbose
just test-coverage  # Rodar testes com cobertura
just test-race      # Rodar testes com race detector
just test-demo      # Demo completa (cria, lista, marca como done, deleta)
```

**Qualidade de Código:**
```bash
just fmt            # Formatar código
just fmt-check      # Verificar formatação sem modificar
just vet            # Verificar problemas comuns
just lint           # Rodar linter (requer golangci-lint)
```

**Desenvolvimento:**
```bash
just run [ARGS]     # Build e rodar com argumentos
just watch          # Recompilar ao detectar mudanças (requer entr)
```

**Pipeline Completo:**
```bash
just pipeline       # clean → fmt → vet → test → build
just check          # fmt-check → vet → test
```

**Dependências:**
```bash
just deps-update    # Atualizar dependências
just deps-clean     # Limpar módulos não usados
just deps-check     # Verificar vulnerabilidades (requer nancy)
```

**Informações:**
```bash
just info           # Mostrar informações do projeto
just list-commands  # Listar todos os comandos disponíveis
just help [CMD]     # Ver ajuda de um comando
```

Para instalar `just`, veja: https://github.com/casey/just

## Estrutura do Projeto

```
togo/
├── cmd/                  # Comandos da aplicação
│   ├── root.go          # Comando raiz
│   ├── create.go        # Comando create
│   ├── delete.go        # Comando delete
│   ├── done.go          # Comando done
│   ├── list.go          # Comando list
│   └── root_test.go     # Testes do comando raiz
├── internal/            # Código interno
│   ├── functions.go     # Lógica dos comandos
│   ├── util.go          # Utilitários JSON
│   └── *_test.go        # Testes unitários
├── models/              # Estruturas de dados
│   └── models.go        # Definição de Task
├── main.go              # Ponto de entrada
├── go.mod              # Definição do módulo
├── go.sum              # Checksums das dependências
├── justfile            # Automação de tarefas
├── README.md           # Este arquivo
└── DEVELOPMENT.md      # Documentação de desenvolvimento
```

## Dependências

- **Cobra** (`github.com/spf13/cobra`): Framework CLI para Go
- **PFlag** (`github.com/spf13/pflag`): Flag parsing library

## Licença

MIT

## Autor

Levyvix
