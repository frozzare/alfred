package app

import (
	"testing"

	"github.com/frozzare/alfred/config"
	"github.com/frozzare/go-assert"
)

func TestApp(t *testing.T) {
	path := "../examples/html/alfred.json"
	config, _ := config.ReadConfig(path)
	app := NewApp(&Options{
		Config: config,
	})

	assert.Nil(t, app.Start())
	assert.Nil(t, app.Stop())
}
