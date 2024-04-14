package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/9Neechan/AvitoTech-2024/internal/database"
	"github.com/9Neechan/AvitoTech-2024/internal/middleware"
	"github.com/9Neechan/AvitoTech-2024/internal/models"
	"github.com/9Neechan/AvitoTech-2024/internal/psql"

	"github.com/gin-gonic/gin"
)

// POST banner - Создание нового баннера
func CreateBannerHander(c *gin.Context) {
	type parameters struct {
		TagIDs    []int32         `json:"tag_ids"`
		FeatureID int32           `json:"feature_id"`
		IsActive  bool            `json:"is_active"`
		Content   json.RawMessage `json:"content"`
	}

	// Парсим параметры
	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't decode parameters: " + err.Error()})
		return
	}

	// Создаем баннер в БД
	banner, err := psql.DB.CreateBanner(c, database.CreateBannerParams{
		FeatureID: params.FeatureID,
		IsActive:  params.IsActive,
		Content:   params.Content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't create banner: " + err.Error()})
		return
	}

	// Добавлям к баннеру Теги
	err = psql.DB.InsertTags(c, database.InsertTagsParams{
		BannerID:  banner.ID,
		FeatureID: params.FeatureID,
		Column3:   params.TagIDs,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't create banner: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"banner_id": banner.ID})
}

// GET user_banner - Получение баннера для пользователя
func GetBannerForUserHandler(c *gin.Context) {
	type parameters struct {
		TagID          int32 `json:"tag_id"`
		FeatureID      int32 `json:"feature_id"`
		UseLastVersion int32  `json:"use_last_version"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't decode parameters: " + err.Error()})
		return
	}

	fmt.Println("************", params.UseLastVersion)

	user, ok := c.Value(middleware.UserKey).(models.User)
	if !ok {
		middleware.SendErrorMsg(c, http.StatusInternalServerError, "user not found in context")
		return
	}

	//if params.UseLastVersion {}
	banner, err := psql.DB.GetUserBanner(c, database.GetUserBannerParams{
		TagID:     params.TagID,
		FeatureID: params.FeatureID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't find user in DB: " + err.Error()})
		return
	}

	if !banner.IsActive && !user.IsAdmin {
		c.Status(http.StatusForbidden)
	} else {
		c.JSON(http.StatusOK, banner.Content)
	}

}

// Получение всех баннеров c фильтрацией по фиче и/или тегу
func GetAllBannersByFeatureOrTagHandler(c *gin.Context) {
	type parameters struct {
		TagID     int32 `json:"tag_id"`
		FeatureID int32 `json:"feature_id"`
		Limit     int32 `json:"limit"`
		Offset    int32 `json:"offset"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't decode parameters: " + err.Error()})
		return
	}

	if params.Limit == 0 {
		params.Limit = 5
	}

	if params.FeatureID != 0 && params.TagID != 0 {
		banner, err := psql.DB.GetUserBanner(c, database.GetUserBannerParams{
			TagID:     params.TagID,
			FeatureID: params.FeatureID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "Couldn't find user in DB: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, banner)

	} else if params.FeatureID == 0 { // по TagID
		banners, err := psql.DB.GetBannerListByTag(
			c, 
			database.GetBannerListByTagParams{
				TagID: params.TagID,
				Limit: params.Limit,
				Offset: params.Offset,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "Couldn't find user in DB: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, banners)
	} else { // по FeatureID
		banners, err := psql.DB.GetBannerListByFeatureId(
			c, 
			database.GetBannerListByFeatureIdParams{
				FeatureID: params.FeatureID,
				Limit: params.Limit,
				Offset: params.Offset,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "Couldn't find user in DB: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, banners)
	}
}

// PATCH banner - Обновить содержимое баннера
func UpdateBannerHandler(c *gin.Context) {
	type parameters struct {
		TagIDs    []int32         `json:"tag_ids"`
		FeatureID int32           `json:"feature_id"`
		IsActive  bool            `json:"is_active"`
		Content   json.RawMessage `json:"content"`
	}

	// Парсим параметры
	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't decode parameters: " + err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	banner, err := psql.DB.GetBannerByID(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't find user in DB: " + err.Error()})
		return
	}

	// update tags
	if params.TagIDs != nil {
		err = psql.DB.DeleteTags(c, database.DeleteTagsParams{
			BannerID:  banner.ID,
			FeatureID: params.FeatureID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "Couldn't update tags (while deleting): " + err.Error()})
			return
		}

		err = psql.DB.InsertTags(c, database.InsertTagsParams{
			BannerID:  banner.ID,
			FeatureID: params.FeatureID,
			Column3:   params.TagIDs,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "Couldn't update tags (while insrting): " + err.Error()})
			return
		}
	}

	// update other params
	query_params := parameters{}

	if params.FeatureID != banner.FeatureID {
		query_params.FeatureID = params.FeatureID
	} else {
		query_params.FeatureID = banner.FeatureID
	}

	if params.IsActive != banner.IsActive {
		query_params.IsActive = params.IsActive
	} else {
		query_params.IsActive = banner.IsActive
	}

	if params.Content != nil {
		query_params.Content = params.Content
	} else {
		query_params.Content = banner.Content
	}

	psql.DB.UpdateBanner(c, database.UpdateBannerParams{
		ID:        banner.ID,
		FeatureID: params.FeatureID,
		IsActive:  params.IsActive,
		Content:   params.Content,
	})

	c.Status(http.StatusOK)
}

// DELETE banner - Удаление баннера
func DeleteBannerHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = psql.DB.DeleteBanner(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Couldn't find banner in DB: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
