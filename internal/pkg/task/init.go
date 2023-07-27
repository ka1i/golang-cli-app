package task

import (
	"runtime"
)

var (
	waitCount int = 0

	waitLatestTaskDone = make(chan bool)

	goroutineMaxProcess = make(chan bool, runtime.NumCPU()*100)
)
