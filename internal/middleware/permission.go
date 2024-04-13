package middleware

import (
	"encoding/json"
	"github.com/9Neechan/AvitoTech-2024/internal/database"
	"github.com/gin-gonic/gin"

	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"
)

type errorMsg struct {
	ErrMsg string `json:"error"`
}

func SendErrorMsg(c *gin.Context, status int, msg string) {
	body, err := json.Marshal(errorMsg{ErrMsg: msg})
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set(contentTypeHeader, contentTypeJSON)
	c.Writer.WriteHeader(status)

	_, err = c.Writer.Write(body)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Value(UserKey).(database.User)
		if !ok {
			SendErrorMsg(c, http.StatusInternalServerError, "user not found in context")
			return
		}

		if !user.IsAdmin {
			c.Writer.WriteHeader(http.StatusForbidden)
		}

		c.Next()
	}
}
