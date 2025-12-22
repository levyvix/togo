package cmd

import (
	"levyvix/togo/internal"
	"log"

	"github.com/spf13/cobra"
)

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
  togo list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			log.Fatalf("Erro: esse comando nao aceita argumentos. voce passou %d argumentos\n", len(args))
		}
		internal.ListFuncDB()
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
