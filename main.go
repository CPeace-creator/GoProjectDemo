package main

import (
	"goDemo/config"
	"goDemo/router"
)

func main() {
	config.InitConfig()
	r := router.SetRouter()
	port := config.AppConfig.App.Port
	if port == "" {
		port = "8000"
	}
	r.Run(port)
}
