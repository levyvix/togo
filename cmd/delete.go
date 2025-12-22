package cmd

import (
	"levyvix/togo/internal"

	"github.com/spf13/cobra"
)

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
