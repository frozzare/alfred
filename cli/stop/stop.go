package stop

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/app"
	"github.com/frozzare/alfred/cli/root"
	"github.com/frozzare/alfred/config"
	"github.com/frozzare/alfred/internal/docker"
	"github.com/frozzare/alfred/internal/log"
	"github.com/pkg/errors"
)

func init() {
	cmd := root.Command("stop", "Stop application container")

	cmd.Action(func(_ *kingpin.ParseContext) error {
		log.Info("Stopping application container")

		c, err := root.Init()
		if err != nil {
			return err
		}

		d, err := docker.NewDocker(&docker.Config{
			Host: config.Global().DockerHost,
		})
		if err != nil {
			return errors.Wrap(err, "Docker")
		}

		app := app.NewApp(&app.Options{
			Config: c,
			Docker: d,
		})

		if err := app.Stop(); err != nil {
			return errors.Wrap(err, "Docker")
		}

		log.Info("Application container stopped")

		return nil
	})
}
