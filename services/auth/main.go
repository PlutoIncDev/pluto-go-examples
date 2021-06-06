package main

import (
	"pluto/pkgs/logging"
	"pluto/pkgs/pluto"
)

func handleHttpAuthLogin() {

}

func handleHttpHealthCheck() {

}

func main() {

	// setup client
	client := pluto.NewClient()

	// setup client options
	client.Options.SetName("auth")
	client.Options.SetHTTPPort("8090")

	// HTTP
	client.Events.Http.Get("/health-check", handleHttpAuthLogin)
	client.Events.Http.Post("/auth/login", handleHttpHealthCheck)

	// Start the service
	logging.Fatal(client.Start())
}
