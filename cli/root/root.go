package root

import (
	"os"

	"github.com/frozzare/alfred/config"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Cmd     = kingpin.New("alfred", "")
	Command = Cmd.Command
	Init    func() (*config.Config, error)
)

func init() {
	workdir := Cmd.Flag("chdir", "Change working directory.").Default(".").Short('C').String()

	Cmd.PreAction(func(ctx *kingpin.ParseContext) error {
		os.Chdir(*workdir)

		Init = func() (*config.Config, error) {
			c, err := config.ReadConfig("alfred.json")

			if err != nil {
				return nil, errors.Wrap(err, "reading config")
			}

			return c, nil
		}

		return nil
	})
}
