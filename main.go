package main

import (
	"log"
	"pluto/pkgs/client"
	"pluto/pkgs/providers/http"
)

func main() {
	c := client.New("users")

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

	c.Start(true)
}