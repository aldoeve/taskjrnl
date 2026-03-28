// Pacage app contains the core application logic.
package app

import (
	"database/sql"
	"fmt"
	util "taskjrnl/pkg/util"
)

func AddMode(db *sql.DB) error {
	user_input := util.ArgsAfterKeyword()
	fmt.Println(user_input, len(user_input))
	return nil
}
