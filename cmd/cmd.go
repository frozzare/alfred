package cmd

import (
	"os"

	_ "github.com/frozzare/alfred/cli/config" // config command
	_ "github.com/frozzare/alfred/cli/proxy"  // proxy command
	"github.com/frozzare/alfred/cli/root"
	_ "github.com/frozzare/alfred/cli/start"  // start command
	_ "github.com/frozzare/alfred/cli/status" // status command
	_ "github.com/frozzare/alfred/cli/stop"   // stop command
	_ "github.com/frozzare/alfred/cli/url"    // url command
	"github.com/frozzare/alfred/internal/log"
)

// Execute executes the command line.
func Execute(version string) {
	root.Cmd.Version(version)

	if _, err := root.Cmd.Parse(os.Args[1:]); err != nil {
		log.Error(err)
	}
}
