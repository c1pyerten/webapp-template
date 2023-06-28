package middleware

import (
	"c1pherten/yet-webapp2/api"
	"c1pherten/yet-webapp2/repository"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// todo:
var secret = "aklfjoijvohwetong,ks "

const authHeader = "Authorization"
const tokenPrefix = "Bearer "

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrUnauthorized = errors.New("unauthorized")
)

type Claim struct {
	jwt.RegisteredClaims
	User *repository.User
}


func Auth(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get(authHeader)
	tokenString = strings.TrimSpace(tokenString)
	if !strings.HasPrefix(tokenString, tokenPrefix) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Unauthorized())
		return
	}

	tokenString = tokenString[len(tokenPrefix):]
	token, err := jwt.ParseWithClaims(tokenString, &Claim{},func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secret), nil
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Unauthorized())
		return
	}
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		ctx.Set("user", claims.User)
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Unauthorized())
		return
	}
}

func newClaim(u *repository.User) *Claim {
	return &Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "webapp",
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		User:             u,
	}
}

func Sign(u *repository.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaim(u))
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}