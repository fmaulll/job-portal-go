package initializers

import "github.com/fmaulll/job-portal-go/models"

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Position{})
}
