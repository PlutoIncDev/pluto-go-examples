package routes

import (
	"pluto/pkgs/providers/http"
	"time"
)

var AuthLoginEndpoint = http.NewEndpoint("POST", "/auth/login", authLoginHandler)

func authLoginHandler(ctx *http.Context) {
	time.Sleep(time.Second)
	ctx.HTTP.JSON(200, http.H{
		"jwt": "<JWT HERE>",
	})
}
