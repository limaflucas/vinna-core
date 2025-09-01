package habit

import (
	"time"

	"github.com/google/uuid"
	"vinna.app/vinna-app/internal/frequency"
	"vinna.app/vinna-app/internal/task"
)

type Habit struct {
	Task   task.BaseTask `json:"task"`
	Status Status        `json:"status"`
}

func New(userID uuid.UUID, name string, start time.Time, end time.Time, status Status, frequency frequency.Frequency) (Habit, error) {
	bt, err := task.NewBaseTask(userID, name, start, end, frequency)
	if err != nil {
		return Habit{}, err
	}
	return Habit{Task: *bt, Status: status}, nil

}
