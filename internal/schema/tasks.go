// Package schema holds schema structs representing the whole database.
package schema

// Pointer feilds are treated as optional args that the user does not have to fill out.
type Tasks struct {
	Id                  int     // Auto-set, Not Null
	Name                string  // Not Null
	Tag                 *string // Nullable
	Create_date         string  //Auto-set, Not Null
	Priority            *string // Not Null
	Importance_variance int     // Not Null
}
