/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/

// Package main é o ponto de entrada da aplicação go-todo-list.
// Inicializa e executa a aplicação CLI usando o framework Cobra.
package main

import (
	"levyvix/togo/cmd"
	"levyvix/togo/internal/database"
)

// main é a função de entrada da aplicação.
// Chama cmd.Execute() para iniciar a interface CLI.
func main() {
	database.InitDB()
	cmd.Execute()
}
