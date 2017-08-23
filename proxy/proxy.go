package proxy

import "github.com/frozzare/alfred/docker"

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
	case "caddy":
		fallthrough
	default:
		return newCaddyProxy(opts)
	}
}
