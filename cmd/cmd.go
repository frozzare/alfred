package cmd

import (
	"github.com/frozzare/alfred/cli/app"
	_ "github.com/frozzare/alfred/cli/config" // start command
	_ "github.com/frozzare/alfred/cli/proxy"  // proxy commands
	_ "github.com/frozzare/alfred/cli/start"  // start command
	_ "github.com/frozzare/alfred/cli/status" // status command
	_ "github.com/frozzare/alfred/cli/stop"   // stop command
	"github.com/frozzare/alfred/internal/log"
)

const version = "master"

// Execute executes the command line.
func Execute() {
	if err := app.Run(version); err != nil {
		log.Error(err)
	}
}
