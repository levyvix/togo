package schema

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	Done        bool
	DoneAt      *time.Time
}
