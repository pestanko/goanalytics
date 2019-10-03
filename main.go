package main

import (
	"fmt"
	"os"

	"github.com/pestanko/goanalytics/goanalytics"
	"gopkg.in/yaml.v2"
)

var log = goanalytics.CreateLogger()

func main() {
	config := loadConfig()
	printConfig(config)
}

func loadConfig() *goanalytics.Config {

	config := goanalytics.CreateConfig()
	configFile := os.Getenv("GO_ANALYTICS_CONFIG_FILE")
	if configFile == "" {
		configFile = "resources/config.yaml"
	}
	config.ReadYaml(configFile)
	config.ReadEnv()

	return config
}

func printConfig(config *goanalytics.Config) {
	str, err := yaml.Marshal(config)
	if err != nil {
		log.Error("Unnable to marshall")
		return
	}

	fmt.Printf("%s\n", str)
}
