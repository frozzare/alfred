package start

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/app"
	"github.com/frozzare/alfred/cli/root"
	"github.com/frozzare/alfred/internal/docker"
	"github.com/frozzare/alfred/internal/log"
	"github.com/pkg/errors"
)

func init() {
	cmd := root.Command("start", "Start application container")

	cmd.Action(func(_ *kingpin.ParseContext) error {
		log.Info("Starting application container")

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
			return errors.Wrap(err, "Docker")
		}

		log.Info("Application started at http://%s", c.Host)

		return nil
	})
}
