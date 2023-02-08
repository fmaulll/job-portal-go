package main

import (
	"github.com/fmaulll/job-portal-go/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.Migrate()
}

func main() {
	router := gin.Default()

	router.GET("/")
}
