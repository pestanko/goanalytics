package goanalytics

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// REF: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64

var log = CreateLogger()

// ConfigServer - server configuration
type configServer struct {
	Port string `envconfig:"SERVER_PORT" yaml:"port"`
	Host string `envconfig:"SERVER_HOST" yaml:"host"`
}

// ConfigRedis Redis Configuration
type configRedis struct {
	Username string `envconfig:"REDIS_USERNAME" yaml:"username"`
	Password string `envconfig:"REDIS_PASSWORD" yaml:"password"`
	Host     string `envconfig:"REDIS_HOST" yaml:"host"`
	Database int    `envconfig:"REDIS_DATABASE" yaml:"database"`
}

// Config - Application configration
type Config struct {
	Server configServer `yaml:"server"`
	Redis  configRedis  `yaml:"redis"`
}

//ReadYaml - Reads configuration from the YAML file
func (cfg *Config) ReadYaml(filePath string) {
	if !Exists(filePath) {
		log.Debugf("[CONFIG] Yaml file not existst: %s", filePath)
		return
	}
	f, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Unable to read a YAML file: %s", filePath)
		return
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Errorf("Unable to decode a file: %s", filePath)
		return
	}
}

// ReadEnv - Reads configuration from the Environment
func (cfg *Config) ReadEnv() {
	err := envconfig.Process("go_analytics", cfg)
	if err != nil {
		log.Error("Unable to read envconfig")
		return
	}
}

// CreateConfig - Creates config
func CreateConfig() *Config {
	return &Config{
		Server: configServer{
			Port: "5000",
			Host: "localhost",
		},
		Redis: configRedis{
			Username: "",
			Password: "",
			Database: 0,
			Host:     "localhost:6987",
		},
	}
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
