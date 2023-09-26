package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This struct defines the model for a console. The josn:"brand" tag tells Gin how to serialize the Brand field to JSON
type console struct {
	ID string `json:"id"`
	Brand string `josn:"brand"`
	Model string `json:"model"`
	Price float64 `json:"price"`
}

// This variable stores a slice of consoles. This is the data that the API will expose
var consoles = []console{
	{ID: "1", Brand: "Sony", Model: "PlayStation 4 Pro, 1TB, Black", Price: 239.99},
	{ID: "2", Brand: "Sony", Model: "PlayStation 5", Price: 459.99},
	{ID: "3", Brand: "Microsoft", Model: "Xbox Series X", Price: 469.99},
	{ID: "4", Brand: "Microsoft", Model: "Xbox Series S 1TB, Black", Price: 349.99},
}

// This function handles GET requests to the /consoles endpoint
// It returns a JSON response containing all of the consoles
func getConsoles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, consoles)
}

// This function handles POST requests to the /consoles endpoint
// It creates a new console from the JSON body of the request and returns a JSON response containing the new console
func postConsoles(c *gin.Context) {
	var newCosole console

	if err := c.BindJSON(&newCosole); err != nil {
		return
	}

	consoles = append(consoles, newCosole)
	c.IndentedJSON(http.StatusCreated, newCosole)
}

// This function handles GET requests to the /consoles/:id endpoint
// It gets the console with the specified ID and returns a JSON response containing the console
// If the console is not found, it returns a JSON response with a 404 status code
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
