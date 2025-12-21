# Documentação de Desenvolvimento

## Setup do Projeto

### Pré-requisitos
- Go 1.25.4 ou superior
- Git

### Instalação das dependências

```bash
go mod download
```

### Build do projeto

```bash
go build -o go-todo-list
```

### Executar testes

```bash
go test ./...
```

## Estrutura do Projeto

```
go-todo-list/
├── cmd/                    # Comandos da aplicação
│   ├── root.go            # Comando raiz (ponto de entrada)
│   ├── create.go          # Comando para criar tarefas
│   └── list.go            # Comando para listar tarefas
├── internal/              # Código interno (não exportado)
│   ├── functions.go       # Lógica dos comandos (CreateFunc, ListFunc)
│   └── util.go            # Utilitários (leitura/escrita JSON)
├── models/                # Estruturas de dados
│   └── models.go          # Definição da struct Task
├── main.go                # Ponto de entrada da aplicação
├── go.mod                 # Definição do módulo Go
├── go.sum                 # Checksums das dependências
├── README.md              # Documentação de uso
└── DEVELOPMENT.md         # Este arquivo
```

## Dependências

- **Cobra** (`github.com/spf13/cobra`): Framework CLI para Go
- **PFlag** (`github.com/spf13/pflag`): Flag parsing library

## Fluxo da Aplicação

```
main.go
  └── cmd.Execute()
      └── rootCmd (Cobra)
          ├── createCmd
          │   └── internal.CreateFunc(args)
          │       ├── Cria Task struct
          │       ├── Chama util.WriteToJson()
          │       └── Printa mensagem de sucesso
          └── listCmd
              └── internal.ListFunc()
                  ├── Chama util.ReadJsonFile()
                  ├── Faz parse do JSON
                  └── Printa todas as tarefas
```

## Como adicionar um novo comando

### 1. Criar arquivo em `cmd/newcommand.go`

```go
package cmd

import (
	"levyvix/go-todo-list/internal"
	"github.com/spf13/cobra"
)

var newcommandCmd = &cobra.Command{
	Use:   "newcommand",
	Short: "Descrição curta",
	Long:  `Descrição longa`,
	Run: func(cmd *cobra.Command, args []string) {
		// Lógica do comando aqui
		internal.NewcommandFunc(args)
	},
}

func init() {
	rootCmd.AddCommand(newcommandCmd)
}
```

### 2. Implementar a função em `internal/functions.go`

```go
func NewcommandFunc(args []string) {
	// Implementação
}
```

### 3. Registrar a importação (automaticamente feita pelo init)

## Padrões de Código

### Tratamento de Erros

```go
if err != nil {
	log.Fatalf("Erro ao fazer algo: %v", err)
	os.Exit(1)
}
```

### Estrutura JSON

As tarefas são persistidas com:
- `description`: string - descrição da tarefa
- `done`: bool - se está concluída
- `createdAt`: time.Time - quando foi criada
- `doneAt`: *time.Time - quando foi concluída (nil se não feita)

### Formatação de Datas

Usar a função `formatDate()` em `internal/functions.go`:

```go
formatDate(time.Time) string // Retorna: "21 Dec 2025 14:30"
```

## Arquivo de Dados

- **Localização**: `tasks.json` (na raiz do projeto)
- **Formato**: JSON array de Task structs
- **Criação automática**: Se não existir, é criado vazio no primeiro write

## Correções e Melhorias Implementadas

### ✅ Corrigidas
- [x] Corrigido formatamento de erros (%w em fmt.Println)
- [x] Trocado `panic()` por `log.Fatalf()` para melhor tratamento
- [x] Adicionado ID sequencial único para tarefas
- [x] Adicionado mutex para sincronização (race conditions)
- [x] Melhorada validação de entrada
- [x] Mensagens em português consistentes
- [x] Output formatado com emojis

### ✅ Funcionalidades Implementadas
- [x] Comando `done` para marcar tarefas como concluídas
- [x] Comando `delete` para remover tarefas
- [x] ID sequencial automático

## Próximos Passos e Melhorias

- [ ] Implementar testes unitários e de integração
- [ ] Adicionar comando `edit` para editar descrição
- [ ] Adicionar filtros (listar apenas concluídas/pendentes)
- [ ] Adicionar persistência em diferentes formatos (YAML, SQLite)
- [ ] Implementar cores na saída com biblioteca como `github.com/fatih/color`
- [ ] Adicionar paginação para listas grandes
- [ ] Implementar busca/grep de tarefas
- [ ] Adicionar comando `clear` para limpar todas as tarefas
- [ ] Exportar tarefas em CSV/JSON para backup
- [ ] Integração com cron para tarefas recorrentes

## Recursos Úteis

- [Documentação Cobra](https://cobra.dev/)
- [Documentação Go](https://golang.org/doc/)
- [Tutorial: Building a CLI in Go](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)
