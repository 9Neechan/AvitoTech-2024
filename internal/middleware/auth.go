package middleware

import (
	"net/http"
	//"context"

	usermodels "github.com/9Neechan/AvitoTech-2024/internal/models"

	"github.com/gin-gonic/gin"
)

const headerTokenName = "token"

//type userKeyT string
//const UserKey userKeyT = "user key"

const UserKey string = "user key"

func AuthMiddleware(userToken string, adminToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user usermodels.User

		token := c.GetHeader(headerTokenName)
		//token := c.Request.Header.Get(headerTokenName)
		switch token {
		case userToken:
			user = usermodels.User{IsAdmin: false}
		case adminToken:
			user = usermodels.User{IsAdmin: true}
		default:
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		//ctx := context.WithValue(c, UserKey, user)
		//next.ServeHTTP(c.Writer, c.Request.WithContext(ctx))

		c.Set(UserKey, user)
		c.Next()
	}
}