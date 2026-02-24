package main

import (
	"sispa-backend/cmd/server"
	"sispa-backend/internal/infrastructure/databases"
	"sispa-backend/internal/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := databases.InitMariaDB()

	handlers := server.InitHandlers(db)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Vue's address
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.Use(middleware.TimeoutMiddleware(5 * time.Second))

	handlers.CustomerHandler.RegisterRoutes(router)
	handlers.ServiceHandler.RegisterRoutes(router)
	handlers.TherapistHandler.RegisterRoutes(router)
	router.Run(":8080")
}
