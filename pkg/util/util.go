// Package util holds commonly used utility functions.
package util

import "flag"

// This function assumes the first token is a keyword and discards it.
func ArgsAfterKeyword() []string {
	return flag.Args()[1:]
}
