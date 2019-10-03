package main

import (
	"fmt"
	"os"

	"github.com/pestanko/goanalytics/goanalytics"
	"gopkg.in/yaml.v2"
)

var log = goanalytics.CreateLogger()

func main() {
	config := goanalytics.CreateConfig()
	configFile := os.Getenv("GO_ANALYTICS_CONFIG_FILE")
	config.ReadYaml(configFile)
	config.ReadEnv()

	str, err := yaml.Marshal(config)
	if err != nil {
		log.Error("Unnable to marshall")
		return
	}

	fmt.Printf("%s\n", str)
}
