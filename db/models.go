package db

import (
	"github.com/jinzhu/gorm"
)

// Site is the ORM model for database table Site
// Site and User (stored in Auth0) is a one-to-one relationship
// Sitename should be unique headr-wide
// Theme is the name of the theme the site is using
type Site struct {
	gorm.Model
	UserID   string `json:"user_id"`
	Sitename string `json:"sitename"`
	Theme    string `json:"theme"`
}
