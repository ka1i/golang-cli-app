package timezone

import (
	"time"

	"github.com/ka1i/cli/internal/pkg/config"
)

var timeFormatStr string = "2006-01-02T15:04:05.000Z0700"

func localTimeZone() *time.Location {
	tz := config.Cfg.Get().TimeZone

	local, err := time.LoadLocation(tz)
	if err != nil {
		local = time.FixedZone(tz, int((time.Hour * 8).Seconds()))
	}

	return local
}

func Init() {
	time.Local = localTimeZone()
}

func TZ() string {
	return localTimeZone().String()
}

func Format(t time.Time) string {
	return t.In(localTimeZone()).Format(timeFormatStr)
}

func Parse(timeStr string) time.Time {
	local := localTimeZone()
	sTime, err := time.ParseInLocation(timeFormatStr, timeStr, local)
	if err != nil {
		sTime = time.Now().In(local)
	}
	return sTime
}
