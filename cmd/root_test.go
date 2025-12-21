package cmd

import (
	"path/filepath"
	"testing"

	"levyvix/go-todo-list/internal"
)

// setupTest cria um arquivo temporário para testes e retorna função de limpeza
func setupTest(t *testing.T) (cleanup func()) {
	t.Helper()

	// Cria diretório temporário (limpeza automática pelo Go)
	tempDir := t.TempDir()

	// Guarda nome original e define novo
	originalFileName := internal.TasksFileName
	internal.TasksFileName = filepath.Join(tempDir, "tasks.json")

	// Retorna função para restaurar nome original
	return func() {
		internal.TasksFileName = originalFileName
	}
}

// TestRootCmdHelp testa se o help funciona
func TestRootCmdHelp(t *testing.T) {
	// Executar com --help
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()

	// --help não deve causar erro (Cobra retorna nil para help)
	if err != nil && err.Error() != "pflag: help requested" {
		t.Errorf("rootCmd.Execute() error = %v", err)
	}
}

// TestRootCmdSubcommands verifica se subcomandos estão registrados
func TestRootCmdSubcommands(t *testing.T) {
	expectedCommands := []string{"create", "list", "done", "delete"}

	// Pegar comandos do rootCmd
	commands := rootCmd.Commands()

	foundCommands := make(map[string]bool)
	for _, cmd := range commands {
		foundCommands[cmd.Name()] = true
	}

	for _, expected := range expectedCommands {
		if !foundCommands[expected] {
			t.Errorf("Expected command '%s' not found", expected)
		}
	}
}

// TestCleanupTasksJsonFile limpa o arquivo de teste após os testes
func TestCleanupTasksJsonFile(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()
}
