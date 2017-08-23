package proxy

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/cli/root"
	"github.com/frozzare/alfred/docker"
	p "github.com/frozzare/alfred/proxy"
)

func init() {
	start := root.Command("proxy start", "Start proxy container")

	typ := start.Flag("type", "Proxy type (supports: caddy)").Default("caddy").Short('t').String()

	start.Action(func(_ *kingpin.ParseContext) error {
		d, err := docker.NewDocker()
		if err != nil {
			return err
		}

		proxy := p.NewProxy(&p.Options{
			Docker: d,
			Name:   "alfred_proxy",
			Type:   *typ,
		})

		if err := proxy.Start(); err != nil {
			return err
		}

		return nil
	})

	stop := root.Command("proxy stop", "Stop proxy container")

	stop.Action(func(_ *kingpin.ParseContext) error {
		d, err := docker.NewDocker()
		if err != nil {
			return err
		}

		proxy := p.NewProxy(&p.Options{
			Docker: d,
			Name:   "alfred_proxy",
			Type:   *typ,
		})

		if err := proxy.Stop(); err != nil {
			return err
		}

		return nil
	})
}
