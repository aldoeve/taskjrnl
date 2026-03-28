// Package schema holds schema structs representing the whole database.
package schema

// Pointer feilds are treated as optional args that the user does not have to fill out.
type Tasks struct {
	// Auto-set, Not Null
	Id int
	// Not Null
	Name string
	//Auto-set, Not Null
	Create_date string
	// Not Null
	Priority *string
	// Not Null
	Importance_variance *int
}
