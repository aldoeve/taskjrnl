// Package exitcodes conatians the common exit codes that the application returns to whoever called it.
package exitcodes

const (
	// Application successfully completed its operation.
	ExitOk = iota

	// Any other error.
	ExitError

	// Any user usage and format issues.
	ExitUsage
)
