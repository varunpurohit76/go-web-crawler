package base

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type RequestContext struct {
	traceId string
	logger  *log.Entry
}

func (ctx *RequestContext) TraceId() string {
	return ctx.traceId
}

func (ctx *RequestContext) SetTraceId() {
	if ctx.traceId != "" {
		return
	}
	ctx.traceId = uuid.New().String()
}

func (ctx *RequestContext) Logger() *log.Entry {
	return ctx.logger
}

func (ctx *RequestContext) SetLogger() {
	ctx.logger = log.WithField("traceID", ctx.traceId)
}

func
NewRequestContext() *RequestContext {
	rc := &RequestContext{}
	rc.SetTraceId()
	rc.SetLogger()
	return rc
}