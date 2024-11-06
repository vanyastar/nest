package logService

import (
	"strconv"
	"sync"
	"time"
)

var duration = time.Now()
var mu sync.Mutex

func Log(service, text string) {
	execTime := float64(time.Now().Local().UnixMicro()-duration.UnixMicro()) / 1000
	println(
		colorCodes[green] + "[NestGo] - " +
			colorCodes[white] + time.Now().Format("2006-01-02 15:04:05") +
			colorCodes[yellow] + " [" + service + "] " +
			colorCodes[green] + text +
			colorCodes[yellow] + " +" + strconv.FormatFloat(execTime, 'f', 1, 64) + "ms",
	)
	mu.Lock()
	duration = time.Now()
	mu.Unlock()
}
