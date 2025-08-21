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

func New(userID uuid.UUID, name string, start time.Time, end time.Time, status Status, frequency frequency.Frequency) RegularTask {
	return RegularTask{
		*NewBaseTask(userID, name, start, end, frequency),
		status}
}
