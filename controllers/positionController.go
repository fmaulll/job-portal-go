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
	fullTime := context.Query("full_time")

	if fullTime == "true" {
		fullTime = "Full Time"
	}

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	var totalRows int64

	offset := (page - 1) * perPage

	var position []models.Position

	result := initializers.DB.Where("LOWER(title) LIKE ? AND LOWER(location) LIKE ? AND type LIKE ?", "%"+description+"%", "%"+location+"%", "%"+fullTime+"%").Order("title").Limit(10).Offset(offset).Find(&position).Count(&totalRows)

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

	result := initializers.DB.Where("id = ?", id).Find(&position)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Position not found"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"result": position})
}
