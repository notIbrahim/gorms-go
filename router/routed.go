package router

import (
	modules "api-go/api"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	route := gin.Default()

	route.POST("/books", modules.CreateBook)
	route.GET("/books", modules.GetBook)
	route.GET("/books/:ID", modules.GetOneBook)
	route.PUT("/books/:ID", modules.UpdatedBook)
	route.DELETE("/books/:ID", modules.DeletedBook)
	return route
}
