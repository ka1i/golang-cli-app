package task

import (
	"sync"
	"time"

	"github.com/ka1i/cli/internal/pkg/config"
	"github.com/ka1i/cli/pkg/logger"
	"go.uber.org/zap"
)

var New = getTask()

func getTask() *taskbBase {
	return &taskbBase{}
}

type taskbBase struct {
	surveillance time.Duration
}

func (t *taskbBase) init() {
	time.Sleep(time.Second * 1)
	interval := config.Cfg.Get().Interval
	if interval == 0 {
		logger.Error("Please check background task interval, Use defaults interval: 1min")
		interval = time.Second * 60
	}
	t.surveillance = interval / 2

	logger.Info("background task", zap.String("interval", interval.String()), zap.String("surveillance", t.surveillance.String()))
}

func (t *taskbBase) Start() {

	t.init()

	var wg sync.WaitGroup
	surveillanceTicker := time.NewTicker(t.surveillance)

	var waitLatestTask bool = false
	for {
		select {
		case <-surveillanceTicker.C:
			// surveillance task
			if waitLatestTask { // check result wait
				waitCount++
				if waitCount == 1 {
					go func() {
						wg.Wait()
						waitLatestTaskDone <- true
					}()
				}
			} else { // start surveillance group by
				wg.Add(1)
				go surveillance(&wg)
				waitLatestTask = true
			}
		case <-waitLatestTaskDone:
			if waitCount > 1 {
				logger.Warnf("注意，请调整任务间隔，建议调整间隔最小为: %v~%v\n", t.surveillance*time.Duration(waitCount-1), t.surveillance*time.Duration(waitCount+1))
			}
			// analyse surveillance task result
			waitCount = 0
			waitLatestTask = false
		}
	}
}
