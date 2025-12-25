package cmd

import (
	"fmt"
	"levyvix/togo/internal"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "limpar todas as tarefas do banco de dados",
	Long: `Limpa todas as tarefas do banco de dados

Usage:
	togo clear`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.ClearDB(args)
		if err != nil {
			fmt.Printf("Erro: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
