package main

import (
	"WebApi/book"
	"WebApi/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/webapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&book.Books{})

	/*
		Find All Books with repository
		bookRepository := book.NewRepository(db)
		books, err := bookRepository.FindAll()
		if err != nil {
			log.Fatal(err)
		}
		for _, book := range books {
			fmt.Println("Books Title :", book.Title)
			fmt.Println(book)
		}
	*/

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	/*
		== insert a record ==
		book := book.Books{}
		book.Title = "Go Programming Language Update 2"
		book.Price = 100
		book.Description = "Go Programming Language is a programming language that makes it easy to build simple, reliable, and efficient software."
		book.Rating = 5
		book.Discount = 10

		err = db.Create(&book).Error
		if err != nil {
			log.Fatal(err)
		}
	*/
	/*
			== Finding a record by ID ==
		var books book.Books
		err = db.Debug().First(&books, 2).Error
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(books.Title)
		fmt.Println(books)
	*/

	/*
		Get all records
		var books []book.Books
		err = db.Debug().Find(&books).Error
		if err != nil {
			log.Fatal(err)
		}
		for _, book := range books {
			fmt.Println("Book Title", book.Title)
			fmt.Println(book)
		}
	*/
	// searching records by rating
	/*
		var books []book.Books
		err = db.Debug().Where("rating= ?", 5).Find(&books).Error
		if err != nil {
			log.Fatal(err)
		}
		for _, book := range books {
			fmt.Println("Book Title", book.Title)
			fmt.Println(book)
		}
	*/
	/*
				== Update a record ==
		var books book.Books
		err = db.Debug().Where("id = ?", 1).First(&books).Error
		if err != nil {
			log.Fatal(err)
		}
		books.Title = "Go Programming Language (updated)"
		err = db.Debug().Save(&books).Error
		if err != nil {
			log.Fatal(err)
		}
	*/
	/*
		var books book.Books
		err = db.Debug().Where("id = ?", 2).First(&books).Error
		if err != nil {
			log.Fatal(err)
		}
		err = db.Debug().Delete(&books).Error
		if err != nil {
			log.Fatal(err)
		}
	*/
	router := gin.Default()
	v1 := router.Group("/v1")
	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("books/:id/:title", bookHandler.BooksHandler)
	// v1.GET("query", bookHandler.QueryHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)
	router.Run(":8000")

}
