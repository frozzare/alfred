package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

var (
	globalConfig *GlobalConfig
)

// GlobalConfig represents global configuration.
type GlobalConfig struct {
	Docker *Docker `json:"docker"`
	Proxy  *Proxy  `json:"proxy"`
	TLD    string  `json:"tld"`
}

// Default sets the default global config.
func (g *GlobalConfig) Default() {
	if len(g.TLD) == 0 {
		g.TLD = "dev"
	}

	if g.Docker == nil {
		g.Docker = &Docker{}
		g.Docker.Default()
	}

	if g.Proxy == nil {
		g.Proxy = &Proxy{}
		g.Proxy.Default()
	}
}

// ReadGlobalConfig tries to reads the global config.
func ReadGlobalConfig() error {
	g := &GlobalConfig{}
	h, err := homedir.Dir()

	if err != nil {
		return err
	}

	path := filepath.Join(h, ".alfred.json")
	b, err := ioutil.ReadFile(path)

	if !os.IsNotExist(err) {
		if err := json.Unmarshal(b, g); err != nil {
			return errors.Wrap(err, "Parsing json")
		}
	}

	g.Default()

	globalConfig = g

	return nil
}

// Global returns the global config.
func Global() *GlobalConfig {
	if globalConfig == nil {
		globalConfig = &GlobalConfig{}
		globalConfig.Default()
	}

	return globalConfig
}

// Docker represents docker configuration.
type Docker struct {
	Host string `json:"host"`
}

// Default sets the default docker global config.
func (d *Docker) Default() {
}

// Default sets the default proxy global config.
func (p *Proxy) Default() {
}

// Proxy represents type configuration.
type Proxy struct {
	Type string `json:"type"`
}
