package api

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Book struct {
// 	IDBook      string `json:"ID-Book"`
// 	Title       string `json:"Title"`
// 	Author      string `json:"Author"`
// 	Description string `json:"Desc"`
// }

// var Data = []Book{}

// func AddBook(RouteContext *gin.Context) {
// 	var AddedBook Book

// 	if err := RouteContext.ShouldBind(&AddedBook); err != nil {
// 		RouteContext.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	Data = append(Data, AddedBook)

// 	RouteContext.JSON(http.StatusCreated, gin.H{
// 		"Book": AddedBook,
// 	})
// }

// func UpdateBook(RouteContext *gin.Context) {
// 	ReferenceBook := RouteContext.Param("ID")
// 	var UpdateBook Book

// 	if err := RouteContext.ShouldBind(&UpdateBook); err != nil {
// 		RouteContext.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	State := false
// 	for i, requirement := range Data {
// 		if requirement.IDBook == ReferenceBook {
// 			Data[i] = UpdateBook
// 			Data[i].IDBook = ReferenceBook
// 			Data[i].Title = UpdateBook.Title
// 			Data[i].Author = UpdateBook.Author
// 			Data[i].Description = UpdateBook.Description
// 			State = true
// 		}
// 	}

// 	if State {
// 		RouteContext.JSON(http.StatusFound, gin.H{
// 			"Status":         "Updated successfully",
// 			"Reference Book": ReferenceBook,
// 			"Book":           Data,
// 		})
// 	} else {
// 		RouteContext.JSON(http.StatusNotFound, gin.H{
// 			"Status":  http.StatusNotFound,
// 			"Message": fmt.Sprintf("Book ID %v not Found", ReferenceBook),
// 		})
// 	}
// }

// func DeteleBook(RouteContext *gin.Context) {
// 	ReferenceBook := RouteContext.Param("ID")
// 	Index := 0
// 	State := false
// 	for i, requirement := range Data {
// 		if ReferenceBook == requirement.IDBook {
// 			State = true
// 			Index = i
// 		}
// 	}
// 	if State {
// 		copy(Data[Index:], Data[Index+1:])
// 		Data[len(Data)-1] = Book{}
// 		Data = Data[:len(Data)-1]
// 		RouteContext.JSON(http.StatusOK, gin.H{
// 			"Status":         "Delete successfully",
// 			"Reference Book": ReferenceBook,
// 		})
// 	}
// }

// func GetAllBook(RouteContext *gin.Context) {
// 	if len(Data) >= 1 {
// 		RouteContext.JSON(http.StatusFound, gin.H{
// 			"Book": Data,
// 		})
// 	} else {
// 		RouteContext.JSON(http.StatusNotFound, gin.H{
// 			"Status":  http.StatusNotFound,
// 			"Message": "GET Book not found",
// 		})
// 	}
// }

// func GetBookByID(RouteContext *gin.Context) {
// 	// Fetch Data from Local Variables using params given
// 	CheckID := RouteContext.Param("ID")
// 	// Make Local Var for temp Data so not overwrite
// 	var TempData Book
// 	var State bool
// 	// Loop Data Here
// 	for i, requirement := range Data {
// 		if CheckID == requirement.IDBook {
// 			// I Assume Right now it just Array
// 			TempData = Data[i]
// 			State = true
// 		}
// 	}

// 	if State {
// 		RouteContext.JSON(http.StatusFound, gin.H{
// 			"Book": TempData,
// 		})
// 	} else {
// 		RouteContext.JSON(http.StatusNotFound, gin.H{
// 			"Status":  http.StatusNotFound,
// 			"Message": fmt.Sprintf("Book ID %v not Found", CheckID),
// 		})
// 	}
// }
