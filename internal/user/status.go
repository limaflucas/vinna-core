package user

type Status string

const (
	Active  Status = "Active"
	Blocked Status = "Blocked"
	Deleted Status = "Deleted"
)
