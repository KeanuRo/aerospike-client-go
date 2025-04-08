package provider

import (
	"log"
	"os"
	"time"

	dynconfig "github.com/aerospike/aerospike-client-go/v8/config"
	"github.com/stretchr/testify/assert/yaml"
)

const defaultFilePath = "~/aerospikeconfig.yaml"

type YamlConfigProvider struct {
	configFilePath string
	oldModTime     time.Time
}

func NewYamlConfigProvider() dynconfig.ConfigProvider {
	return &YamlConfigProvider{
		configFilePath: defaultFilePath,
	}
}

func NewYamlConfigProviderWithPath(configFilePath string) dynconfig.ConfigProvider {
	return &YamlConfigProvider{
		configFilePath: configFilePath,
	}
}

func (yc *YamlConfigProvider) LoadConfig() *dynconfig.Config {
	// Load the YAML configuration file and parse it into a Config struct
	// This is a placeholder implementation. Actual implementation would involve
	// reading from a YAML file and unmarshalling it into the Config struct.
	info, err := os.Stat(defaultFilePath)
	if err != nil {
		// handle error
	}

	modTime := info.ModTime()
	// Compare to previously stored modTime
	if modTime.After(yc.oldModTime) {
		// file changed
		yc.oldModTime = modTime
		// re-unmarshal your struct

		data, err := os.ReadFile("config.yaml")
		if err != nil {
			log.Fatalf("failed to read config file: %v", err)
		}
		var config dynconfig.Config
		if err := yaml.Unmarshal(data, &config); err != nil {
			log.Fatalf("failed to unmarshal yaml: %v", err)
		}

		return &config
	}

	return nil
}
