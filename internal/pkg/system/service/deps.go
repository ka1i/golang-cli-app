package service

import (
	"sync"
)

type dependencies struct {
	mutex sync.RWMutex
	item  map[string]string
}

func (deps *dependencies) Get() map[string]string {
	deps.mutex.TryRLock()
	defer deps.mutex.RUnlock()
	return deps.item
}

func (deps *dependencies) Set(key, value string) {
	deps.mutex.Lock()
	defer deps.mutex.Unlock()
	deps.item[key] = value
}

func getDependencies() *dependencies {
	return &dependencies{
		item: make(map[string]string),
	}
}

var Dependencies *dependencies = getDependencies()
