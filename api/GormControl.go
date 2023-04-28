package api

import (
	"api-go/handler"
	"api-go/handler/database"
	"api-go/handler/structures"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBook(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		ResponseContext.JSON(http.StatusInternalServerError, gin.H{
			"Error":   handler.BaseError(err, http.StatusInternalServerError, "Database Failed to initialize"),
			"Message": "Failed on Database",
		})
	}
	Book := []structures.Book{}
	QueryCheck := DB.Find(&Book).Error
	if QueryCheck != nil {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Messsage": "Book not found",
		})
		return
	} else {
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Data": Book,
		})
		return
	}

}

func GetOneBook(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		ResponseContext.JSON(http.StatusInternalServerError, gin.H{
			"Error":   handler.BaseError(err, http.StatusInternalServerError, "Database Failed to initialize"),
			"Message": "Failed on Database",
		})
	}
	Book := structures.Book{}
	Param, _ := strconv.Atoi(ResponseContext.Param("ID"))
	QueryCheck := DB.Where("id = ?", Param).Find(&Book).Error
	if QueryCheck != nil {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Messsage": "Book not found",
		})
		return
	} else {
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Data": Book,
		})
		return
	}
}

func UpdatedBook(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		ResponseContext.JSON(http.StatusInternalServerError, gin.H{
			"Error":   handler.BaseError(err, http.StatusInternalServerError, "Database Failed to initialize"),
			"Message": "Failed on Database",
		})
	}
	// Check Temp Data
	Temps := structures.Temp{}
	Book := structures.Book{}

	Param, _ := strconv.Atoi(ResponseContext.Param("ID"))

	if ResponseContext.GetHeader("Content-Type") == "application/json" {
		ResponseContext.ShouldBindJSON(&Book)
	} else {
		ResponseContext.ShouldBind(&Book)
	}

	QueryFind := DB.Model(&Book).Where("id = ?", Param).Take(&Temps)

	if QueryFind.RowsAffected == 0 || QueryFind.RowsAffected == -1 {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Messsage": "Book not found",
		})
		return
	} else {
		Book.IDBook = Temps.IDBook
	}

	QueryCheck := DB.Model(&Book).Where("ID = ?", Param).Where("id_book = ?", Book.IDBook).Updates(structures.Book{Title: Book.Title, Author: Book.Author, Description: Book.Description}).Take(&Book)
	if QueryCheck.RowsAffected == 1 {
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Data": Book,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Messsage": "Book Failed to Updated",
		})
		return
	}
}

func CreateBook(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		ResponseContext.JSON(http.StatusInternalServerError, gin.H{
			"Error":   handler.BaseError(err, http.StatusInternalServerError, "Database Failed to initialize"),
			"Message": "Failed on Database",
		})
	}

	Book := structures.Book{}

	if ResponseContext.GetHeader("Content-Type") == "application/json" {
		ResponseContext.ShouldBindJSON(&Book)
	} else {
		ResponseContext.ShouldBind(&Book)
	}

	QueryCheck := DB.Create(&Book).Error
	if QueryCheck != nil {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Messsage": "Failed To Add Book",
		})
		return
	} else {
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Data": Book,
		})
		return
	}
}
func DeletedBook(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		ResponseContext.JSON(http.StatusInternalServerError, gin.H{
			"Error":   handler.BaseError(err, http.StatusInternalServerError, "Database Failed to initialize"),
			"Message": "Failed on Database",
		})
	}
	Temps := structures.Temp{}
	Book := structures.Book{}
	Param, _ := strconv.Atoi(ResponseContext.Param("ID"))

	if ResponseContext.GetHeader("Content-Type") == "application/json" {
		ResponseContext.ShouldBindJSON(&Book)
	} else {
		ResponseContext.ShouldBind(&Book)
	}

	QueryFind := DB.Model(&Book).Where("id = ?", Param).Take(&Temps)

	if QueryFind.RowsAffected == 0 || QueryFind.RowsAffected == -1 {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Messsage": "Book not found",
		})
		return
	}

	QueryCheck := DB.Debug().Where("ID = ?", Param).Where("IDBook = ?", Book.IDBook).Delete(&Book)
	if QueryCheck.RowsAffected == 1 {
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Data": Book,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Messsage": "Failed to Delete Book",
		})
		return
	}
}
