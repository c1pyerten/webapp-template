package app

import (
	"c1pherten/yet-webapp2/api"
	"c1pherten/yet-webapp2/api/handler"
	"c1pherten/yet-webapp2/appctx"
	"c1pherten/yet-webapp2/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userRouter *handler.UserHandler
	msgRouter *handler.MessageHandler
	wsRouter *handler.WsHandler
}

func (r *Router) Routes(e *gin.Engine) {
	v1 := e.Group("/v1")
	userGroup := v1.Group("/user")
	{
		userGroup.POST("/login", r.userRouter.Login)
		userGroup.GET("/login/qrcode", r.userRouter.QrcodeID)
		userGroup.POST("", r.userRouter.CreateNewUser)

		// login required handlers
		userGroup.GET("/:id", middleware.Auth, r.userRouter.GetUserByID)
	}

	msgGroup := v1.Group("/messages")
	{
		msgGroup.GET("/target/:id", r.msgRouter.GetMsgById)
	}

	// ws
	v1.GET("/ws", r.wsRouter.HandleWs)

	e.NoRoute(noRoute)
}

// newRouter init router instance
func newRouter(c appctx.Container) *Router {
	userRouter := handler.NewUserHandler(c)
	msgRouter := handler.NewMessageHandler(c)
	wsRouter := handler.NewWsHandler(c)

	return &Router{
		userRouter: userRouter,
		msgRouter:  msgRouter,
		wsRouter:   wsRouter,
	}
}

func noRoute(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, api.NotFound())
}