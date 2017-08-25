package proxy

import (
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/cli/root"
	"github.com/frozzare/alfred/config"
	"github.com/frozzare/alfred/internal/docker"
	"github.com/frozzare/alfred/internal/log"
	p "github.com/frozzare/alfred/proxy"
	"github.com/pkg/errors"
)

func init() {
	cmd := root.Command("proxy", "Proxy container")

	typ := cmd.Flag("type", "Proxy type (supports: caddy, nginx)").Default("caddy").Short('t').String()

	cmd.Command("start", "Start proxy container").Action(func(_ *kingpin.ParseContext) error {
		log.Info("Starting proxy container")

		d, err := docker.NewDocker(&docker.Config{
			Host: config.Global().Docker.Host,
		})
		if err != nil {
			return errors.Wrap(err, "Docker")
		}

		proxy := p.NewProxy(&p.Options{
			Docker: d,
			Name:   "alfred_proxy",
			Type:   *typ,
		})

		if err := proxy.Start(); err != nil {
			if strings.Contains(err.Error(), "container already exists") {
				return errors.New("Proxy container already exists")
			}

			return errors.Wrap(err, "Docker")
		}

		log.Info("Proxy container started")

		return nil
	})

	cmd.Command("stop", "Stop proxy container").Action(func(_ *kingpin.ParseContext) error {
		log.Info("Stopping proxy container")

		d, err := docker.NewDocker(&docker.Config{
			Host: config.Global().Docker.Host,
		})
		if err != nil {
			return errors.Wrap(err, "Docker")
		}

		proxy := p.NewProxy(&p.Options{
			Docker: d,
			Name:   "alfred_proxy",
			Type:   *typ,
		})

		if err := proxy.Stop(); err != nil {
			return errors.Wrap(err, "Docker")
		}

		log.Info("Proxy container stopped")

		return nil
	})
}
