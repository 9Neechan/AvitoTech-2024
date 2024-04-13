package middleware

import (
	"net/http"
	"context"

	usermodels "github.com/9Neechan/AvitoTech-2024/internal/models"

	"github.com/gin-gonic/gin"
)

const headerTokenName = "token"

type userKeyT string

const UserKey userKeyT = "user key"

func AuthMiddleware(userToken string, adminToken string, next *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user usermodels.User

		token := c.Request.Header.Get(headerTokenName)
		switch token {
		case userToken:
			user = usermodels.User{IsAdmin: false}
		case adminToken:
			user = usermodels.User{IsAdmin: true}
		default:
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		//c.Set(UserKey, user)
		ctx := context.WithValue(c, UserKey, user)

		//c.Next()
		next.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}