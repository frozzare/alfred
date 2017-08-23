package app

import (
	"fmt"
	"strings"

	"github.com/frozzare/alfred/config"
	"github.com/frozzare/alfred/docker"
)

// App interface represents the application.
type App interface {
	Start() error
	Stop() error
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

func (a *app) Start() error {
	err := a.opts.Docker.CreateContainer(&docker.CreateContainerOptions{
		Env:          a.opts.Config.Env,
		Name:         "/" + a.opts.Config.Host,
		Image:        a.opts.Config.Image,
		ExposedPorts: []string{fmt.Sprintf("%d", a.opts.Config.Port)},
		Volumes:      []string{a.opts.Config.Path},
	})

	if err != nil && strings.Contains(err.Error(), "container already exists") {
		return nil
	}

	return err
}

func (a *app) Stop() error {
	return a.opts.Docker.RemoveContainer(a.opts.Config.Host)
}
