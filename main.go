package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Bbarbosa7/ineed-full/app/service-search/handlers"
)

func main() {
	r := gin.Default()

	// Rotas da API
	r.GET("/users", handlers.GetUserHandler)
	r.POST("/users", handlers.CreateUserHandler)
	r.DELETE("/users/:id", handlers.DeleteUserHandler)

	// Servir frontend
	r.Static("/static", "./frontend/assets")
	r.LoadHTMLGlob("frontend/*.html")
	r.GET
