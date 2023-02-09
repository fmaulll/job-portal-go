package main

import (
	"os"

	"github.com/fmaulll/job-portal-go/controllers"
	"github.com/fmaulll/job-portal-go/initializers"
	"github.com/fmaulll/job-portal-go/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.Migrate()
}

func main() {
	router := gin.Default()
	router.Static("/images/", "./images")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	// router.Use(cors.Default())

	router.POST("/api/signup", controllers.Signup)
	router.POST("/api/login", controllers.Login)

	router.GET("/api/position", middleware.RequireAuth, controllers.GetPosition)
	router.GET("/api/position/:id", middleware.RequireAuth, controllers.GetPositionDetail)

	router.Run(":" + os.Getenv("PORT"))
}
