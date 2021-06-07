package main

import (
	"pluto/pkgs/pluto"
	"pluto/pkgs/providers/http"
)

func main() {
	client := pluto.NewClient("test")

	httpProvider := http.NewProvider("8080")

	httpProvider.RegisterEndpoint("GET", "/health-check", func() {

	})

	// Register the providers
	client.RegisterProvider(httpProvider)

	// Start the service
	client.Start()
}
