package config

import (
	"strings"
	"testing"

	"github.com/frozzare/go-assert"
)

func TestDefaultConfig(t *testing.T) {
	c := &Config{}
	c.Default()

	assert.Equal(t, c.Port, 2015)
	assert.Equal(t, c.Image, "joshix/caddy")
	assert.NotEmpty(t, c.Host)
	assert.NotEmpty(t, c.Path)
	assert.NotEmpty(t, c.Env)
	assert.Empty(t, c.Links)
}

func TestReadConfig(t *testing.T) {
	path := "../examples/html/alfred.json"
	c, err := ReadConfig(path)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(c.Path, "/public"))
}
