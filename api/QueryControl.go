package api

// import (
// 	Errno "api-go/handler"
// 	common "api-go/handler/common/response"
// 	"api-go/handler/database"
// 	"api-go/handler/entity"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// var db, _ = database.Connect()
// var QueryBook = entity.Book{}

// func AddedBook(ResponseContext *gin.Context) {
// 	var AddedBook entity.Book

// 	if err := ResponseContext.ShouldBind(&AddedBook); err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorFailedToBindJSON, "Failed to Bind JSON"),
// 		})
// 	}

// 	// Set SQL Statement MYSQL Thing
// 	// SQL, err := db.Prepare("INSERT INTO book (IDBook, Title, Author, Description) VALUES (?, ?, ?,?)")
// 	SQL, err := db.Prepare("INSERT INTO book (Title, Author, Description) VALUES ($1, $2, $3)")
// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseFailedStatement, "Failed to Initialize Statement"),
// 		})
// 	}

// 	res, err := SQL.Exec(&AddedBook.IDBook, &AddedBook.Title, &AddedBook.Author, &AddedBook.Description)
// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseInsertion, "Failed to Insert"),
// 		})

// 	}

// 	affected, err := res.RowsAffected()
// 	if err != nil {
// 		return
// 	}

// 	ResponseContext.JSON(http.StatusCreated, gin.H{
// 		"Status":       http.StatusCreated,
// 		"Database":     common.Success,
// 		"Result Query": affected,
// 	})
// }

// func GetAllBook(ResponseContext *gin.Context) {
// 	var GetBook = []entity.Book{}

// 	// SQL Statement : Select All (Which is Bad)

// 	SQL := "SELECT * FROM book"
// 	SQLRes, err := db.Query(SQL)

// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseQueryExecution, "Failed to Execute Query"),
// 		})
// 	}

// 	for SQLRes.Next() {
// 		// Guess Scan Need Pointer to read after all so need Reference Memory
// 		err = SQLRes.Scan(&QueryBook.IDBook, &QueryBook.Title, &QueryBook.Author, &QueryBook.Description)
// 		if err != nil {
// 			ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 				"Error": Errno.BaseError(err, common.ErrorDatabaseFailure, ""),
// 			})
// 		}

// 		GetBook = append(GetBook, QueryBook)
// 	}

// 	ResponseContext.JSON(http.StatusOK, gin.H{
// 		"Book": GetBook,
// 	})
// }

// func UpdateBook(ResponseContext *gin.Context) {
// 	Reference := ResponseContext.Param("ID")
// 	var UpdateBook entity.Book

// 	if err := ResponseContext.ShouldBind(&UpdateBook); err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorFailedToBindJSON, "Failed to Bind JSON"),
// 		})
// 	}

// 	// Quick Param Builder if one of those Parameter not Filled
// 	// Title = ?,  Author = ?, Description = ? WHERE IDBook = ?
// 	SQLUpdate := "UPDATE book SET "
// 	Param := []interface{}{}

// 	if UpdateBook.Title != "" {
// 		SQLUpdate += "Title = ?, "
// 		Param = append(Param, *&UpdateBook.Title)
// 	}

// 	if UpdateBook.Author != "" {
// 		SQLUpdate += "Author = ?, "
// 		Param = append(Param, *&UpdateBook.Author)
// 	}

// 	if UpdateBook.Description != "" {
// 		SQLUpdate += "Description = ?, "
// 		Param = append(Param, *&UpdateBook.Description)
// 	}

// 	SQLUpdate = strings.TrimSuffix(SQLUpdate, ", ") + " WHERE IDBook = ?"
// 	Param = append(Param, Reference)

// 	SQL, err := db.Prepare(SQLUpdate)
// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseQueryExecution, "Failed to Execute Query"),
// 		})
// 	}

// 	res, err := SQL.Exec(Param...)
// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseUpdate, "Failed to Update Book"),
// 		})
// 	}

// 	affected, err := res.RowsAffected()
// 	if err != nil {
// 		return
// 	}

// 	ResponseContext.JSON(http.StatusAccepted, gin.H{
// 		"Status":       http.StatusAccepted,
// 		"Database":     common.Success,
// 		"Result Query": affected,
// 	})
// }
// func GetBookByID(ResponseContext *gin.Context) {
// 	var OneBook = []entity.Book{}
// 	Reference := ResponseContext.Param("ID")
// 	SQLFind := "SELECT * FROM book WHERE IDBook = ?" // Bad Query do not consider

// 	SQL, err := db.Prepare(SQLFind)
// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseQueryExecution, "Failed to Execute Query"),
// 		})
// 	}
// 	SQLRes, _ := SQL.Query(Reference)
// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseFind, "Failed to Update Book"),
// 		})
// 	}

// 	for SQLRes.Next() {
// 		err = SQLRes.Scan(&QueryBook.IDBook, &QueryBook.Title, &QueryBook.Author, &QueryBook.Description)
// 		if err != nil {
// 			ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 				"Error": Errno.BaseError(err, common.ErrorDatabaseFailure, "Failed to Receive Data"),
// 			})
// 		}

// 		OneBook = append(OneBook, QueryBook)
// 	}
// 	ResponseContext.JSON(http.StatusFound, gin.H{
// 		"Status":         http.StatusFound,
// 		"Database":       common.Success,
// 		"Reference Book": OneBook,
// 	})
// 	return
// }

// func DeleteBook(ResponseContext *gin.Context) {
// 	Reference := ResponseContext.Param("ID")
// 	SQLFind := "DELETE FROM book WHERE IDBook = ?"
// 	SQL, err := db.Prepare(SQLFind)
// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseQueryExecution, "Failed to Execute"),
// 		})
// 	}

// 	SQLRes, err := SQL.Exec(Reference)
// 	if err != nil {
// 		ResponseContext.JSON(http.StatusBadRequest, gin.H{
// 			"Error": Errno.BaseError(err, common.ErrorDatabaseDelete, "Failed to Delete Book Reference"),
// 		})
// 	}
// 	affected, err := SQLRes.RowsAffected()
// 	if err != nil {
// 		return
// 	}
// 	ResponseContext.JSON(http.StatusAccepted, gin.H{
// 		"Status":       http.StatusAccepted,
// 		"Database":     common.Success,
// 		"Result Query": affected,
// 	})
// }
