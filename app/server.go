package app

import (
	"c1pherten/yet-webapp2/appctx"
	"c1pherten/yet-webapp2/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/websocket"
)

const (
	A = iota
	B
	C
)

type server struct {
	engine *gin.Engine
	ws     *websocket.Upgrader
}

func newServer(c appctx.Container) *server {
	ws := newWs()
	e := gin.Default()

	e.Use(cors.Default())

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("trim", middleware.TrimTag)
	// }
	binding.Validator = &middleware.CustomValidator{}

	// init router
	router := newRouter(c)
	router.Routes(e)

	return &server{
		engine: e,
		ws:     ws,
	}
}

func (s *server) Run(addr ...string) error {
	return s.engine.Run(addr...)
}
