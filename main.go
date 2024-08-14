package main

import (
	"fmt"
	"net/http"

	// add gin package
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// add handlers package
	"qrCode/pkg/auth"
	"qrCode/pkg/handlers"
)

func main() {
	// create a new gin router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Update with your frontend origin
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// define a route
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("Hello, World!")
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/api/login", handlers.Login)
	r.POST("/api/register", handlers.Register)
	r.POST("/api/Generate", handlers.Generate)
	r.GET("/api/qr/:site", handlers.GO)
	r.GET("/api/users/:user", handlers.GetUser)
	r.GET("/api/mySites", auth.AuthMiddleware(), handlers.GetSites)
	http.ListenAndServe(":8080", r)
}
