package config

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Modules []*Module `yaml:"modules"`
}

type Module struct {
	Package  string  `yaml:"package"`
	Type     string  `yaml:"type"`
	Target   string  `yaml:"target"`
	Sources  sources `yaml:"sources,omitempty"`
	Redirect string  `yaml:"redirect,omitempty"`
}

// create custom type so we can override String()
type sources []string

func (list sources) String() string {
	return strings.Join(list, " ")
}

// LoadConfig loads the config file
func (c *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}
