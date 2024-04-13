package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/9Neechan/AvitoTech-2024/internal/database"
	"github.com/9Neechan/AvitoTech-2024/internal/psql"
	"github.com/9Neechan/AvitoTech-2024/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUserHander(c *gin.Context) {
	type parameters struct {
		IsAdmin bool `json:"is_admin"`
	}
	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't decode parameters" + err.Error()})
		//respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := psql.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		IsAdmin:   params.IsAdmin,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't create user" + err.Error()})
		//respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	c.JSON(http.StatusOK, models.DatabaseUserToUser(user))
	//c.JSON(http.StatusOK, user)
}
