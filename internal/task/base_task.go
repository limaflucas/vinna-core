package task

import (
	"time"

	"github.com/google/uuid"
	"vinna.app/vinna-app/internal/frequency"
)

type TaskType string

const (
	Habit TaskType = "Habit"
	Task  TaskType = "Task"
)

type BaseTask struct {
	ID         uuid.UUID           `json:"id"`
	UserID     uuid.UUID           `json:"user_id"`
	Name       string              `json:"name"`
	StartDate  time.Time           `json:"start_date"`
	EndDate    time.Time           `json:"end_date"`
	Frequency  frequency.Frequency `json:"frequency"`
	CreatedAt  time.Time           `json:"created_at"`
	ModifiedAt time.Time           `json:"modified_at"`
}

func NewBaseTask(userID uuid.UUID, name string, start time.Time, end time.Time, frequency frequency.Frequency) *BaseTask {
	return &BaseTask{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		StartDate: start,
		EndDate:   end,
		Frequency: frequency,
		CreatedAt: time.Now(),
	}
}
