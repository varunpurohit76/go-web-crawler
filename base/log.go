package base

import (
	log "github.com/sirupsen/logrus"
)

func InitLog(level log.Level) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(level)
}
