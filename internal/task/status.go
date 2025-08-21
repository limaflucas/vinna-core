package task

type Status string

const (
	Recorded Status = "Logged"
	Todo     Status = "To-do"
	Doing    Status = "Doing"
	Done     Status = "Done"
	Canceled Status = "Canceled"
)
