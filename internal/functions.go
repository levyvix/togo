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

const (
	MsgTaskCreated = "Tarefa criada com sucesso."
	MsgTaskDone    = "Tarefa marcada como concluida."
	MsgTaskDeleted = "Tarefa deletada com sucesso."
	MsgTaskUpdated = "Tarefa atualizada com sucesso."
)

func CreateFuncDB(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("este comando aceita apenas um argumento, você passou %d", len(args))
	}

	if args[0] == "" || strings.TrimSpace(args[0]) == "" {
		return fmt.Errorf("a descrição não pode estar vazia")
	}

	descricao := args[0]
	novaTask := schema.Task{
		Description: descricao,
		Done:        false,
		DoneAt:      nil,
	}

	database.DB.Create(&novaTask)
	fmt.Println(MsgTaskCreated)
	return nil
}

func DoneFuncDB(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("este comando aceita apenas um argumento, você passou %d argumentos", len(args))
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("nao foi possivel convertar '%v' para inteiro", args[0])
	}

	var t schema.Task
	result := database.DB.First(&t, id)
	if result.Error != nil {
		return fmt.Errorf("tarefa com ID %d não existe: %w", id, result.Error)
	}

	// tarefa já está feita?
	if t.Done {
		return fmt.Errorf("tarefa %d já está concluída", id)
	}

	t.Done = true
	now := time.Now()
	t.DoneAt = &now
	result = database.DB.Save(&t)
	if result.Error != nil {
		return fmt.Errorf("erro ao salvar a tarefa no banco de dados: %w", result.Error)
	}
	fmt.Println(MsgTaskDone)
	return nil
}

func DeleteFuncDB(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("este comando aceita apenas um argumento, você passou %d", len(args))
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("o ID deve ser um número inteiro, você passou '%v'", args[0])
	}

	result := database.DB.Delete(&schema.Task{}, taskID)
	if result.RowsAffected == 0 {
		return fmt.Errorf("tarefa com ID %d não existe", taskID)
	}
	fmt.Println(MsgTaskDeleted)
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
		fmt.Println("nenhuma task para mostrar. Crie uma usando o comando 'create'")
		return nil
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
		return fmt.Errorf("este comando aceita dois argumentos (ID e descrição), você passou %d", len(args))
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("o ID deve ser um número inteiro, você passou '%v'", args[0])
	}

	novaDescricao := args[1]
	if novaDescricao == "" || strings.TrimSpace(novaDescricao) == "" {
		return fmt.Errorf("a descrição não pode estar vazia")
	}

	var t schema.Task
	result := database.DB.First(&t, taskID)
	if result.Error != nil {
		return fmt.Errorf("tarefa com ID %d não existe: %w", taskID, result.Error)
	}

	t.Description = novaDescricao
	result = database.DB.Save(&t)
	if result.Error != nil {
		return fmt.Errorf("erro ao salvar a tarefa no banco de dados: %w", result.Error)
	}

	fmt.Println(MsgTaskUpdated)
	return nil
}
