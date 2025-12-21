# Documentação de Testes

## Visão Geral

O projeto inclui testes unitários para validar a funcionalidade principal da aplicação. Todos os testes usam o framework padrão `testing` do Go.

## Executando os Testes

### Rodar todos os testes

```bash
go test ./...
```

### Rodar testes com output verboso

```bash
go test ./... -v
```

### Rodar testes de um pacote específico

```bash
go test ./internal -v
go test ./cmd -v
```

### Rodar um teste específico

```bash
go test -run TestGetNextID ./internal -v
```

### Ver cobertura de testes

```bash
go test ./... -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Estrutura dos Testes

### `internal/functions_test.go`

Testes para as funções principais de negócio.

| Teste | O que testa |
|-------|-----------|
| `TestGetNextID` | Cálculo do próximo ID sequencial |
| `TestFormatDate` | Formatação de datas |
| `TestCreateTaskStructure` | Criação correta de struct Task |
| `TestTaskMarkedAsDone` | Marcação de tarefa como concluída |
| `TestWriteAndReadTask` | Salvar e ler uma tarefa |
| `TestMultipleTasks` | Operações com múltiplas tarefas |
| `TestRemoveTask` | Remoção de tarefa de um slice |

### `internal/util_test.go`

Testes para funções utilitárias de I/O e manipulação de arquivo JSON.

| Teste | O que testa |
|-------|-----------|
| `TestReadJsonFileNotExists` | Ler quando arquivo não existe |
| `TestReadJsonFileExists` | Ler quando arquivo existe |
| `TestWriteToJsonNewFile` | Escrever em arquivo novo |
| `TestWriteToJsonAppend` | Adicionar tarefa a arquivo existente |
| `TestSaveTasksToFileEmpty` | Salvar lista vazia |
| `TestSaveTasksToFileMultiple` | Salvar múltiplas tarefas |
| `TestJsonFileFormatting` | Validar formatação JSON |
| `TestJsonSpecialCharacters` | Validar caracteres especiais |

### `cmd/root_test.go`

Testes para os comandos Cobra.

| Teste | O que testa |
|-------|-----------|
| `TestRootCmdHelp` | Comando `--help` funciona |
| `TestRootCmdSubcommands` | Todos os subcomandos estão registrados |
| `TestCleanupTasksJsonFile` | Limpeza de arquivos de teste |

## Cobertura de Testes

Atualmente implementados:
- ✅ Testes unitários para funções de negócio
- ✅ Testes para I/O de arquivo
- ✅ Testes para parsing JSON
- ✅ Testes para validação de estruturas
- ✅ Testes para comandos Cobra

Não implementados (futuros):
- ⏳ Testes de integração completos (criar → listar → done → delete)
- ⏳ Testes de erro com arquivo corrompido
- ⏳ Testes de concorrência (race conditions)
- ⏳ Testes de performance
- ⏳ Testes end-to-end de CLI

## Boas Práticas nos Testes

### Limpeza de Recursos

Todos os testes que criam `tasks.json` usam `defer` para limpeza:

```go
func TestExample(t *testing.T) {
    defer os.Remove(JsonFileName)
    // ... seu teste aqui
}
```

### Table-Driven Tests

Para testes com múltiplos cenários, usamos a abordagem table-driven:

```go
tests := []struct {
    name        string
    input       string
    expected    int
    shouldError bool
}{
    {
        name:     "caso 1",
        input:    "valor1",
        expected: 100,
    },
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // teste aqui
    })
}
```

### Nomenclatura

- Nome do teste: `Test<FunctionName>` ou `Test<FunctionName><Scenario>`
- Exemplo: `TestGetNextID`, `TestWriteToJsonAppend`

## Rodando Testes Automaticamente

### Com Git Hooks

```bash
# Criar arquivo .git/hooks/pre-commit
#!/bin/bash
go test ./... -v
exit $?

chmod +x .git/hooks/pre-commit
```

### Com Makefile

Crie um arquivo `Makefile`:

```makefile
test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: test test-coverage
```

Depois execute:

```bash
make test
make test-coverage
```

## Debugging de Testes

### Rodar com output de debug

```bash
go test ./... -v -run TestGetNextID
```

### Adicionar prints no teste

```go
t.Logf("valor: %v", minhaVariavel)  // Aparece apenas em modo -v
```

### Usar debugger

```bash
go test -debug ./... # Usar com delve ou similar
```

## Manutenção de Testes

### Quando adicionar novos testes

1. Sempre que adicionar uma nova função pública
2. Quando corrigir um bug (adicionar teste que falha, depois corrigir)
3. Para casos extremos ou condições de erro

### Quando atualizar testes

1. Se mudar o comportamento de uma função (atualizar teste)
2. Se adicionar novo campo em struct (atualizar validações)
3. Se mudar formato de dados (atualizar parser JSON)

## Próximos Passos

- [ ] Adicionar testes de integração E2E
- [ ] Adicionar testes de concorrência
- [ ] Aumentar cobertura para 80%+
- [ ] Adicionar testes de erro/edge cases
- [ ] Criar CI/CD pipeline com GitHub Actions

## Troubleshooting

### Erro: "no test files in package main"

Testes não encontrados. Certifique-se que:
- Arquivo termina em `_test.go`
- Funções começam com `Test`
- Arquivo está no pacote correto

### Testes deixam `tasks.json` para trás

Adicione `defer os.Remove(JsonFileName)` no teste:

```go
func TestExample(t *testing.T) {
    defer os.Remove(JsonFileName)
    // seu teste
}
```

### Race condition detectada durante testes

Execute com detector de race:

```bash
go test -race ./...
```

Isso detectará problemas de concorrência no código.
