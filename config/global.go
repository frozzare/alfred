package config

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

// SetGlobal sets the global config.
func SetGlobal(g *GlobalConfig) {
	if g == nil {
		return
	}

	globalConfig = g
	globalConfig.Default()
}

// Global returns the global config.
func Global() *GlobalConfig {
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
