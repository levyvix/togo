// Package models define as estruturas de dados usadas na aplicação.
package models

import "time"

// Task representa uma tarefa no sistema.
//
// Fields:
//   - ID: Identificador único sequencial da tarefa
//   - Description: Descrição ou título da tarefa (obrigatório)
//   - Done: Status de conclusão da tarefa (false = pendente, true = concluída)
//   - CreatedAt: Timestamp de quando a tarefa foi criada
//   - DoneAt: Timestamp de quando a tarefa foi marcada como concluída (nil se pendente)
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"createdAt"`
	DoneAt      *time.Time `json:"doneAt,omitempty"`
}
