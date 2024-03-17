package main

import (
	"github.com/Bbarbosa7/ineed-full/app/service-search/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// API routes
	r.GET("/service-types", handlers.GetServiceTypesHandler)

	// Server the frontend for the Application
	r.Static("/static", "./client/assets")
	r.LoadHTMLGlob("client/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Run the server
	r.Run(":1234")
}
