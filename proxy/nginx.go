package proxy

import (
	"strings"

	"github.com/frozzare/alfred/docker"
)

type nginxProxy struct {
	opts *Options
}

func newNginxProxy(opts *Options) Proxy {
	return &nginxProxy{opts}
}

func (p *nginxProxy) Start() error {
	err := p.opts.Docker.CreateContainer(&docker.CreateContainerOptions{
		Name:    "/" + p.opts.Name,
		Image:   "jwilder/nginx-proxy",
		Ports:   []string{"80:80"},
		Volumes: []string{"/var/run/docker.sock:/tmp/docker.sock:ro"},
	})

	if err != nil && strings.Contains(err.Error(), "container already exists") {
		return nil
	}

	return err
}

func (p *nginxProxy) Stop() error {
	return p.opts.Docker.RemoveContainer(p.opts.Name)
}
