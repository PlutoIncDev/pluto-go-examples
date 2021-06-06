package main

import (
	"pluto/pkgs/logging"
	"pluto/pkgs/pluto"
)

func handleRpcFindUserByID() {

}

func handleHttpUserByID() {

}

func handleHttpHealthCheck() {

}

func main() {

	// setup client
	client := pluto.NewClient()

	// setup client options
	client.Options.SetName("users")

	// HTTP
	client.Events.Http.Get("/health-check", handleHttpHealthCheck)
	client.Events.Http.Get("/user/:id", handleHttpUserByID)

	// RPC
	client.Events.Rpc.Register("findUserByID", handleRpcFindUserByID)

	// Start the service
	logging.Fatal(client.Start())
}
