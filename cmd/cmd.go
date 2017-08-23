package cmd

import (
	"github.com/frozzare/alfred/cli/app"
	_ "github.com/frozzare/alfred/cli/config" // start command
	_ "github.com/frozzare/alfred/cli/proxy"  // proxy commands
	_ "github.com/frozzare/alfred/cli/start"  // start command
)

const version = "master"

// Execute executes the command line.
func Execute() {
	app.Run(version)
}
