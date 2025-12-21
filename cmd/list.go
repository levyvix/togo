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

// listCmd representa o comando para listar todas as tarefas.
// Não aceita nenhum argumento.
//
// Exibe todas as tarefas em ordem de criação, mostrando:
// - ID sequencial único
// - Status (✓ concluída ou ⏳ pendente)
// - Descrição
// - Data de criação
// - Data de conclusão (se aplicável)
//
// Exemplos:
//
//	go-todo-list list
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas as tarefas",
	Long: `Exibe todas as tarefas salvas com seus detalhes:
- ID sequencial único
- Status (✓ = concluída, ⏳ = pendente)
- Descrição
- Data de criação
- Data de conclusão (se aplicável)

Este comando não aceita argumentos.

Exemplo:
  go-todo-list list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Printf("Erro: este comando não aceita argumentos. Você passou %d argumentos\n", len(args))
			os.Exit(1)
		}
		internal.ListFunc()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
