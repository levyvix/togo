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
}
