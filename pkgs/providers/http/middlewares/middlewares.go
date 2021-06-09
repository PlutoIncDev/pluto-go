package middlewares

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"pluto/pkgs/providers/http"
	"time"
)

func Logging() http.Handler {
	return func(ctx *http.Context) {

		startedAt := time.Now()
		ctx.HTTP.Next()
		requestDuration := time.Now().Sub(startedAt)

		// todo: check that this fires even on error
		log.WithFields(log.Fields{
			"method":      ctx.HTTP.Request.Method,
			"url":         ctx.HTTP.Request.URL,
			"status_code": ctx.HTTP.Writer.Status(),
		}).Info(fmt.Sprintf("[%s] %d (%s)	 %s		%s", ctx.HTTP.Request.Method, ctx.HTTP.Writer.Status(), ctx.HTTP.Request.URL, ctx.HTTP.Request.Host, requestDuration))
	}
}

func Limiter(n int) http.Handler {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }

	return func(ctx *http.Context) {
		acquire()       // before request
		defer release() // after request

		ctx.HTTP.Next()
	}
}
