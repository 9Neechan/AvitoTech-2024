package main

import (
	//"fmt"
	//"net/http"

	"github.com/9Neechan/AvitoTech-2024/internal/handlers"
	"github.com/9Neechan/AvitoTech-2024/internal/middleware"

	//"github.com/9Neechan/AvitoTech-2024/internal/psql"
	//"github.com/9Neechan/AvitoTech-2024/internal/database"

	"github.com/gin-gonic/gin"
)

// export GOOSE_DRIVER=postgres
// export GOOSE_DBSTRING="host=localhost port=5432 user=postgres password=123 dbname=banners sslmode=disable"
// goose up

func main() {
	//db := psql.ConnectDB()
	//defer db.Close()

	router := gin.Default()

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware("user", "adm"))
	{
		authorized.GET("/user_banner", handlers.GetBannerForUserHandler)                               // Получение баннера для пользователя
		authorized.GET("/banner", middleware.OnlyAdmin(), handlers.GetAllBannersByFeatureOrTagHandler) // Получение всех баннеров c фильтрацией по фиче и/или тегу
		authorized.POST("/banner", middleware.OnlyAdmin(), handlers.CreateBannerHander)                // Создание нового баннера
		authorized.PATCH("/banner/:id", middleware.OnlyAdmin(), handlers.UpdateBannerHandler)          // Обновление содержимого баннера
		authorized.DELETE("/banner/:id", middleware.OnlyAdmin(), handlers.DeleteBannerHandler)         // Удаление баннера по идентификатору
	}

	// Run the Gin server
	router.Run(":8080")
}

//{"{\"name\": \"John\", \"age\": 30}": true}
