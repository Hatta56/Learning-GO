package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)

	router.GET("/helo", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.Run()
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":  "Hatta ",
		"title": "Student",
	})
}
func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "I LOVE 3000.",
		"subtitle": "apalah apadeh",
	})
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
