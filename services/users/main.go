package main

import (
	"pluto/pkgs/pluto"
	"pluto/pkgs/providers/http"
)

func main() {
	client := pluto.NewClient("test")

	// Create Providers
	httpProvider := http.NewProvider("8080")

	httpProvider.RegisterEndpoint(http.GetMethod, "/health-check", func() {

	})

	// TODO; think about how to add a middleware which could authenticate the user
	// TODO; think about how to add the :userID
	httpProvider.RegisterEndpoint(http.GetMethod, "/user/:userID", func() {

	})

	// Register the Providers
	client.RegisterProvider(httpProvider)

	// Start the service
	client.Start()
}
