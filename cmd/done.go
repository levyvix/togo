/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"levyvix/go-todo-list/internal"
	"os"

	"github.com/spf13/cobra"
)

// doneCmd representa o comando para marcar uma tarefa como concluída.
// Aceita um argumento: o ID da tarefa.
//
// Exemplos:
//
//	go-todo-list done 1
//	go-todo-list done 5
var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Marcar uma tarefa como concluída",
	Long: `Marca uma tarefa como concluída registrando o timestamp de conclusão.

O ID deve ser fornecido como um argumento numérico.

Exemplo:
  go-todo-list done 1`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Erro: você deve fornecer o ID da tarefa a marcar como concluída")
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Printf("Erro: apenas 1 argumento é permitido. Você passou %d argumentos\n", len(args))
			os.Exit(1)
		}
		internal.DoneFunc(args)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
