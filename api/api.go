package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MsgInternalErr  = "internal error"
	MsgTokenInvalid = "invalid token"
	MsgUnauthorized = "unauthorized"
	

	CodeInternalErr = 10001
	CodeNotFound = 10002
	CodeDebug    = 99
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) Response {
	return Response{
		Code: 0,
		Msg:  "",
		Data: data,
	}
}

func Fail(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func FailWithErr(code int, err error) Response {
	return Response{
		Code: 0,
		Msg:  err.Error(),
		Data: nil,
	}
}

func EmptyResponse() Response {
	return Response{
		Code: -1,
		Msg:  "",
		Data: nil,
	}
}

func Unauthorized() Response {
	return Response{
		Code: CodeInternalErr,
		Msg:  MsgUnauthorized,
		Data: nil,
	}
}

func Internal() Response {
	return Response{
		Code: CodeInternalErr,
		Msg:  MsgInternalErr,
		Data: nil,
	}
}

func NotFound() Response {
	return Response{
		Code: CodeNotFound,
		Msg:  "not found",
		Data: nil,
	}
}

func DebugErr(err error) Response {
	return Response{
		Code: CodeDebug,
		Msg:  err.Error(),
		Data: nil,
	}
}

func BadReq(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, DebugErr(err))
	return
}

func InternalError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, DebugErr(err))
	return
}

func OK(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Success(data))
	return
}