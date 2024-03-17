package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Bbarbosa7/ineed-full/app/service-search/handlers"
	"net/http"
)

func main() {
	r := gin.Default()

	// API routes
	r.GET("/service-types", handlers.GetServiceTypesHandler)
	r.Run(":1234")

	// Server the frontend for the Application
	r.Static("/static", "./client/assets")
	r.LoadHTMLGlob("client/*.html")
	r.GET("/", func(c *gin.Context) {
		// Código para renderizar a página inicial (index.html)
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}