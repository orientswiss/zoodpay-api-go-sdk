package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config represents an application configuration.
type Config struct {
	// Data Source Endpoint/Name
	Host string `yaml:"api_url"`
	// Api Version
	Version string `yaml:"version"`
	// Default Error Code
	DefaultCode string `yaml:"default_code"`
	// Credential Source Type
	CST string `yaml:"cst"`
	// Yml Configurations
	YML map[string]string `yaml:"yml,omitempty"`
	// Database Configurations
	DB map[string]string `yaml:"db,omitempty"`
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string) (*Config, error) {
	// default config
	c := Config{}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	return &c, err
}
