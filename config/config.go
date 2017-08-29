package config

import (
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// Config represents application configuration.
type Config struct {
	Env   []string `json:"env"`
	Image string   `json:"image"`
	Host  string   `json:"host"`
	Links []string `json:"links"`
	Port  int      `json:"port"`
	Path  string   `json:"path"`
}

// Default sets the default config.
func (c *Config) Default() error {
	// Set default port.
	if c.Port == 0 {
		c.Port = 80
	}

	// Set host based on folder.
	if c.Host == "" {
		path, err := os.Getwd()
		if err != nil {
			return err
		}

		parts := strings.Split(path, "/")
		c.Host = fmt.Sprintf("%s.%s", parts[len(parts)-1], Global().TLD)
	}

	// Set default image, works for HTML.
	if c.Image == "" {
		c.Image = "joshix/caddy"
		c.Port = 2015
	}

	// Add missing parts to path.
	c.Path = path(c.Path)

	// Count number of ":" char in path.
	r := regexp.MustCompile("\\:")
	index := suffixarray.New([]byte(c.Path))
	result := index.FindAllIndex(r, -1)

	// Add missing parts to application path.
	if len(result) == 0 {
		c.Path = c.Path + ":/var/www/html:rw"
	} else if len(result) == 1 {
		c.Path = c.Path + ":rw"
	}

	// Set virtual host and virtual port default values.
	c.Env = append(c.Env, "VIRTUAL_HOST="+c.Host)
	c.Env = append(c.Env, fmt.Sprintf("VIRTUAL_PORT=%d", c.Port))

	// Set empty slice as default value for links.
	if len(c.Links) == 0 {
		c.Links = []string{}
	}

	return nil
}

// ReadConfig tries to reads the given path or create a default config.
func ReadConfig(path string) (*Config, error) {
	c := &Config{}
	b, err := ioutil.ReadFile(path)

	if !os.IsNotExist(err) {
		if err := json.Unmarshal(b, c); err != nil {
			return nil, errors.Wrap(err, "Parsing json")
		}
	}

	if err := c.Default(); err != nil {
		return nil, errors.Wrap(err, "Default config")
	}

	return c, nil
}
