package main

import (
	//"fmt"

	"github.com/9Neechan/AvitoTech-2024/internal/handlers"
	"github.com/9Neechan/AvitoTech-2024/internal/middleware"

	//"github.com/9Neechan/AvitoTech-2024/internal/psql"
	//"github.com/9Neechan/AvitoTech-2024/internal/database"

	"github.com/gin-gonic/gin"
)

// 7.26.00

// export GOOSE_DRIVER=postgres
// export GOOSE_DBSTRING="host=localhost port=5432 user=postgres password=123 dbname=banners sslmode=disable"
// goose up

func main() {
	//db := psql.ConnectDB()
	//defer db.Close()

	router := gin.Default()

	//middleware.AuthMiddleware("u", "a", router)

	router.Use(middleware.AuthMiddleware("u", "a", router))
	router.Use(middleware.OnlyAdmin())

	router.GET("/user_banner", handlers.GetBannerForUserHandler)       // Получение баннера для пользователя
	router.GET("/banner", handlers.GetAllBannersByFeatureOrTagHandler) // Получение всех баннеров c фильтрацией по фиче и/или тегу
	router.POST("/banner", handlers.CreateBannerHander)                // Создание нового баннера
	router.PATCH("/banner/:id", handlers.UpdateBannerHandler)          // Обновление содержимого баннера
	router.DELETE("/banner/:id", handlers.DeleteBannerHandler)         // Удаление баннера по идентификатору

	router.POST("/user", handlers.CreateUserHander)

	// Run the Gin server
	router.Run(":8080")
}

//{"{\"name\": \"John\", \"age\": 30}": true}
