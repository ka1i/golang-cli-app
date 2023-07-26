package config

import (
	"sync"
)

type cfg struct {
	mutex   sync.RWMutex
	options Options
}

func (cfg *cfg) Get() *Options {
	cfg.mutex.TryRLock()
	defer cfg.mutex.RUnlock()
	return &cfg.options
}

func (cfg *cfg) Set(newOptions *Options) {
	cfg.mutex.TryLock()
	cfg.options = *newOptions
	cfg.mutex.Unlock()
}

func getCfg() *cfg {
	return &cfg{
		options: Options{},
	}
}

var Cfg *cfg = getCfg()
