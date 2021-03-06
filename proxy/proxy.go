package proxy

import "github.com/frozzare/alfred/internal/docker"

// Proxy interface represents the application.
type Proxy interface {
	Start() error
	Stop() error
}

// Options represents the proxy options.
type Options struct {
	Docker *docker.Docker
	Name   string
	Type   string
}

// NewProxy creates a new proxy with the given options.
func NewProxy(opts *Options) Proxy {
	switch opts.Type {
	case "nginx":
		return newNginxProxy(opts)
	case "caddy":
		fallthrough
	default:
		return newCaddyProxy(opts)
	}
}
