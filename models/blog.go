package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Image  string `json:"image"`
	Title string `json:"title"`
	Sub string `json:"sub"`
	Url string `json:"url"`
}
