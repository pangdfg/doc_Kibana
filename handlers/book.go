package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"elasticsearch-basic/config"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
)

var genres = []string{
	"Fantasy",
	"Science",
	"Mystery",
	"Historical",
	"Romance",
	"Horror",
	"Biography",
	"Adventure",
}

func randomGenre() string {
	return genres[rand.Intn(len(genres))]
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API Running",
	})
}

func InitBooks(c *gin.Context) {

	for i := 0; i < 10000; i++ {

		book := Book{
			Title:  gofakeit.ProductName(),
			Author: gofakeit.Name(),
			Genre:  randomGenre(),
		}

		fmt.Printf("position %d title %s\n", i, book.Title)

		data, _ := json.Marshal(book)

		res, err := config.ES.Index(
			"old_books",
			bytes.NewReader(data),
			config.ES.Index.WithContext(context.Background()),
		)

		if err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		res.Body.Close()
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Books indexed!",
	})
}