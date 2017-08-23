package app

import (
	"os"

	"github.com/frozzare/alfred/cli/root"
)

func Run(version string) error {
	root.Cmd.Version(version)
	_, err := root.Cmd.Parse(os.Args[1:])
	return err
}
