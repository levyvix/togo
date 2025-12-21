package cmd

import (
	"os"
	"testing"
)

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
	// Chamado no final para limpeza
	os.Remove("tasks.json")
}
