package cmd

import (
	"fmt"
	"levyvix/togo/internal"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edita a descricao de uma tarefa",
	Long: `Edita a descrição de uma tarefa

Exemplo:
	togo edit <id> <nova descrição>`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.EditFuncDB(args)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
