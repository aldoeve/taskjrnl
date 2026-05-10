package schema

// Basic struct that represents the Positions table.
type Positions struct {
	// Auto-generated. Id of the position row.
	Id int

	// Tasks.Id that the position is related to.
	TaskId int

	// What position a task is in relative to the others.
	Position int
}
