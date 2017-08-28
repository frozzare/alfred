package root

import (
	"os"

	"github.com/frozzare/alfred/config"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// Cmd is the root command.
	Cmd = kingpin.New("alfred", "")

	// Command registers a command.
	Command = Cmd.Command

	// Init function.
	Init func() (*config.Config, error)
)

func init() {
	workdir := Cmd.Flag("wd", "Change working directory.").Default(".").Short('W').String()
	path := Cmd.Flag("config", "Path to config file.").Default("alfred.json").Short('C').String()

	Cmd.PreAction(func(ctx *kingpin.ParseContext) error {
		os.Chdir(*workdir)

		g, err := config.ReadGlobalConfig()
		if err != nil {
			return err
		}

		config.SetGlobal(g)

		Init = func() (*config.Config, error) {
			c, err := config.ReadConfig(*path)

			if err != nil {
				return nil, errors.Wrap(err, "Reading config")
			}

			return c, nil
		}

		return nil
	})
}
