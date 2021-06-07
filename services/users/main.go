package main

import (
	"log"
	"pluto/pkgs/pluto"
	"pluto/pkgs/providers/http"
)

func handleHealthCheck(ctx *http.Context) {
	log.Println("health check")

	ctx.HTTP.JSON(200, http.H{"success": "true"})
}

func handleGetUserById(ctx *http.Context) {
	log.Println("get user by id")

	ctx.HTTP.JSON(200, http.H{"user": http.H{"id": "123", "email": "calumpeterwebb@icloud.com"}})
}

func CreateClient() *pluto.Client {
	// TODO: add config file / environment file?
	client := pluto.NewClient("test")

	// Create HTTP Provider
	httpProvider := http.NewProvider("8080")

	httpProvider.RegisterEndpoint(http.GetMethod, "/health-check", handleHealthCheck)
	httpProvider.RegisterEndpoint(http.GetMethod, "/user/:id", handleGetUserById)

	// Register the Providers
	client.RegisterProvider(httpProvider)

	return client
}

func RunClient() {
	// Build the service
	client := CreateClient()

	// Start the service
	client.Start()
}

func main() {
	RunClient()
}
