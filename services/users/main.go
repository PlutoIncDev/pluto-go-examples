package main

import (
	"pluto/pkgs/entrypoints/http"
	"pluto/pkgs/pluto"
)


func main() {
	// setup client
	client, err := pluto.NewClient("users")
	if err != nil {
		panic(err)
	}

	httpEntrypoint := http.New()
	client.RegisterEntrypoint(httpEntrypoint)

	httpEntrypoint.HTTP("GET", "/health-check", func() {

	})

	httpEntrypoint.HTTP("GET", "/health-check", func() {

	})


	//
	//// HTTP
	//client.Events.Http.Get("/health-check", handleHttpHealthCheck)
	//client.Events.Http.Get("/user/:id", handleHttpUserByID)
	//
	//// RPC
	//client.Events.Rpc.Register("findUserByID", handleRpcFindUserByID)

	// Start the service
	client.Start()
}
