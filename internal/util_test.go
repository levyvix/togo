package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"levyvix/go-todo-list/models"
)

// setupTest cria um arquivo temporário para testes e retorna função de limpeza
func setupTest(t *testing.T) (cleanup func()) {
	t.Helper()

	// Cria diretório temporário (limpeza automática pelo Go)
	tempDir := t.TempDir()

	// Guarda nome original e define novo
	originalFileName := TasksFileName
	TasksFileName = filepath.Join(tempDir, "tasks.json")

	// Retorna função para restaurar nome original
	return func() {
		TasksFileName = originalFileName
	}
}

// TestReadJsonFileNotExists testa leitura quando arquivo não existe
func TestReadJsonFileNotExists(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	data, err := ReadJsonFile()

	// Não deve retornar erro, apenas slice vazio
	if err != nil {
		t.Errorf("ReadJsonFile() error = %v, want nil", err)
	}

	if len(data) != 0 {
		t.Errorf("ReadJsonFile() returned non-empty data for missing file")
	}
}

// TestReadJsonFileExists testa leitura quando arquivo existe
func TestReadJsonFileExists(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	// Criar arquivo de teste
	testData := []models.Task{
		{ID: 1, Description: "Test", Done: false, CreatedAt: time.Now()},
	}
	jsonData, _ := json.MarshalIndent(testData, "", "  ")
	os.WriteFile(TasksFileName, jsonData, 0644)

	// Ler arquivo
	data, err := ReadJsonFile()

	if err != nil {
		t.Errorf("ReadJsonFile() error = %v, want nil", err)
	}

	if len(data) == 0 {
		t.Errorf("ReadJsonFile() returned empty data for existing file")
	}

	// Validar conteúdo
	var tasks []models.Task
	json.Unmarshal(data, &tasks)
	if len(tasks) != 1 || tasks[0].ID != 1 {
		t.Errorf("ReadJsonFile() content mismatch")
	}
}

// TestWriteToJsonNewFile testa escrita em arquivo novo
func TestWriteToJsonNewFile(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	task := models.Task{
		ID:          1,
		Description: "Nova tarefa",
		Done:        false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}

	err := WriteToJson(task)

	if err != nil {
		t.Errorf("WriteToJson() error = %v", err)
	}

	// Validar que arquivo foi criado
	if _, err := os.Stat(TasksFileName); os.IsNotExist(err) {
		t.Errorf("WriteToJson() did not create file")
	}

	// Validar conteúdo
	data, _ := ReadJsonFile()
	var tasks []models.Task
	json.Unmarshal(data, &tasks)

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

// TestWriteToJsonAppend testa adição em arquivo existente
func TestWriteToJsonAppend(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	// Primeira tarefa
	task1 := models.Task{
		ID:          1,
		Description: "Tarefa 1",
		Done:        false,
		CreatedAt:   time.Now(),
	}
	WriteToJson(task1)

	// Segunda tarefa
	task2 := models.Task{
		ID:          2,
		Description: "Tarefa 2",
		Done:        false,
		CreatedAt:   time.Now(),
	}
	WriteToJson(task2)

	// Validar conteúdo
	data, _ := ReadJsonFile()
	var tasks []models.Task
	json.Unmarshal(data, &tasks)

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].ID != 1 || tasks[1].ID != 2 {
		t.Errorf("Task IDs mismatch")
	}
}

// TestSaveTasksToFileEmpty testa salvar lista vazia
func TestSaveTasksToFileEmpty(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	tasks := []models.Task{}

	err := saveTasksToFile(tasks)

	if err != nil {
		t.Errorf("saveTasksToFile() error = %v", err)
	}

	// Arquivo deve conter JSON array vazio
	data, _ := ReadJsonFile()
	if string(data) != "[]" {
		t.Errorf("Expected '[]', got %s", string(data))
	}
}

// TestSaveTasksToFileMultiple testa salvar múltiplas tarefas
func TestSaveTasksToFileMultiple(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	now := time.Now()
	tasks := []models.Task{
		{ID: 1, Description: "Task 1", Done: false, CreatedAt: now},
		{ID: 2, Description: "Task 2", Done: true, CreatedAt: now, DoneAt: &now},
		{ID: 3, Description: "Task 3", Done: false, CreatedAt: now},
	}

	err := saveTasksToFile(tasks)

	if err != nil {
		t.Errorf("saveTasksToFile() error = %v", err)
	}

	// Validar conteúdo
	data, _ := ReadJsonFile()
	var readTasks []models.Task
	json.Unmarshal(data, &readTasks)

	if len(readTasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(readTasks))
	}

	// Validar segunda tarefa tem DoneAt
	if readTasks[1].DoneAt == nil {
		t.Errorf("Task 2 should have DoneAt timestamp")
	}
}

// TestJsonFileFormatting testa que o arquivo JSON está bem formatado
func TestJsonFileFormatting(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	tasks := []models.Task{
		{ID: 1, Description: "Test", Done: false, CreatedAt: time.Now()},
	}

	saveTasksToFile(tasks)

	// Ler arquivo como string
	data, _ := os.ReadFile(TasksFileName)
	content := string(data)

	// Deve ser indentado (conter espacos)
	if !contains(content, "  ") {
		t.Errorf("JSON file is not properly indented")
	}

	// Deve ser válido JSON
	var parsedTasks []models.Task
	err := json.Unmarshal(data, &parsedTasks)
	if err != nil {
		t.Errorf("JSON file is not valid JSON: %v", err)
	}
}

// TestJsonSpecialCharacters testa descrições com caracteres especiais
func TestJsonSpecialCharacters(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	specialDesc := `Tarefa com "aspas" e 'apóstrofos' e \barra`
	task := models.Task{
		ID:          1,
		Description: specialDesc,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	WriteToJson(task)

	// Ler de volta
	data, _ := ReadJsonFile()
	var tasks []models.Task
	json.Unmarshal(data, &tasks)

	if tasks[0].Description != specialDesc {
		t.Errorf("Special characters lost: got %s", tasks[0].Description)
	}
}

// Helper function para checar se string contém substring
func contains(s, substr string) bool {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
