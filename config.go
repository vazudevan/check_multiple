package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config
type Config struct {
	Checks []Checks `yaml:"Checks"`
}

// Checks
type Checks struct {
	Protocol   string     `yaml:"Protocol"`
	Parameters Parameters `yaml:"Parameters"`
}

// Parameters
type Parameters struct {
	Host       string  `yaml:"Host"`
	Port       []int   `yaml:"Port"`
	SecurePort []int   `yaml:"SecurePort"`
	Timeout    float64 `yaml:"Timeout"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Alternate method using Unmarshal function
	/* yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	} */

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)
	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateAndPrepare(c *Config) ([]checkTcp, error) {
	r := []checkTcp{}
	for _, c := range c.Checks {
		if strings.ToLower(c.Protocol) == "tcp" {
			chk := checkTcp{}
			for _, p := range c.Parameters.Port {
				chk.network = fmt.Sprintf("%s:%d", c.Parameters.Host, p)
				if c.Parameters.Timeout > 0 {
					chk.timeout = c.Parameters.Timeout
				}
				if chk.network != "" {
					r = append(r, chk)
				}
			}
		}
	}
	return r, nil
}

func (c *Config) Validate() error {
	for _, c := range c.Checks {
		switch c.Protocol {
		case "tcp", "TCP":
			if c.Parameters.Host == "" {
				return errors.New("missing Host parameter in for TCP check")
			}
			if len(c.Parameters.Port) == 0 && len(c.Parameters.SecurePort) == 0 {
				return errors.New("missing Port or SecurePort details for TCP check")
			}
		default:
			return errors.New("missing one or more configurations 'Checks'")
			//return nil
		}
	}
	return nil
}
