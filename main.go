package main

import (
	"c1pherten/yet-webapp2/app"
	"c1pherten/yet-webapp2/config"
	"c1pherten/yet-webapp2/container"
	"c1pherten/yet-webapp2/log"
	"embed"
)

//go:embed resources/config/application.*.yml
var appYmlFile embed.FS

//go:embed resources/config/messages.properties
var properties embed.FS

func main() {
	cfg, env := config.LoadAppConfig(appYmlFile)
	messages := config.LoadMessagesConfig(properties)
	l := log.NewLogger()
	l.Info(cfg)
	l.Info(env)
	l.Info(messages)

	container := container.NewContainer(l, cfg, env)

	app := app.NewApp(container)
	app.Run("0.0.0.0:8080")
}
