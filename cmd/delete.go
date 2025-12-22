/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"levyvix/togo/internal"

	"github.com/spf13/cobra"
)

// deleteCmd representa o comando para deletar uma tarefa.
// Aceita um argumento: o ID da tarefa.
//
// Exemplos:
//
//	togo delete 1
//	togo delete 5
var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Deletar uma tarefa",
	Long: `Remove permanentemente uma tarefa do sistema.

O ID deve ser fornecido como um argumento numérico.

Exemplo:
  togo delete 1
  togo delete 5`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.DeleteFuncDB(args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
