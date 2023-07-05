package app

import (
	"c1pherten/yet-webapp2/container"
	"fmt"
)

// App app instance
type App struct {
	server *server
	c      container.Container
}

// Run app running interface
func (a *App) Run(addr string) {
	if addr == "" {
		addr = "0.0.0.0:8080"
	}

	if err := a.server.Run(addr); err != nil {
		fmt.Println(err)
		return
	}

	return
}

// NewApp init app instance
func NewApp(c container.Container) *App {
	server := newServer(c)
	return &App{
		server: server,
		c:      c,
	}
}
