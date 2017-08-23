package docker

import (
	api "github.com/fsouza/go-dockerclient"
)

// pullImage will pull a image if it don't exists.
func (d *Docker) pullImage(image string) error {
	if dockerImage, _ := d.client.InspectImage(image); dockerImage == nil {
		repository, tag := api.ParseRepositoryTag(image)
		if err := d.client.PullImage(api.PullImageOptions{
			Repository: repository,
			Tag:        tag,
		}, api.AuthConfiguration{}); err != nil {
			return err
		}
	}

	return nil
}
