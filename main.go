package main

import (
	"levyvix/togo/cmd"
	"levyvix/togo/internal/database"
)

func main() {
	database.InitDB()
	cmd.Execute()
}
