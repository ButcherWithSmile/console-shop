package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type console struct {
	ID string `json:"id"`
	Brand string `josn:"brand"`
	Model string `json:"model"`
	Price float64 `json:"price"`
}

var consoles = []console{
	{ID: "1", Brand: "Sony", Model: "PlayStation 4 Pro, 1TB, Black", Price: 239.99},
	{ID: "2", Brand: "Sony", Model: "PlayStation 5", Price: 459.99},
	{ID: "3", Brand: "Microsoft", Model: "Xbox Series X", Price: 469.99},
	{ID: "4", Brand: "Microsoft", Model: "Xbox Series S 1TB, Black", Price: 349.99},
}

func getConsoles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, consoles)
}

func postConsoles(c *gin.Context) {
	var newCosole console

	if err := c.BindJSON(&newCosole); err != nil {
		return
	}

	consoles = append(consoles, newCosole)
	c.IndentedJSON(http.StatusCreated, newCosole)
}

func getConsoleByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range consoles {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "console not found"})
}

func main() {
	router := gin.Default()
	router.GET("/consoles", getConsoles)
	router.GET("/consoles/:id", getConsoleByID)
	router.POST("/consoles", postConsoles)

	router.Run("localhost:8880")
}