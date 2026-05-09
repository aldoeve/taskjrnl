// Package schema holds schema structs representing the whole database.
package schema

// Basic struct that represents the table Tasks.
// Pointer feilds are treated as optional args that the user does not have to fill out.
type Tasks struct {
	// Auto-set, Not Null. Id of task.
	Id int

	// Not Null. Name of task.
	Name string

	// Nullable, defaults to empty str. Task tag.
	Tag *string

	// Auto-set, Not Null. Date task was created.
	DateCreated string

	// Not Null, defaults to "L". Priority rank of the task.
	Priority *string

	// Not Null. The importance modifier that can change a tasks importance relative to others.
	ImportanceVariance int
}
