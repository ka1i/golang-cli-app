package task

import (
	"sync"

	"github.com/ka1i/cli/pkg/logger"
)

func surveillance(wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	goroutineMaxProcess <- true
	go func() {
		defer wg.Done()
		defer func() { <-goroutineMaxProcess }()
		logger.Info("new task running...")
	}()
}
