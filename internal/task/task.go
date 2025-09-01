package task

import (
	"time"

	"github.com/google/uuid"
	"vinna.app/vinna-app/internal/frequency"
)

type RegularTask struct {
	Task   BaseTask `json:"task"`
	Status Status   `json:"status"`
}

func New(userID uuid.UUID, name string, start time.Time, end time.Time, status Status, frequency frequency.Frequency) (RegularTask, error) {

	bt, err := NewBaseTask(userID, name, start, end, frequency)
	if err != nil {
		return RegularTask{}, err
	}
	return RegularTask{Task: *bt, Status: status}, nil
}
