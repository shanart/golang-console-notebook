package main

import (
	"notebook/app"
	"notebook/app/config"
)

func main() {
	config := config.GetConfig()
	app.Run(config)

}
