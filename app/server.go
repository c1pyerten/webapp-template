package app

import (
	"c1pherten/yet-webapp2/container"
	"c1pherten/yet-webapp2/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func newServer(c container.Container) *gin.Engine {
	e := gin.Default()

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("trim", middleware.TrimTag)
	// }
	binding.Validator = &middleware.CustomValidator{}

	// init router
	router := newRouter(c)
	router.Routes(e)

	return e
}