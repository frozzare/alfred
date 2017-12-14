package proxy

import (
	"strings"
	"testing"

	"github.com/frozzare/alfred/internal/docker"
	"github.com/frozzare/go-assert"
)

func TestApp(t *testing.T) {
	d, err := docker.NewDocker()
	if err != nil {
		t.Fatal(err)
	}

	proxy := NewProxy(&Options{
		Docker: d,
		Name:   "alfred_proxy",
	})

	assert.Nil(t, proxy.Start())
	assert.True(t, strings.Contains(proxy.Start().Error(), "API error (400)"))
	assert.Nil(t, proxy.Stop())
}
