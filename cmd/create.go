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

// createCmd representa o comando para criar uma nova tarefa.
// Aceita um argumento obrigatório: a descrição da tarefa.
//
// A descrição não pode estar vazia e será armazenada com um ID sequencial único.
//
// Exemplos:
//
//	togo create "Estudar Go"
//	togo create "Fazer compras"
var createCmd = &cobra.Command{
	Use:   "create <descrição>",
	Short: "Criar uma nova tarefa",
	Long: `Cria uma nova tarefa com um ID sequencial único e a persiste no arquivo JSON.

A descrição deve ser fornecida como um argumento de string e não pode estar vazia.
A tarefa é criada com status pendente (não concluída).

Exemplos:
  togo create "Estudar Go"
  togo create "Fazer compras"
  togo create "Revisar código"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Erro: você deve fornecer uma descrição para a tarefa")
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Printf("Erro: apenas 1 argumento é permitido. Você passou %d argumentos\n", len(args))
			os.Exit(1)
		}

		description := args[0]
		if description == "" {
			fmt.Println("Erro: a descrição da tarefa não pode estar vazia")
			os.Exit(1)
		}

		internal.CreateFunc(args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
