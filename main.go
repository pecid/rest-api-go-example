package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pecid/rest-api-go-example/cmd/handlers"
	"github.com/pecid/rest-api-go-example/internal/book"
	"github.com/pecid/rest-api-go-example/models"
)

var DB *gorm.DB

func main() {
	server := gin.Default()

	database, error := gorm.Open("sqlite3", "test.db")

	if error != nil {
		fmt.Println("error")
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.Book{})

	DB = database

	bookRepo := book.NewRespository(DB)
	bookServ := book.NewService(bookRepo)
	bookHandler := handlers.NewProduct(bookServ)

	server.GET("/books", bookHandler.GetAll())
	server.POST("/books", bookHandler.Store())

	server.Run(":8082")
}
