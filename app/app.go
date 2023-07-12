package app

import (
	"c1pherten/yet-webapp2/appctx"
	"fmt"
)

// App app instance
type App struct {
	server *server
	c      appctx.Container
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
func NewApp(c appctx.Container) *App {
	server := newServer(c)
	return &App{
		server: server,
		c:      c,
	}
}
