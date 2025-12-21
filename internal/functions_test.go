package internal

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"levyvix/togo/models"
)

// TestGetNextID testa a função que calcula o próximo ID sequencial
func TestGetNextID(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	tests := []struct {
		name        string
		tasks       []models.Task
		expectedID  int
		shouldError bool
	}{
		{
			name:       "Arquivo vazio retorna ID 1",
			tasks:      []models.Task{},
			expectedID: 1,
		},
		{
			name: "Com uma tarefa retorna ID 2",
			tasks: []models.Task{
				{ID: 1, Description: "Task 1", Done: false, CreatedAt: time.Now()},
			},
			expectedID: 2,
		},
		{
			name: "Com múltiplas tarefas retorna max+1",
			tasks: []models.Task{
				{ID: 1, Description: "Task 1", Done: false, CreatedAt: time.Now()},
				{ID: 5, Description: "Task 5", Done: false, CreatedAt: time.Now()},
				{ID: 3, Description: "Task 3", Done: false, CreatedAt: time.Now()},
			},
			expectedID: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Se há tarefas, salvar no arquivo
			if len(tt.tasks) > 0 {
				data, _ := json.MarshalIndent(tt.tasks, "", "  ")
				os.WriteFile(TasksFileName, data, 0644)
			}

			got, err := getNextID()

			if (err != nil) != tt.shouldError {
				t.Errorf("getNextID() erro = %v, shouldError %v", err, tt.shouldError)
			}

			if got != tt.expectedID {
				t.Errorf("getNextID() = %d, want %d", got, tt.expectedID)
			}
		})
	}
}

// TestFormatDate testa a formatação de datas
func TestFormatDate(t *testing.T) {
	testTime := time.Date(2025, 12, 21, 14, 30, 0, 0, time.UTC)
	expected := "21 Dez 2025 14:30"
	got := formatDate(testTime)

	if got != expected {
		t.Errorf("formatDate() = %s, want %s", got, expected)
	}
}

// TestCreateTaskStructure testa se a struct Task é criada corretamente
func TestCreateTaskStructure(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	// Simular criação de tarefa
	nextID := 1
	task := models.Task{
		ID:          nextID,
		Description: "Teste tarefa",
		Done:        false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}

	// Validações
	if task.ID != 1 {
		t.Errorf("Task ID = %d, want 1", task.ID)
	}
	if task.Description != "Teste tarefa" {
		t.Errorf("Task Description = %s, want 'Teste tarefa'", task.Description)
	}
	if task.Done != false {
		t.Errorf("Task Done = %v, want false", task.Done)
	}
	if task.DoneAt != nil {
		t.Errorf("Task DoneAt = %v, want nil", task.DoneAt)
	}
}

// TestTaskMarkedAsDone testa se uma tarefa pode ser marcada como concluída
func TestTaskMarkedAsDone(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	// Criar tarefa
	task := models.Task{
		ID:          1,
		Description: "Teste",
		Done:        false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}

	// Marcar como concluída
	now := time.Now()
	task.Done = true
	task.DoneAt = &now

	// Validações
	if task.Done != true {
		t.Errorf("Task Done = %v, want true", task.Done)
	}
	if task.DoneAt == nil {
		t.Errorf("Task DoneAt = nil, want timestamp")
	}
}

// TestWriteAndReadTask testa salvar e ler uma tarefa
func TestWriteAndReadTask(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	// Criar e salvar tarefa
	task := models.Task{
		ID:          1,
		Description: "Tarefa teste",
		Done:        false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}

	err := saveTasksToFile([]models.Task{task})
	if err != nil {
		t.Fatalf("saveTasksToFile() error = %v", err)
	}

	// Ler tarefa
	data, err := ReadJsonFile()
	if err != nil {
		t.Fatalf("ReadJsonFile() error = %v", err)
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		t.Fatalf("Unmarshal error = %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].ID != 1 || tasks[0].Description != "Tarefa teste" {
		t.Errorf("Task data mismatch: %+v", tasks[0])
	}
}

// TestMultipleTasks testa com múltiplas tarefas
func TestMultipleTasks(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	tasks := []models.Task{
		{ID: 1, Description: "Task 1", Done: false, CreatedAt: time.Now()},
		{ID: 2, Description: "Task 2", Done: false, CreatedAt: time.Now()},
		{ID: 3, Description: "Task 3", Done: false, CreatedAt: time.Now()},
	}

	err := saveTasksToFile(tasks)
	if err != nil {
		t.Fatalf("saveTasksToFile() error = %v", err)
	}

	// Ler todas
	data, err := ReadJsonFile()
	if err != nil {
		t.Fatalf("ReadJsonFile() error = %v", err)
	}

	var readTasks []models.Task
	json.Unmarshal(data, &readTasks)

	if len(readTasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(readTasks))
	}

	for i, task := range readTasks {
		expectedID := i + 1
		if task.ID != expectedID {
			t.Errorf("Task %d: expected ID %d, got %d", i, expectedID, task.ID)
		}
	}
}

// TestRemoveTask testa remoção de tarefa de um slice
func TestRemoveTask(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	tasks := []models.Task{
		{ID: 1, Description: "Task 1", Done: false, CreatedAt: time.Now()},
		{ID: 2, Description: "Task 2", Done: false, CreatedAt: time.Now()},
		{ID: 3, Description: "Task 3", Done: false, CreatedAt: time.Now()},
	}

	// Remover tarefa com ID 2
	newTasks := make([]models.Task, 0)
	for _, t := range tasks {
		if t.ID != 2 {
			newTasks = append(newTasks, t)
		}
	}

	if len(newTasks) != 2 {
		t.Errorf("Expected 2 tasks after removal, got %d", len(newTasks))
	}

	if newTasks[0].ID != 1 || newTasks[1].ID != 3 {
		t.Errorf("Task IDs after removal incorrect: %v", newTasks)
	}
}
