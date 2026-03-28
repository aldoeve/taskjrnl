// Package schema holds schema structs representing the whole database.
package schema

type Task struct {
	// Auto-incremented, Not Null
	id int
	// Not Null
	name string
	// Not Null
	create_date string
	// Not Null
	priority rune
	// Not Null
	importance int
}
