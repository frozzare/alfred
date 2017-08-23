package docker

import (
	"archive/tar"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	api "github.com/fsouza/go-dockerclient"
)

// CreateContainerOptions is the options for creating a container.
type CreateContainerOptions struct {
	Env          []string
	ExposedPorts []string
	IP           string
	Image        string
	Labels       map[string]string
	Name         string
	Recreate     bool
	Ports        []string
	Volumes      []string
}

// createOptions will create container options struct
func createOptions(opts *CreateContainerOptions) api.CreateContainerOptions {
	ip := "0.0.0.0"
	publishedPorts := map[api.Port][]api.PortBinding{}
	for _, port := range opts.Ports {
		parts := strings.Split(port, ":")

		if len(parts) < 2 {
			continue
		}

		containerPort := api.Port(parts[1])

		publishedPorts[containerPort+"/tcp"] = []api.PortBinding{{HostIP: ip, HostPort: parts[0]}}
		publishedPorts[containerPort+"/udp"] = []api.PortBinding{{HostIP: ip, HostPort: parts[0]}}
	}

	exposedPorts := map[api.Port]struct{}{}
	for _, port := range opts.ExposedPorts {
		containerPort := api.Port(port)
		exposedPorts[containerPort+"/tcp"] = struct{}{}
		exposedPorts[containerPort+"/udp"] = struct{}{}
	}

	options := api.CreateContainerOptions{
		Name: opts.Name,
		Config: &api.Config{
			Env:          opts.Env,
			Image:        opts.Image,
			Volumes:      map[string]struct{}{},
			ExposedPorts: exposedPorts,
			Labels:       opts.Labels,
		},
		HostConfig: &api.HostConfig{
			Binds:           []string{},
			PublishAllPorts: false,
			PortBindings:    publishedPorts,
			RestartPolicy:   api.AlwaysRestart(),
		},
	}

	for _, volume := range opts.Volumes {
		parts := strings.Split(volume, ":")

		if len(parts) >= 2 {
			if string(parts[0][0]) == "." {
				path, err := os.Getwd()
				if err != nil {
					continue
				}

				volume = filepath.Join(path, volume)
			}

			options.HostConfig.Binds = append(options.HostConfig.Binds, volume)
			options.Config.Volumes[parts[1]] = struct{}{}
		} else {
			options.Config.Volumes[volume] = struct{}{}
		}
	}

	return options
}

// findContainer will find a container with a name.
func (d *Docker) findContainer(name string) (*api.Container, error) {
	containers, err := d.client.ListContainers(api.ListContainersOptions{
		All: true,
	})

	if err != nil {
		return nil, err
	}

	containerName := name
	if containerName[0] != '/' {
		containerName = "/" + containerName
	}

	for _, container := range containers {
		found := false
		for _, name := range container.Names {
			if name == containerName {
				found = true
				break
			}
		}

		if !found {
			continue
		}

		container, err := d.client.InspectContainer(container.ID)
		if err != nil {
			return nil, fmt.Errorf("Failed to inspect container %s: %s", container.ID, err)
		}

		return container, nil
	}

	return nil, nil
}

// CreateContainer will create a container with the given options.
func (d *Docker) CreateContainer(opts *CreateContainerOptions) error {
	// Check if image exists or pull it.
	d.pullImage(opts.Image)

CREATE:
	// Create container if it don't exists.
	container, err := d.client.CreateContainer(createOptions(opts))

	if err != nil {
		// Try to destroy the container if it exists and should be recreated.
		if strings.Contains(err.Error(), "container already exists") && opts.Recreate {
			container, err := d.findContainer(opts.Name)
			if err != nil {
				return err
			}

			if err := d.removeContainer(container); err != nil {
				return err
			}

			goto CREATE
		}

		return err
	}

	return d.startContainer(container)
}

// startContainer will start the container or try to start the container five times before it stops.
func (d *Docker) startContainer(container *api.Container) error {
	attempted := 0
START:
	if err := d.client.StartContainer(container.ID, nil); err != nil {
		// If it is a 500 error it is likely we can retry and be successful.
		if strings.Contains(err.Error(), "API error (500)") {
			if attempted < 5 {
				attempted++
				time.Sleep(1 * time.Second)
				goto START
			}
		}

		return err
	}

	return nil
}

// stopContainer will stop the container or try to stop the container five times before it stops.
func (d *Docker) stopContainer(container *api.Container) error {
	attempted := 0
STOP:
	if err := d.client.StopContainer(container.ID, 0); err != nil {
		if strings.Contains(err.Error(), "API error (500)") {
			if attempted < 5 {
				attempted++
				time.Sleep(1 * time.Second)
				goto STOP
			}
		}

		if strings.Contains(strings.ToLower(err.Error()), "container not running") {
			return nil
		}

		return err
	}

	return nil
}

// removeContainer will remove the container or try to remove the container five times before it stops.
func (d *Docker) removeContainer(container *api.Container) error {
	// No need to remove nil container.
	if container == nil {
		return nil
	}

	attempted := 0
REMOVE:
	if err := d.client.RemoveContainer(api.RemoveContainerOptions{
		ID:    container.ID,
		Force: true,
	}); err != nil {
		if strings.Contains(err.Error(), "API error (500)") {
			if attempted < 5 {
				attempted++
				time.Sleep(1 * time.Second)
				goto REMOVE
			}
		}

		return err
	}

	return nil
}

type File struct {
	Name string
	Body string
}

func (d *Docker) UploadToContainer(name string, files []File) error {
	buf := new(bytes.Buffer)

	// Create a new tar archive.
	tw := tar.NewWriter(buf)

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
	// Make sure to check the error on Close.
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}
	// End .tar file

	// Start upload .tar
	uploadOption := api.UploadToContainerOptions{
		InputStream:          buf,
		Path:                 "/",
		NoOverwriteDirNonDir: true,
	}

	return d.client.UploadToContainer(name, uploadOption)
}
