package main

import (
	"os"

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
	Host       string `yaml:"Host"`
	Port       []int  `yaml:"Port"`
	SecurePort []int  `yaml:"SecurePort"`
	Timeout    int    `yaml:"Timeout"`
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
