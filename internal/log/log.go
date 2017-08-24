package log

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stderr, "", log.LstdFlags)
}

// Error outputs a error row and exits with code 1 if no second bool parameter is provided.
func Error(err error, exit ...bool) {
	red := color.New(color.FgRed).SprintFunc()
	msg := fmt.Sprintf("%s %s", red("==>"), err.Error())
	logger.Println(msg)
	if len(exit) == 0 {
		os.Exit(1)
	}
}

// Info outputs a info row.
func Info(str string, a ...interface{}) {
	green := color.New(color.FgGreen).SprintFunc()
	msg := fmt.Sprintf("%s %s", green("==>"), fmt.Sprintf(str, a...))
	logger.Println(msg)
}
