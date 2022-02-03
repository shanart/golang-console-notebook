package main

import (
	"notebook/app"
	"notebook/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run()
}
