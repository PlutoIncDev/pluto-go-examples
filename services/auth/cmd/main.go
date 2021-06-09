package main

import (
	"auth/pkgs/routes"
	log "github.com/sirupsen/logrus"
	"pluto/pkgs/client"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	c := client.New("auth")
	httpProvider := routes.SetupRoute()
	c.RegisterProvider(httpProvider)
	c.Start(true)
}
