/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"levyvix/togo/internal"
	"os"

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
  togo delete 1`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Erro: você deve fornecer o ID da tarefa a deletar")
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Printf("Erro: apenas 1 argumento é permitido. Você passou %d argumentos\n", len(args))
			os.Exit(1)
		}
		internal.DeleteFunc(args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
