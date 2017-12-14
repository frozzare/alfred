package app

import (
	"strings"
	"testing"

	"github.com/frozzare/alfred/config"
	"github.com/frozzare/alfred/internal/docker"
	"github.com/frozzare/go-assert"
)

func TestApp(t *testing.T) {
	path := "../examples/html/alfred.json"
	config, err := config.ReadConfig(path)
	if err != nil {
		t.Fatal(err)
	}

	d, err := docker.NewDocker()
	if err != nil {
		t.Fatal(err)
	}

	app := NewApp(&Options{
		Config: config,
		Docker: d,
	})

	assert.Nil(t, app.Start())
	assert.True(t, strings.Contains(app.Start().Error(), "API error (400)"))
	assert.Nil(t, app.Stop())
}
