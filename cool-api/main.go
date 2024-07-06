package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type cool struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Points int     `json:"points"`
	Price  float64 `json:"price"`
}

var coolItems = []cool{
	{ID: "1", Name: "Cool Item 1", Points: 100, Price: 10.00},
	{ID: "2", Name: "Cool Item 2", Points: 200, Price: 20.00},
	{ID: "3", Name: "Cool Item 3", Points: 300, Price: 30.00},
	{ID: "4", Name: "Cool Item 4", Points: 400, Price: 40.00},
}

func getCoolItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, coolItems)
}

func main() {
	router := gin.Default()
	router.GET("/items", getCoolItems)
	router.Run("localhost:8080")
}
