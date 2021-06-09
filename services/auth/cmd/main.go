package main

import (
	"auth/pkgs/routes"
	"github.com/plutoincdev/pluto-go/pkgs/client"
	log "github.com/sirupsen/logrus"
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
