package cmd

import (
	"levyvix/togo/internal"

	"github.com/spf13/cobra"
)

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
		internal.CreateFuncDB(args)
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
