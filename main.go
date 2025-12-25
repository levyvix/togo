package main

import (
	"levyvix/togo/cmd"
	"levyvix/togo/internal/database"
	"log"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}
	cmd.Execute()
}
