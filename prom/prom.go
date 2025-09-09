package prom

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Global GlobalConfig `yaml:"global"`
}

type GlobalConfig struct {
	ScrapeInterval     time.Duration     `yaml:"scrape_interval"`
	EvaluationInterval time.Duration     `yaml:"evaluation_interval"`
	ScrapeTimeout      time.Duration     `yaml:"scrape_timeout"`
	ExternalLabels     map[string]string `yaml:"external_labels"`
}

func ConfigFromYAML(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	c := Config{
		GlobalConfig{
			ScrapeTimeout: 10 * time.Second, // set default
		},
	}
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
