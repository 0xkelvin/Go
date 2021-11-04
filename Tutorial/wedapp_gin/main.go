package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by GIN
	router = gin.Default()

	// Process the templates at the start so that they dont have to be loaded
	// from the disk again. This makes serving HTML pages very fast.

	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler, Later one, we'll create
	// standalone functions that will be used as route handlers
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the context to render a template
		c.HTML(
			// Set the HTTP status to 200 OK
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses
			gin.H{
				"title": "Home Page",
			},
		)
	})
	// start serving the application
	router.Run()
}
