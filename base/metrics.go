package base

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func LogLatency(key string, f log.Fields, startTime time.Time)  {
	log.WithFields(f).WithFields(log.Fields{"value": time.Now().Sub(startTime).Seconds()}).Info(key)
}