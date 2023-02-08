package controllers

import (
	"math"
	"net/http"

	"strconv"

	"github.com/fmaulll/job-portal-go/initializers"
	"github.com/fmaulll/job-portal-go/models"
	"github.com/gin-gonic/gin"
)

func GetPosition(context *gin.Context) {
	perPage := 10
	page := 1
	pageStr := context.Query("page")
	description := context.Query("description")
	location := context.Query("location")
	fullTime := context.Query("fullTime")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	var totalRows int64

	offset := (page - 1) * perPage

	var position []models.Position

	result := initializers.DB.Where("LOWER(title) LIKE ? AND LOWER(location) LIKE ? AND full_time LIKE ?", description+"%", location+"%", fullTime+"%").Order("title").Limit(10).Offset(offset).Find(&position).Count(&totalRows)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Position not found"})

		return
	}

	totalPages := math.Ceil(float64(totalRows) / float64(perPage))

	context.JSON(http.StatusOK, gin.H{"results": position, "totalPages": totalPages, "totalRows": totalRows})
}

func GetPositionDetail(context *gin.Context) {
	id := context.Param("id")

	var position models.Position

	result := initializers.DB.First(&position, id)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Position not found"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"result": position})
}

func AddPosition(context *gin.Context) {
	var body struct {
		Company     string
		Location    string
		Title       string
		Description string
		Steps       string
		FullTime    string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})

		return
	}

	position := models.Position{Company: body.Company, Location: body.Location, Title: body.Title, Description: body.Description, Steps: body.Steps, LogoUrl: "", FullTime: body.FullTime}

	result := initializers.DB.Create(&position)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to add employee"})

		return

	}

	context.JSON(http.StatusCreated, gin.H{"message": "Position " + body.Title + " at " + body.Company + " created"})
}
