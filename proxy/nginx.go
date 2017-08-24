package proxy

import (
	"github.com/frozzare/alfred/internal/docker"
)

type nginxProxy struct {
	opts *Options
}

func newNginxProxy(opts *Options) Proxy {
	return &nginxProxy{opts}
}

func (p *nginxProxy) Start() error {
	return p.opts.Docker.CreateContainer(&docker.CreateContainerOptions{
		Image:   "jwilder/nginx-proxy",
		Labels:  map[string]string{"alfred": "true"},
		Name:    "/" + p.opts.Name,
		Ports:   []string{"80:80"},
		Volumes: []string{"/var/run/docker.sock:/tmp/docker.sock:ro"},
	})
}

func (p *nginxProxy) Stop() error {
	return p.opts.Docker.RemoveContainer(p.opts.Name)
}
