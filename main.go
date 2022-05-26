package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

// basic api structure
type allthing struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

var t = time.Now()

// basic api payload
var allthings = []allthing{
	{Message: "Automate all the things!", Timestamp: t.UTC()},
}

// get the api payload
func getAllthings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allthings)
}

func main() {
	router := gin.Default()
	// set a trusted proxy
	router.SetTrustedProxies([]string{"192.168.0.1"})

	router.GET("/allthings", getAllthings)

	// set root to getAllthings as well
	router.GET("/", getAllthings)

	router.Run("localhost:8080")
}
