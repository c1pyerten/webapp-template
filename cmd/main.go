package main

import (
	"c1pherten/yet-webapp2/app"
	"c1pherten/yet-webapp2/config"
	"c1pherten/yet-webapp2/container"
	"c1pherten/yet-webapp2/log"
	"embed"
)

// go:embed resources/config/application.*.yml
var appYmlFile embed.FS

func main() {
	config, env := config.LoadAppConfig(appYmlFile)
	l := log.NewLogger()
	container := container.NewContainer(l, config, env)

	app := app.NewApp(container)
	app.Run(":8080")

}
