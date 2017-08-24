package app

import (
	"testing"

	"github.com/frozzare/alfred/config"
	"github.com/frozzare/alfred/internal/docker"
	"github.com/frozzare/go-assert"
)

func TestApp(t *testing.T) {
	path := "../examples/html/alfred.json"
	config, _ := config.ReadConfig(path)
	d, err := docker.NewDocker()
	if err != nil {
		t.Fatal(err)
	}
	app := NewApp(&Options{
		Config: config,
		Docker: d,
	})

	assert.Nil(t, app.Start())
	assert.Nil(t, app.Stop())
}
