package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

// easter egg structure
type easteregg struct {
	Message string `json:"message"`
}

// alltings api structure
type allthing struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// for the timestamp
var t = time.Now()

// easter egg payload
var eastereggs = []easteregg{
	{Message: "Hello there! this is the easter egg. The required stuff is at root."},
}

// allthings payload
var allthings = []allthing{
	{Message: "Automate all the things!", Timestamp: t.UTC()},
}

// get the root payload
func getEasterEgg(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, eastereggs)
}

// get the api payload
func getAllthings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allthings)
}

func main() {
	router := gin.Default()
	// set a trusted proxy
	router.SetTrustedProxies([]string{"128.0.0.1/24"})

	router.GET("/", getAllthings)
	router.GET("/easteregg", getEasterEgg)

	router.Run("localhost:8180")
}
