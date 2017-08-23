package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Config struct for alfred applications.
type Config struct {
	Env   []string `json:"env"`
	Image string   `json:"image"`
	Host  string   `json:"host"`
	Port  int      `json:"port"`
	Path  string   `json:"path"`
}

// Default sets the default config.
func (c *Config) Default() error {
	if c.Port == 0 {
		c.Port = 80
	}

	if c.Host == "" {
		path, err := os.Getwd()
		if err != nil {
			return err
		}

		parts := strings.Split(path, "/")
		c.Host = fmt.Sprintf("%s.dev", parts[len(parts)-1])
	}

	c.Path = path(c.Path)

	if c.Image == "" {
		c.Image = "joshix/caddy"
		c.Port = 2015
	}

	if !strings.Contains(c.Path, ":") {
		c.Path = c.Path + ":/var/www/html:ro"
	}

	c.Env = append(c.Env, "VIRTUAL_HOST="+c.Host)
	c.Env = append(c.Env, fmt.Sprintf("VIRTUAL_PORT=%d", c.Port))

	return nil
}

func path(p string) string {
	if strings.HasPrefix(p, "/") {
		return p
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(p, "./") {
		p = strings.Replace(p, "./", "", -1)
	}

	return filepath.Join(path, p)
}

// ReadConfig tries to reads the given path or create a default config.
func ReadConfig(path string) (*Config, error) {
	c := &Config{}
	b, err := ioutil.ReadFile(path)

	if !os.IsNotExist(err) {
		if err := json.Unmarshal(b, c); err != nil {
			return nil, errors.Wrap(err, "parsing json")
		}
	}

	if err := c.Default(); err != nil {
		return nil, errors.Wrap(err, "default config error")
	}

	return c, nil
}
