package docker

import (
	"os"
	"runtime"

	api "github.com/fsouza/go-dockerclient"
)

// Docker represents a docker client.
type Docker struct {
	client *api.Client
	host   string
}

// NewDocker creates a new docker client.
func NewDocker() (*Docker, error) {
	var client *api.Client
	var err error
	var host string

	// Find docker host for local machine.
	if os.Getenv("DOCKER_HOST") != "" {
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

// Prune removes all unused containers, volumes, networks and images (both dangling and unreferenced).
func (d *Docker) Prune() error {
	if _, err := d.client.PruneContainers(api.PruneContainersOptions{}); err != nil {
		return err
	}

	if _, err := d.client.PruneImages(api.PruneImagesOptions{}); err != nil {
		return err
	}

	if _, err := d.client.PruneVolumes(api.PruneVolumesOptions{}); err != nil {
		return err
	}

	if _, err := d.client.PruneNetworks(api.PruneNetworksOptions{}); err != nil {
		return err
	}

	return nil
}

func (d *Docker) RemoveContainer(name string) error {
	container, err := d.findContainer(name)

	if err != nil {
		return err
	}

	return d.removeContainer(container)
}
