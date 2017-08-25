package app

import (
	"fmt"

	"github.com/frozzare/alfred/config"
	"github.com/frozzare/alfred/internal/docker"
)

// App interface represents the application.
type App interface {
	Start() error
	Stop() error
	URL() string
}

// Options represents the application options.
type Options struct {
	Config *config.Config
	Docker *docker.Docker
}

type app struct {
	opts *Options
}

// NewApp creates a new app with the given options.
func NewApp(opts *Options) App {
	return &app{opts}
}

// Start application container.
func (a *app) Start() error {
	return a.opts.Docker.CreateContainer(&docker.CreateContainerOptions{
		Env:          a.opts.Config.Env,
		ExposedPorts: []string{fmt.Sprintf("%d", a.opts.Config.Port)},
		Image:        a.opts.Config.Image,
		Labels:       map[string]string{"alfred": "true"},
		Links:        a.opts.Config.Links,
		Name:         "/" + a.opts.Config.Host,
		Volumes:      append([]string{a.opts.Config.Path}, a.opts.Config.Files...),
	})
}

// Stop and remove application container.
func (a *app) Stop() error {
	return a.opts.Docker.RemoveContainer(a.opts.Config.Host)
}

// Get application URL.
func (a *app) URL() string {
	return fmt.Sprintf("http://%s", a.opts.Config.Host)
}
