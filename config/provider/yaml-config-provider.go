package provider

import (
	"fmt"
	"os"
	"time"

	dynconfig "github.com/aerospike/aerospike-client-go/v8/config"
	"github.com/aerospike/aerospike-client-go/v8/logger"
	"gopkg.in/yaml.v3"
)

const defaultFilePath = "aerospikeconfig.yaml"

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
	if info == nil {
		logger.Logger.Debug("File does not exist %s . Nothing to do...", defaultFilePath)
		return nil
	}

	modTime := info.ModTime()
	// Compare to previously stored modTime
	if modTime.After(yc.oldModTime) {
		yc.oldModTime = modTime

		data, err := os.ReadFile(defaultFilePath)
		if err != nil {
			logger.Logger.Error("Failed to read file %s. Error: %v", defaultFilePath, err)
		}
		var config dynconfig.Config
		if err := yaml.Unmarshal(data, &config); err != nil {
			fmt.Printf("Failed to serialize file %s to object. Error: %v", defaultFilePath, err)
			logger.Logger.Error("Failed to serialize file %s to object. Error: %v", defaultFilePath, err)
		}

		return &config
	}

	return nil
}

type Duration time.Duration

func (d *Duration) UnmarshalYAML(b []byte) error {
	var value int64
	if err := yaml.Unmarshal(b, &value); err != nil {
		return err
	}
	*d = Duration(time.Duration(value))
	return nil
}
