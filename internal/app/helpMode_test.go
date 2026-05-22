package app_test

import (
	"database/sql"
	"taskjrnl/internal/app"
	testingutils "taskjrnl/testingUtils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_screenOutput(t *testing.T) {
	correctOutput := `taskjrnl — a simple command line task & journal                          
                                                                         
                                                                         
Usage:                                                                   
                                                                         
    taskjrnl | task [options] <command>                                  
                                                                         
                                                                         
Commands:                                                                
    help            Show help                                            
    add             Adds task. <taskName> [priority:L|M|H] [tag:"string"]
    list            Lists all tasks                                      
                                                                         
                                                                         
Options:                                                                 
    -h, --help      Show help                                            
                                                                         
                                                                         
`

	output := testingutils.CaptureOutput(t, app.HelpMode, (*sql.DB)(nil))

	assert.Equal(t, correctOutput, output)
}
