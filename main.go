package main

import (
	"log"
	"pluto/pkgs/client"
	"pluto/pkgs/providers/http"
	"pluto/pkgs/providers/rpc"
	"time"
)

func setupHTTP(c *client.Client) {
	httpProvider := http.NewProvider("8080")

	httpProvider.RegisterEndpoint("GET", "/health-check", func(ctx *http.Context) {
		log.Println("GET /health-check")
		ctx.HTTP.JSON(200, http.H{"success": "true"})
	})

	httpProvider.RegisterEndpoint("GET", "/users", func(ctx *http.Context) {
		log.Println("GET /users")

		ctx.HTTP.JSON(200, http.H{})
	})

	c.RegisterProvider(httpProvider)
}

func setupRPC(c *client.Client) {
	rpcProvider := rpc.NewProvider()

	// todo: add handlers?
	c.RegisterProvider(rpcProvider)
}

func main() {

	var c *client.Client
	c = client.New("users")

	setupHTTP(c)
	setupRPC(c)

	c.Start(false)

	time.Sleep(10 * time.Second)
	c.Stop()
}
