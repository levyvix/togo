# Documentação de Desenvolvimento

## Setup do Projeto

### Pré-requisitos
- Go 1.25.4 ou superior
- Git
- Justfiles (opcional, mas recomendado)

### Instalação das dependências

```bash
go mod download
```

### Build do projeto

```bash
go build -o togo
```

### Executar testes

```bash
go test ./...
```

## Estrutura do Projeto

```
togo/
├── cmd/                         # Comandos da aplicação
│   ├── root.go                 # Comando raiz (ponto de entrada)
│   ├── create.go               # Comando para criar tarefas
│   ├── list.go                 # Comando para listar tarefas
│   ├── done.go                 # Comando para marcar como concluída
│   ├── delete.go               # Comando para deletar tarefas
│   ├── edit.go                 # Comando para editar descrição
│   └── root_test.go            # Testes do comando raiz
├── internal/                    # Código interno (não exportado)
│   ├── functions.go            # Lógica dos comandos (CreateFuncDB, ListFuncDB, etc)
│   └── database/
│       └── db.go               # Inicialização e configuração do banco SQLite
├── schema/                      # Definições de esquema do banco
│   └── task.go                 # Definição da struct Task (modelo GORM)
├── main.go                      # Ponto de entrada da aplicação
├── go.mod                       # Definição do módulo Go
├── go.sum                       # Checksums das dependências
├── README.md                    # Documentação de uso
└── DEVELOPMENT.md              # Este arquivo
```

## Dependências

- **Cobra** (`github.com/spf13/cobra`): Framework CLI para Go
- **GORM** (`gorm.io/gorm`): ORM (Object-Relational Mapping) para Go
- **SQLite Driver** (`gorm.io/driver/sqlite`): Driver SQLite para GORM
- **go-sqlite3** (`github.com/mattn/go-sqlite3`): Binding SQLite3 (dependência indireta)

## Fluxo da Aplicação

```
main.go
  ├── database.InitDB()          # Inicializa SQLite e cria/migra tabelas
  └── cmd.Execute()
      └── rootCmd (Cobra)
          ├── createCmd
          │   └── internal.CreateFuncDB(args)
          │       ├── Valida argumentos
          │       ├── Cria Task struct
          │       ├── database.DB.Create(&task)  # INSERT na tabela tasks
          │       └── Printa mensagem de sucesso
          ├── listCmd
          │   └── internal.ListFuncDB()
          │       ├── database.DB.Order("id asc").Find(&tasks)  # SELECT
          │       └── Printa todas as tarefas formatadas
          ├── doneCmd
          │   └── internal.DoneFuncDB(args)
          │       ├── database.DB.First(&task, id)  # SELECT WHERE id
          │       ├── Atualiza Done=true e DoneAt
          │       ├── database.DB.Save(&task)  # UPDATE
          │       └── Printa confirmação
          ├── deleteCmd
          │   └── internal.DeleteFuncDB(args)
          │       └── database.DB.Delete(&task, id)  # DELETE FROM
          └── editCmd
              └── internal.EditFuncDB(args)
                  ├── database.DB.First(&task, id)  # SELECT WHERE id
                  ├── Atualiza Description
                  ├── database.DB.Save(&task)  # UPDATE
                  └── Printa confirmação
```

## Como adicionar um novo comando

### 1. Criar arquivo em `cmd/newcommand.go`. ou rodar o comando com o `cobra-cli`: `cobra-cli add newCommand`

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
}
```

### Modelo de Dados (GORM)

```go
type Task struct {
	gorm.Model                    // Fornece: ID, CreatedAt, UpdatedAt, DeletedAt
	Description string            // Descrição da tarefa
	Done        bool              // Status de conclusão
	DoneAt      *time.Time        // Timestamp de quando foi concluída
}
```

**Campos automáticos do `gorm.Model`:**
- `ID`: uint - Chave primária auto-incrementada
- `CreatedAt`: time.Time - Data de criação (automática)
- `UpdatedAt`: time.Time - Última atualização (automática)
- `DeletedAt`: *time.Time - Soft delete (NULL enquanto ativo)

### Operações no Banco de Dados

**Create (INSERT):**
```go
task := schema.Task{Description: "Estudar", Done: false}
database.DB.Create(&task)
```

**Read (SELECT):**
```go
var task schema.Task
database.DB.First(&task, id)  // Find by primary key

var tasks []schema.Task
database.DB.Order("id asc").Find(&tasks)  // Find all
```

**Update:**
```go
task.Done = true
task.DoneAt = &now
database.DB.Save(&task)
```

**Delete (Soft delete por padrão):**
```go
database.DB.Delete(&task, id)
```

### Formatação de Datas

Usar a função `formatDate()` em `internal/functions.go`:

```go
formatDate(time.Time) string // Retorna: "21 Dez 2025 14:30"
```

## Banco de Dados

- **Tipo**: SQLite
- **Localização**: `tasks.db` (na raiz do projeto)
- **ORM**: GORM
- **Criação automática**: Tabelas são criadas automaticamente via `AutoMigrate()` em `database.InitDB()`
- **Schema**: Definido pela struct `Task` em `schema/task.go`

### Inicialização do Banco

Em `main.go`:

```go
package main

import (
	"levyvix/togo/cmd"
	"levyvix/togo/internal/database"
)

func main() {
	database.InitDB()  // Inicializa SQLite e cria/migra tabelas
	cmd.Execute()      // Executa o comando CLI
}
```

O `database.InitDB()` faz:
1. Abre/cria o arquivo `tasks.db`
2. Executa `AutoMigrate()` para criar/atualizar a tabela `tasks`
3. Configura o logger do GORM para modo silencioso

## Correções e Melhorias Implementadas

### ✅ Corrigidas
- [x] Corrigido formatamento de erros (%w em fmt.Println)
- [x] Trocado `panic()` por `log.Fatalf()` para melhor tratamento
- [x] Adicionado ID sequencial único para tarefas
- [x] Adicionado mutex para sincronização (race conditions)
- [x] Melhorada validação de entrada
- [x] Mensagens em português consistentes
- [x] Output formatado com emojis
- [x] Migração de JSON para SQLite com GORM

### ✅ Funcionalidades Implementadas
- [x] Comando `done` para marcar tarefas como concluídas
- [x] Comando `delete` para remover tarefas
- [x] Comando `edit` para editar descrição de tarefas
- [x] ID sequencial automático via banco de dados
- [x] Timestamps automáticos (CreatedAt, UpdatedAt, DeletedAt)
- [x] Persistência em SQLite com GORM ORM

## Próximos Passos e Melhorias

- [ ] Implementar testes unitários e de integração abrangentes
- [x] Adicionar comando `edit` para editar descrição (implementado)
- [x] Adicionar persistência em SQLite com GORM (implementado)
- [ ] Adicionar filtros (listar apenas concluídas/pendentes)
- [ ] Implementar cores na saída com biblioteca como `github.com/fatih/color`
- [ ] Adicionar paginação para listas grandes
- [ ] Implementar busca/grep de tarefas
- [ ] Adicionar comando `clear` para limpar todas as tarefas
- [ ] Exportar tarefas em CSV/JSON para backup
- [ ] Integração com cron para tarefas recorrentes
- [ ] Adicionar suporte a categorias/tags de tarefas
- [ ] Implementar prioridades para tarefas
- [ ] Adicionar recurso de agendamento (data de vencimento)

## Recursos Úteis

- [Documentação Cobra](https://cobra.dev/)
- [Documentação Go](https://golang.org/doc/)
- [Tutorial: Building a CLI in Go](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)
- [GORM Documentation](https://gorm.io/)
- [GORM SQLite Driver](https://gorm.io/docs/databases/sqlite.html)
- [SQLite Official Documentation](https://www.sqlite.org/docs.html)
