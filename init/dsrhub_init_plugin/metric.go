package main

import (
	"fmt"

	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func (p *DSRHubInitPlugin) setupMetrics() error {
	if !p.StatsdEnabled || !p.StatsdAPMEnabled {
		return nil
	}

	p.service.Server.WithCustomMiddlewares(
		gintrace.Middleware(fmt.Sprintf("%s-http-server", p.StatsdAPMServiceName)))

	tracer.Start(
		tracer.WithAgentAddr(fmt.Sprintf("%s:%s", p.StatsdHost, p.StatsdAPMPort)),
		tracer.WithService(p.StatsdAPMServiceName),
	)

	return nil
}
