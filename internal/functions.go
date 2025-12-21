// Package internal contém as funções de lógica de negócio da aplicação.
package internal

import (
	"encoding/json"
	"fmt"
	"levyvix/go-todo-list/models"
	"log"
	"os"
	"sync"
	"time"
)

// mu protege o acesso ao arquivo JSON contra race conditions
var mu sync.Mutex

// CreateFunc cria uma nova tarefa e a salva no arquivo JSON.
//
// Args:
//   - args: slice de strings contendo a descrição da tarefa no índice 0
//
// A função cria um novo Task struct com:
// - Description: valor fornecido em args[0]
// - ID: próximo ID sequencial
// - Done: false (nova tarefa é sempre pendente)
// - CreatedAt: timestamp atual
// - DoneAt: nil (não marcada como concluída)
//
// Em caso de erro ao salvar, exibe mensagem de erro e encerra com status 1.
func CreateFunc(args []string) {
	description := args[0]

	// Obter próximo ID
	nextID, err := getNextID()
	if err != nil {
		log.Fatalf("Erro ao obter próximo ID: %v\n", err)
	}

	t := models.Task{
		ID:          nextID,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}

	err = WriteToJson(t)
	if err != nil {
		log.Fatalf("Erro ao salvar tarefa: %v\n", err)
	}

	fmt.Printf("✓ Tarefa criada! ID: %d | '%s'\n", t.ID, t.Description)
}

// getNextID calcula o próximo ID sequencial disponível.
// Lê todas as tarefas e retorna max(ID) + 1.
// Se não houver tarefas, retorna 1.
func getNextID() (int, error) {
	currentData, err := ReadJsonFile()
	if err != nil {
		return 0, err
	}

	var tasks []models.Task
	if len(currentData) > 0 {
		err = json.Unmarshal(currentData, &tasks)
		if err != nil {
			return 0, err
		}
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	return maxID + 1, nil
}

// DoneFunc marca uma tarefa como concluída pelo ID.
//
// Args:
//   - args: slice contendo o ID da tarefa no índice 0
//
// Se a tarefa for encontrada, marca como Done=true e registra DoneAt com timestamp atual.
// Se não encontrada, exibe erro.
func DoneFunc(args []string) {
	if len(args) == 0 {
		log.Fatal("Erro: você deve fornecer o ID da tarefa a marcar como concluída")
	}

	var taskID int
	_, err := fmt.Sscanf(args[0], "%d", &taskID)
	if err != nil {
		log.Fatalf("Erro: ID deve ser um número. Você passou '%s'\n", args[0])
	}

	mu.Lock()
	defer mu.Unlock()

	currentData, err := ReadJsonFile()
	if err != nil {
		log.Fatalf("Erro ao ler tarefas: %v\n", err)
	}

	var tasks []models.Task
	if len(currentData) > 0 {
		err = json.Unmarshal(currentData, &tasks)
		if err != nil {
			log.Fatalf("Erro ao decodificar JSON: %v\n", err)
		}
	}

	found := false
	now := time.Now()
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Done = true
			tasks[i].DoneAt = &now
			found = true
			break
		}
	}

	if !found {
		log.Fatalf("Erro: tarefa com ID %d não encontrada\n", taskID)
	}

	err = saveTasksToFile(tasks)
	if err != nil {
		log.Fatalf("Erro ao salvar tarefas: %v\n", err)
	}

	fmt.Printf("✓ Tarefa %d marcada como concluída!\n", taskID)
}

// DeleteFunc remove uma tarefa pelo ID.
//
// Args:
//   - args: slice contendo o ID da tarefa no índice 0
//
// Se a tarefa for encontrada, a remove da lista.
// Se não encontrada, exibe erro.
func DeleteFunc(args []string) {
	if len(args) == 0 {
		log.Fatal("Erro: você deve fornecer o ID da tarefa a deletar")
	}

	var taskID int
	_, err := fmt.Sscanf(args[0], "%d", &taskID)
	if err != nil {
		log.Fatalf("Erro: ID deve ser um número. Você passou '%s'\n", args[0])
	}

	mu.Lock()
	defer mu.Unlock()

	currentData, err := ReadJsonFile()
	if err != nil {
		log.Fatalf("Erro ao ler tarefas: %v\n", err)
	}

	var tasks []models.Task
	if len(currentData) > 0 {
		err = json.Unmarshal(currentData, &tasks)
		if err != nil {
			log.Fatalf("Erro ao decodificar JSON: %v\n", err)
		}
	}

	newTasks := make([]models.Task, 0)
	found := false

	for _, t := range tasks {
		if t.ID != taskID {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		log.Fatalf("Erro: tarefa com ID %d não encontrada\n", taskID)
	}

	if len(newTasks) > 0 {
		err = saveTasksToFile(newTasks)
		if err != nil {
			log.Fatalf("Erro ao salvar tarefas: %v\n", err)
		}
	} else {
		// Se não há mais tarefas, limpa o arquivo
		err = os.WriteFile("tasks.json", []byte("[]"), 0644)
		if err != nil {
			log.Fatalf("Erro ao limpar arquivo de tarefas: %v\n", err)
		}
	}

	fmt.Printf("✓ Tarefa %d deletada!\n", taskID)
}

// formatDate formata um time.Time para um formato legível.
//
// Formato: "02 Jan 2006 15:04"
// Exemplo: "21 Dec 2025 14:30"
func formatDate(t time.Time) string {
	return fmt.Sprintf("%s", t.Format("02 Jan 2006 15:04"))
}

// ListFunc lê todas as tarefas do arquivo JSON e as exibe formatadas.
//
// A função:
// 1. Lê o arquivo JSON de tarefas
// 2. Faz o parse dos dados para um slice de Task structs
// 3. Se houver tarefas, exibe cada uma com:
//   - ID
//   - Descrição
//   - Status de conclusão
//   - Data de criação
//   - Data de conclusão (se disponível)
//
// 4. Se não houver tarefas, exibe mensagem apropriada
//
// Em caso de erro ao ler o arquivo, exibe mensagem de erro e encerra com código 1.
func ListFunc() {
	currentData, err := ReadJsonFile()
	if err != nil {
		log.Fatalf("Erro ao ler arquivo de tarefas: %v\n", err)
	}

	var tasks []models.Task
	if len(currentData) > 0 {
		err = json.Unmarshal(currentData, &tasks)
		if err != nil {
			log.Fatalf("Erro ao decodificar arquivo JSON: %v\n", err)
		}
	}

	if len(tasks) == 0 {
		fmt.Println("📭 Nenhuma tarefa encontrada. Use 'create' para adicionar uma.")
		return
	}

	fmt.Println("\n📋 Lista de Tarefas:")
	fmt.Println("==================================================")
	for _, t := range tasks {
		status := "⏳"
		if t.Done {
			status = "✓"
		}

		fmt.Printf("[%d] %s %s\n", t.ID, status, t.Description)
		fmt.Printf("    Criada em: %s\n", formatDate(t.CreatedAt))
		if t.DoneAt != nil {
			fmt.Printf("    Concluída em: %s\n", formatDate(*t.DoneAt))
		}
		fmt.Println("--------------------------------------------------")
	}
}
