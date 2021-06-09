package routes

import (
	"pluto/pkgs/providers/http"
	"pluto/pkgs/providers/http/middlewares"
)

func SetupRoute() *http.Provider {
	httpProvider := http.NewProvider("8080")

	httpProvider.RegisterMiddleware(middlewares.Logging())

	httpProvider.RegisterEndpoint(HealthCheckEndpoint)
	httpProvider.RegisterEndpoint(AuthLoginEndpoint)
	return httpProvider
}
