// Package internal contém as funções de lógica de negócio da aplicação.
package internal

import (
	"fmt"
	"levyvix/togo/internal/database"
	"levyvix/togo/schema"
	"strconv"
	"strings"
	"time"
)

func CreateFuncDB(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Comando aceita apenas um argumento. Voce passou %d argumentos\n", len(args))
	}

	descricao := args[0]
	novaTask := schema.Task{
		Description: descricao,
		Done:        false,
		DoneAt:      nil,
	}

	database.DB.Create(&novaTask)
	fmt.Println("Tarefa Criada!")
	return nil
}

func DoneFuncDB(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Esse comando aceita somente 1 argumento. Você passou %d argumentos\n", len(args))
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Voce precisa passar um numero. Voce passou '%v'\n", args[0])
	}
	var t schema.Task
	database.DB.First(&t, id)

	t.Done = true
	now := time.Now()
	t.DoneAt = &now
	database.DB.Save(&t)
	fmt.Println("Tarefa marcada como concluida!")
	return nil

}

func DeleteFuncDB(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Erro: voce deve fornecer o ID da tarefa a deletar\n")
	}

	id := args[0]
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("Voce precisa passar um inteiro para o ID. voce passou %v\n", id)
	}

	result := database.DB.Delete(&schema.Task{}, taskID)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Nao foi possivel encontrar a tarefa com ID %v\n", id)
	}
	fmt.Printf("Tarefa %v deletada com sucesso!\n", id)
	return nil
}

// formatDate formata um time.Time para um formato legível.
//
// Formato: "02 Jan 2006 15:04"
// Exemplo: "21 Dec 2025 14:30"
func formatDate(t time.Time) string {
	ingles := t.Format("02 Jan 2006 15:04")

	repl := strings.NewReplacer(
		"Jan", "Jan",
		"Feb", "Fev",
		"Mar", "Mar",
		"Apr", "Abr",
		"May", "Mai",
		"Jun", "Jun",
		"Jul", "Jul",
		"Aug", "Ago",
		"Sep", "Set",
		"Oct", "Out",
		"Nov", "Nov",
		"Dec", "Dez",
	)

	portugues := repl.Replace(ingles)

	return portugues
}

func ListFuncDB() error {
	var tasks []schema.Task

	result := database.DB.Order("id asc").Find(&tasks)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Nenhuma task para mostrar. Crie uma usando o comando 'create'")
	}

	fmt.Println("\n📋 Lista de Tarefas:")
	fmt.Println("==================================================")
	for _, t := range tasks {
		status := "⏳"
		if t.Done {
			status = "✓"
		}

		fmt.Printf("[%d] %s %s\n", t.ID, status, t.Description)
		fmt.Printf("    Criada em: %s\n", formatDate(t.CreatedAt))
		if t.DoneAt != nil {
			fmt.Printf("    Concluída em: %s\n", formatDate(*t.DoneAt))
		}
		fmt.Println("--------------------------------------------------")
	}
	return nil
}

func EditFuncDB(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Somente ID e Nova Descrição são permitidos. Voce passou %d argumentos\n", len(args))
	}
	id := args[0]
	novaDescricao := args[1]

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("Voce precisa passar um inteiro como ID. voce passou %v\n", id)
	}

	var t schema.Task
	result := database.DB.First(&t, taskID)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Nao achei tarefa com id %v\n", id)
	}

	t.Description = novaDescricao
	database.DB.Save(&t)
	fmt.Println("Tarefa atualizada com sucesso!")
	return nil

}
