package routes

import "github.com/plutoincdev/pluto-go/pkgs/providers/http"

var HealthCheckEndpoint = http.NewEndpoint("GET", "/health-check", healthCheckHandler)

func healthCheckHandler(ctx *http.Context) {
	ctx.HTTP.JSON(200, http.H{
		"status": "healthy",
	})
}
