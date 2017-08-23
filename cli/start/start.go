package start

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/app"
	"github.com/frozzare/alfred/cli/root"
	"github.com/frozzare/alfred/docker"
)

func init() {
	cmd := root.Command("start", "Start application container")

	cmd.Action(func(_ *kingpin.ParseContext) error {
		c, err := root.Init()
		if err != nil {
			return err
		}

		d, err := docker.NewDocker()
		if err != nil {
			return err
		}

		app := app.NewApp(&app.Options{
			Config: c,
			Docker: d,
		})

		if err := app.Start(); err != nil {
			return err
		}

		return nil
	})
}
