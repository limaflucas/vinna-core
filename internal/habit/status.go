package habit

type Status string

const (
	Active       Status = "Active"
	Paused       Status = "Paused"
	Inactive     Status = "Inactive"
	Accomplished Status = "Accomplished"
	Abandoned    Status = "Abandoned"
)
