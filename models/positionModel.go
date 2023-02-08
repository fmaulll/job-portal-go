package models

import "time"

type Position struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Steps       string    `json:"steps"`
	LogoUrl     string    `json:"logoUrl"`
	FullTime    string    `json:"fullTime"`
}
