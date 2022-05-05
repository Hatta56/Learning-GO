package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"WebApi/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookHandler struct {
	BookService book.Service
}

func NewBookHandler(bookService book.Service) *BookHandler {
	return &BookHandler{bookService}
}

func convertBookResponse(b book.Books) book.BookResponse {
	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
	return bookResponse
}

func (h *BookHandler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	bookID, _ := strconv.Atoi(id)
	books, err := h.BookService.FindByID(bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookResponse := convertBookResponse(books)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *BookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.BookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (handler *BookHandler) CreateBook(ctx *gin.Context) {
	var bookRequest book.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)
	errorMessages := []string{}
	if err != nil {
		//error input validation
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s ", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}
	book, err := handler.BookService.Create(bookRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": convertBookResponse(book),
	})
}

func (handler *BookHandler) UpdateBook(ctx *gin.Context) {
	var BookRequest book.BookRequest

	err := ctx.ShouldBindJSON(&BookRequest)
	errorMessages := []string{}
	if err != nil {
		//error input validation
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s ", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}
	id := ctx.Param("id")
	bookID, _ := strconv.Atoi(id)
	book, err := handler.BookService.Update(bookID, BookRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": convertBookResponse(book),
	})
}

func (h *BookHandler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	bookID, _ := strconv.Atoi(id)
	books, err := h.BookService.Delete(bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookResponse := convertBookResponse(books)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// func (handler *BookHandler) QueryHandler(ctx *gin.Context) {
// 	title := ctx.Query("title")
// 	price := ctx.Query("price")
// 	ctx.JSON(http.StatusOK, gin.H{"title": title, "price": price})
// }

// func (handler *BookHandler) BooksHandler(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	title := ctx.Param("title")
// 	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
// }

// func (handler *BookHandler) RootHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"name":      "Hatta",
// 		"biography": "I am a software engineer",
// 	})
// }

// func (handler *BookHandler) HelloHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"title":    "Hello World",
// 		"subtitle": "I am a software engineer",
// 	})
// }
