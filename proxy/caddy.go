package proxy

import (
	"github.com/frozzare/alfred/internal/docker"
)

type caddyProxy struct {
	opts *Options
}

func newCaddyProxy(opts *Options) Proxy {
	return &caddyProxy{opts}
}

func (p *caddyProxy) Start() error {
	return p.opts.Docker.CreateContainer(&docker.CreateContainerOptions{
		Image:   "frozzare/caddy-proxy",
		Labels:  map[string]string{"alfred": "true"},
		Name:    "/" + p.opts.Name,
		Ports:   []string{"80:80"},
		Volumes: []string{"/var/run/docker.sock:/tmp/docker.sock:ro"},
	})
}

func (p *caddyProxy) Stop() error {
	return p.opts.Docker.RemoveContainer(p.opts.Name)
}
