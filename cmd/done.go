package cmd

import (
	"fmt"
	"levyvix/togo/internal"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Marcar uma tarefa como concluída",
	Long: `Marca uma tarefa como concluída registrando o timestamp de conclusão.

O ID deve ser fornecido como um argumento numérico.

Exemplo:
  togo done 1`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.DoneFuncDB(args)
		if err != nil {
			fmt.Println("Erro:", err)

		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
