package start

import (
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/app"
	"github.com/frozzare/alfred/cli/root"
	"github.com/frozzare/alfred/internal/docker"
	"github.com/frozzare/alfred/internal/log"
	"github.com/pkg/errors"
)

func init() {
	cmd := root.Command("start", "Start application container")

	cmd.Action(func(p *kingpin.ParseContext) error {
		log.Info("Starting application container")

		c, err := root.Init()
		if err != nil {
			return err
		}

		d, err := docker.NewDocker(&docker.Config{
			Host: c.Global().DockerHost,
		})
		if err != nil {
			return errors.Wrap(err, "Docker")
		}

		app := app.NewApp(&app.Options{
			Config: c,
			Docker: d,
		})

		if err := app.Start(); err != nil {
			if strings.Contains(err.Error(), "container already exists") {
				return errors.New("Application container already exists")
			}

			return errors.Wrap(err, "Docker")
		}

		log.Info("Application started at %s", app.URL())

		return nil
	})
}
