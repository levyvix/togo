/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/

// Package cmd define os comandos CLI da aplicação usando o framework Cobra.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd representa o comando raiz da aplicação.
// É o ponto de entrada para todos os subcomandos (create, list, done, delete).
var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "Gerenciador de tarefas em linha de comando",
	Long: `togo é uma aplicação CLI simples para gerenciar sua lista de tarefas.

Comandos disponíveis:
  create <descrição>  - Criar uma nova tarefa
  list                - Listar todas as tarefas
  done <id>           - Marcar uma tarefa como concluída
  delete <id>         - Deletar uma tarefa

Exemplos:
  togo create "Estudar Go"
  togo list
  togo done 1
  togo delete 2

Use "togo [command] --help" para mais informações sobre um comando.`,
}

// Execute adiciona todos os subcomandos ao comando raiz e configura as flags.
// É chamado uma única vez por main.main().
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
