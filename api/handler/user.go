package handler

import (
	"c1pherten/yet-webapp2/api"
	"c1pherten/yet-webapp2/appctx"
	"c1pherten/yet-webapp2/dto"
	"c1pherten/yet-webapp2/middleware"
	// "c1pherten/yet-webapp2/repository"
	"c1pherten/yet-webapp2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	loginTypeQr = "qrcode"
)

type UserHandler struct {
	c appctx.Container
	s *service.UserService
	msgService *service.MsgService
	
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	id, b := ctx.Params.Get("id")
	if !b {
		// ctx.JSON(http.StatusBadRequest, api.EmptyResponse())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := h.s.GetUserByID(id)
	if err != nil {
		// TODO: log
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, api.Success(&user))
}

func (h *UserHandler) CreateNewUser(ctx *gin.Context) {
	var u dto.CreateUser
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.DebugErr(err))
		return
	}
	h.c.Logger().Info(u)
	// if err := u.Validate(); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, api.DebugErr(err))
	// 	return 
	// }
	
	rUser, err := h.s.CreateNewUser(u)
	if err != nil {
		// log.Info(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.DebugErr(err))
		return
	}
	tokenString, err := middleware.Sign(&rUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Internal())
		return
	}

	ctx.JSON(http.StatusOK, api.Success(tokenString))
	
	return
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var login dto.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		// ctx.AbortWithStatusJSON(http.StatusBadRequest, api.DebugErr(err))
		api.BadReq(ctx, err)
		return
	}

	// ctx.JSON(http.StatusOK, api.Success(login))
	token, err := h.s.Login(login)
	if err != nil {
		api.InternalError(ctx, err)
		return 
	}
	
	api.OK(ctx, token)
	return
}

func (h *UserHandler) QrcodeID(ctx *gin.Context) {
	id, err := h.s.QrcodeID()
	if err != nil {
		api.InternalError(ctx, err)
		return
	}

	api.OK(ctx, id)
	return
}

func (h *UserHandler) QrcodeLogin(ctx *gin.Context) {
	
}


// func NewUserHandler(s *service.UserService) *UserHandler {
func NewUserHandler(c appctx.Container) *UserHandler {
	s := service.NewUserService(c)
	msgService := service.NewMsgService(c)
	
	return &UserHandler{
		c:          c,
		s:          s,
		msgService: msgService,
	}
}