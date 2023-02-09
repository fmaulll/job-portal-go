package models

type Position struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Url         string `json:"url"`
	Type        string `json:"type"`
	CreatedAt   string `json:"createdAt"`
	Company     string `json:"company"`
	CompanyUrl  string `json:"companyUrl"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"howToApply"`
	CompanyLogo string `json:"companyLogo"`
}
