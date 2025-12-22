package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "Gerenciador de tarefas em linha de comando",
	Long: `togo é uma aplicação CLI simples para gerenciar sua lista de tarefas.

Comandos disponíveis:
  create <descrição>  - Criar uma nova tarefa
  list                - Listar todas as tarefas
  done <id>           - Marcar uma tarefa como concluída
  delete <id>         - Deletar uma tarefa
	edit <id> <nova descricao> - Editar a descricao de uma tarefa

Exemplos:
  togo create "Estudar Go"
  togo list
  togo done 1
  togo delete 2
	togo edit 1 "nova descricao"


Use "togo [command] --help" para mais informações sobre um comando.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-todo-list.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
