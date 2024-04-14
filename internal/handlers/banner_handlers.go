package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/9Neechan/AvitoTech-2024/internal/database"
	"github.com/9Neechan/AvitoTech-2024/internal/models"
	"github.com/9Neechan/AvitoTech-2024/internal/psql"

	"github.com/gin-gonic/gin"
)

// POST banner - Создание нового баннера
func CreateBannerHander(c *gin.Context) {
	type parameters struct {
		Feature     int32           `json:"feature"`
		Tag         int32           `json:"tag"`
		JsonContent json.RawMessage `json:"json_content"`
		IsActive    bool            `json:"is_active"`
	}
	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't decode parameters: " + err.Error()})
		//respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	banner, err := psql.DB.CreateBanner(c, database.CreateBannerParams{
		Feature:     params.Feature,
		Tag:         params.Tag,
		JsonContent: params.JsonContent,
		IsActive:    params.IsActive,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't create banner: " + err.Error()})
		//respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	c.JSON(http.StatusOK, models.DatabaseBannerToBanner(banner))
}

func GetBannerForUserHandler(c *gin.Context) {
	
	
}

// Получение всех баннеров c фильтрацией по фиче и/или тегу
func GetAllBannersByFeatureOrTagHandler(c *gin.Context) {

}

func UpdateBannerHandler(c *gin.Context) {

}

func DeleteBannerHandler(c *gin.Context) {

}
