package base

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func LogLatency(key string, f log.Fields, startTime time.Time) {
	log.WithFields(f).WithFields(log.Fields{"value": time.Now().Sub(startTime).Seconds()}).Info(key)
}
