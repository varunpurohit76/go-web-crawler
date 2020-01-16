package base

import (
	log "github.com/sirupsen/logrus"
)

func InitLog() {
	log.SetFormatter(&log.JSONFormatter{})
	level, _ := log.ParseLevel(Config.Log.Level)
	log.SetLevel(level)
}
