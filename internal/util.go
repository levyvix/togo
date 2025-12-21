// Package internal contém funções utilitárias para manipulação de dados.
package internal

import (
	"encoding/json"
	"fmt"
	"levyvix/go-todo-list/models"
	"os"
)

const (
	// JsonFileName é o nome do arquivo onde as tarefas são persistidas
	JsonFileName = "tasks.json"
)

// TasksFileName é o nome do arquivo JSON usado para armazenamento.
// Pode ser sobrescrito nos testes para usar arquivos temporários.
var TasksFileName = JsonFileName

// ReadJsonFile lê o arquivo JSON de tarefas.
//
// Retorna:
//   - []byte: conteúdo do arquivo em bytes
//   - error: nil se sucesso, erro em caso de falha na leitura
//
// Se o arquivo não existir, retorna um slice vazio e nil de erro.
// Se o arquivo existir mas houver erro ao ler, retorna erro apropriado.
func ReadJsonFile() ([]byte, error) {
	// Tenta ler o arquivo de tarefas
	data, err := os.ReadFile(TasksFileName)
	if err != nil {
		if os.IsNotExist(err) {
			// Arquivo não existe - retorna slice vazio (normal na primeira execução)
			var d []byte
			return d, nil
		} else {
			// Erro ao ler arquivo existente
			return nil, fmt.Errorf("Error reading json file %w", err)
		}
	}

	return data, nil
}

// WriteToJson salva uma tarefa adicionando-a ao arquivo JSON.
//
// NOTA: Esta função assume que o mutex já foi adquirido pelo chamador!
//
// Args:
//   - task: Task struct a ser salva
//
// Processo:
// 1. Lê tarefas existentes do arquivo
// 2. Desserializa dados JSON em um slice de Tasks
// 3. Adiciona a nova tarefa ao slice
// 4. Serializa o slice de volta para JSON (com indentação)
// 5. Escreve o arquivo com as novas tarefas
//
// Retorna:
//   - error: nil se sucesso, erro em caso de qualquer falha
func WriteToJson(task models.Task) error {
	// Lê dados existentes
	currentData, err := ReadJsonFile()
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo JSON: %w", err)
	}

	var taskList []models.Task
	// Se existem dados, faz parse
	if len(currentData) > 0 {
		err = json.Unmarshal(currentData, &taskList)
		if err != nil {
			return fmt.Errorf("erro ao desserializar dados JSON: %w", err)
		}
	}

	// Adiciona nova tarefa ao slice
	taskList = append(taskList, task)

	// Converte slice para JSON com indentação (2 espaços)
	task_json, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao converter struct para JSON: %w", err)
	}

	// Escreve arquivo com permissões 0644 (rw-r--r--)
	err = os.WriteFile(TasksFileName, task_json, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar arquivo JSON: %w", err)
	}

	return nil
}

// saveTasksToFile salva uma lista de tarefas no arquivo JSON.
//
// NOTA: Esta função assume que o mutex já foi adquirido pelo chamador!
//
// Args:
//   - tasks: slice de Task structs a ser salvo
//
// Retorna:
//   - error: nil se sucesso, erro em caso de qualquer falha
func saveTasksToFile(tasks []models.Task) error {
	// Converte slice para JSON com indentação (2 espaços)
	task_json, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao converter struct para JSON: %w", err)
	}

	// Escreve arquivo com permissões 0644 (rw-r--r--)
	err = os.WriteFile(TasksFileName, task_json, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar arquivo JSON: %w", err)
	}

	return nil
}
