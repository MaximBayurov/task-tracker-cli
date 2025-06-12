package storage

import (
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
	"time"
)

// DataRow describe stored data about task
type DataRow struct {
	ID          int64
	Description string
	Status      mark.Mark
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
