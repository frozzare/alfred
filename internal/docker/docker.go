package docker

import (
	"os"
	"runtime"

	api "github.com/fsouza/go-dockerclient"
)

// Config represents the docker configuration.
type Config struct {
	Host string
}

// Docker represents a docker client.
type Docker struct {
	client *api.Client
	host   string
}

// NewDocker creates a new docker client.
func NewDocker(args ...*Config) (*Docker, error) {
	var client *api.Client
	var err error
	var host string

	// Find docker host for local machine.
	if len(args) > 0 && len(args[0].Host) > 0 {
		host = args[0].Host
	} else if os.Getenv("DOCKER_HOST") != "" {
		host = os.Getenv("DOCKER_HOST")
	} else if runtime.GOOS == "windows" {
		host = "http://localhost:2375"
	} else {
		host = "unix:///var/run/docker.sock"
	}

	client, err = api.NewClient(host)

	if err != nil {
		return nil, err
	}

	return &Docker{
		client: client,
		host:   host,
	}, nil
}

// Host will return the docker host that is used.
func (d *Docker) Host() string {
	return d.host
}

// RemoveContainer removes a container by name.
func (d *Docker) RemoveContainer(name string) error {
	container, err := d.findContainer(name)

	if err != nil {
		return err
	}

	return d.removeContainer(container)
}
