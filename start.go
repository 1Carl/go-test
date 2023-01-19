package main

import (
	"github/1Carl/go-test/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/books", func(ctx *gin.Context) {
		var bookToAdd *book.Book
		ctx.BindJSON(bookToAdd)
		if book, err := book.AddBook(*bookToAdd); err != nil {
			// todo return error
		} else {
			ctx.JSON(http.StatusOK, book)
		}
	})

	r.GET("/books/:id", func(ctx *gin.Context) {
		// var bookinfo *struct {
		// 	BookId int
		// }

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// response error
			return
		}
		// ctx.BindJSON(bookinfo)

		if book, err := book.Get(id); err != nil {
			//response error
			return
		} else {
			ctx.JSON(200, book)
		}
	})
}
