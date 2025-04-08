package aerospike

import (
	"fmt"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	dynconfig "github.com/aerospike/aerospike-client-go/v8/config"
	"github.com/aerospike/aerospike-client-go/v8/logger"
)

var configProviderMu sync.RWMutex
var configProvider *dynconfig.ConfigProvider

type DynConfig struct {
	lock               sync.RWMutex
	config             *dynconfig.Config
	wgConfig           sync.WaitGroup
	configInitialized  *atomic.Bool
	clientPolicy       *ClientPolicy
	configWatchChannel chan struct{}
}

func NewDynConfig(policy *ClientPolicy) *DynConfig {
	dynConfig := &DynConfig{
		clientPolicy:       policy,
		configWatchChannel: make(chan struct{}),
		configInitialized:  &atomic.Bool{},
	}
	dynConfig.wgConfig.Add(1)

	return dynConfig
}

func register(provider *dynconfig.ConfigProvider) {
	configProviderMu.Lock()
	defer configProviderMu.Unlock()

	if provider == nil {
		panic("provider is nil")
	}

	configProvider = provider
}

func (dc *DynConfig) loadConfig() error {
	configProviderMu.RLock()
	defer configProviderMu.RUnlock()

	if !dc.configInitialized.Load() && configProvider != nil {
		logger.Logger.Debug("Initializing configuration...")
		dc.initConfig()
		dc.configInitialized.Store(true)
	} else {
		dc.providerLoadConfig()
	}

	return nil
}

func (dc *DynConfig) providerLoadConfig() {
	loadedConfig := (*configProvider).LoadConfig()
	dc.lock.Lock()
	if loadedConfig != nil {
		dc.config.Dynamic = loadedConfig.Dynamic
	}
	dc.lock.Unlock()
}

func (dc *DynConfig) initConfig() {
	loadedConfig := (*configProvider).LoadConfig()
	dc.lock.Lock()
	if loadedConfig != nil {
		dc.config = loadedConfig
	}
	dc.lock.Unlock()
}

func (dc *DynConfig) watchConfig() {
	logger.Logger.Info("Starting the config watch goroutine...")

	defer func() {
		// TODO: Add exponential backoff here to resource starvation
		if r := recover(); r != nil {
			logger.Logger.Error("Watch config goroutine crashed: %s", debug.Stack())
			fmt.Printf("Watch config goroutine crashed: %s\n", debug.Stack())
			go dc.watchConfig()
		}
	}()

	defer dc.wgConfig.Done()

	configInterval := dc.clientPolicy.ConfigInterval
	if configInterval <= 10*time.Millisecond {
		configInterval = 10 * time.Millisecond
	}
Loop:
	for {
		if !dc.configInitialized.Load() {
			logger.Logger.Debug("Initializing configuration...")
			tm := time.Now()
			if err := dc.loadConfig(); err != nil {
				logger.Logger.Warn(err.Error())
			}
			if configDuration := time.Since(tm); configDuration > dc.clientPolicy.ConfigInterval {
				logger.Logger.Warn("Reload took %s, but your requested ConfigInterval is %s. "+
					"Reload is slower than the interval and may fall behind changes.",
					configDuration, dc.clientPolicy.ConfigInterval)
			}
		}

		select {
		case <-dc.configWatchChannel:
			logger.Logger.Debug("Watch config channel closed. Stopping watch goroutine.")
			break Loop
		case <-time.After(configInterval):
			tm := time.Now()
			if err := dc.loadConfig(); err != nil {
				logger.Logger.Warn(err.Error())
			}

			if configDuration := time.Since(tm); configDuration > dc.clientPolicy.ConfigInterval {
				logger.Logger.Warn("Watching took %s, while your requested ClientPolicy.TendInterval is %s. Config fetches are slower than the interval, and may be falling behind the changes.", configDuration, dc.clientPolicy.TendInterval)
			}
		}
	}
}
