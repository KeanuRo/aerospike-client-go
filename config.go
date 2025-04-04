package aerospike

import (
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	dynconfig "github.com/aerospike/aerospike-client-go/v8/config"
	"github.com/aerospike/aerospike-client-go/v8/logger"
)

var configProviderMu sync.RWMutex
var configProviders = make(map[string]dynconfig.ConfigProvider)

const defaultConfigProviderName = "yaml"

type DynConfig struct {
	lock               sync.RWMutex
	config             *dynconfig.Config
	wgConfig           sync.WaitGroup
	configInitialized  *atomic.Bool
	clientPolicy       *ClientPolicy
	configWatchChannel chan struct{}
}

func NewDynConfig(policy *ClientPolicy) *DynConfig {
	return &DynConfig{
		clientPolicy:       policy,
		configWatchChannel: make(chan struct{}),
	}
}

func register(name string, provider dynconfig.ConfigProvider) dynconfig.ConfigProvider {
	configProviderMu.Lock()
	defer configProviderMu.Unlock()

	if _, ok := configProviders[name]; ok {
		panic("provider already registered")
	}

	if provider == nil {
		panic("provider is nil")
	}

	if _, dup := configProviders[name]; dup {
		panic("config provider called twice" + name)
	}

	configProviders[name] = provider

	return configProviders[name]
}

func (dc *DynConfig) loadConfig() error {
	configProviderMu.RLock()
	defer configProviderMu.RUnlock()

	if !dc.configInitialized.Load() {
		// Get entire config and update/initialize the config
		if len(configProviders) == 0 {
			// We should never get here, just in case  we do initialize to default config
			logger.Logger.Warn("Configuration provider has not been registered. Using default config provider.")
			register(defaultConfigProviderName, nil)
		} else {
			for _, provider := range configProviders {
				_, loadedConfig := provider.LoadConfig()
				dc.lock.Lock()
				dc.config = loadedConfig
				dc.lock.Unlock()
				break
			}
		}
	} else {
		// get static and dynamic parts of the config
		if len(configProviders) > 1 {
			logger.Logger.Warn("Multiple config providers registered. Using the first one.")
			if defaultConfigProvider, ok := configProviders[defaultConfigProviderName]; ok {
				_, loadedConfig := defaultConfigProvider.LoadConfig()
				dc.lock.Lock()
				dc.config.Dynamic = loadedConfig.Dynamic
				dc.lock.Unlock()
			}
		} else if len(configProviders) == 1 {
			for _, provider := range configProviders {
				_, loadedConfig := provider.LoadConfig()
				dc.lock.Lock()
				dc.config.Dynamic = loadedConfig.Dynamic
				dc.lock.Unlock()
				break
			}
		}
	}

	return nil
}

func (dc *DynConfig) watchConfig() {
	logger.Logger.Info("Starting the config watch goroutine...")

	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error("Watch config goroutine crashed: %s", debug.Stack())
			go dc.watchConfig()
		}
	}()

	defer dc.wgConfig.Done()

	tendInterval := dc.clientPolicy.TendInterval
	if tendInterval <= 10*time.Millisecond {
		tendInterval = 10 * time.Millisecond
	}

Loop:
	for {
		select {
		case <-dc.configWatchChannel:
			logger.Logger.Debug("Watch config channel closed. Stopping watch goroutine.")
			break Loop
		case <-time.After(tendInterval):
			tm := time.Now()
			if err := dc.loadConfig(); err != nil {
				logger.Logger.Warn(err.Error())
			}

			// Tending took longer than requested tend interval.
			// Tending is too slow for the cluster, and may be falling behind schedule.
			if tendDuration := time.Since(tm); tendDuration > dc.clientPolicy.TendInterval {
				logger.Logger.Warn("Watching took %s, while your requested ClientPolicy.TendInterval is %s. Config fetches are slower than the interval, and may be falling behind the changes.", tendDuration, dc.clientPolicy.TendInterval)
			}
		}
	}
}

func (dc *DynConfig) updateConfig(config *dynconfig.Config) {
	dc.lock.Lock()
	defer dc.lock.Unlock()

	if config == nil {
		logger.Logger.Error("Config is nil")
		return
	}

	dc.config = config
}
