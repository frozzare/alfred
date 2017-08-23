package proxy

import (
	"strings"

	"github.com/frozzare/alfred/docker"
)

type caddyProxy struct {
	opts *Options
}

func newCaddyProxy(opts *Options) Proxy {
	return &caddyProxy{opts}
}

func (p *caddyProxy) Start() error {
	err := p.opts.Docker.CreateContainer(&docker.CreateContainerOptions{
		Name:    "/" + p.opts.Name,
		Image:   "frozzare/caddy-proxy",
		Ports:   []string{"80:80"},
		Volumes: []string{"/var/run/docker.sock:/tmp/docker.sock:ro"},
	})

	if err != nil && strings.Contains(err.Error(), "container already exists") {
		return nil
	}

	return err
}

func (p *caddyProxy) Stop() error {
	return p.opts.Docker.RemoveContainer(p.opts.Name)
}
