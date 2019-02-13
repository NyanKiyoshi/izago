package managebot

import "time"

var startTime = time.Now()

func getUptime() time.Duration {
	return time.Since(startTime)
}
