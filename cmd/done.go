/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"levyvix/togo/internal"

	"github.com/spf13/cobra"
)

// doneCmd representa o comando para marcar uma tarefa como concluída.
// Aceita um argumento: o ID da tarefa.
//
// Exemplos:
//
//	togo done 1
//	togo done 5
var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Marcar uma tarefa como concluída",
	Long: `Marca uma tarefa como concluída registrando o timestamp de conclusão.

O ID deve ser fornecido como um argumento numérico.

Exemplo:
  togo done 1`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.DoneFuncDB(args)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
