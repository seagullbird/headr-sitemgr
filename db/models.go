package db

import (
	"github.com/jinzhu/gorm"
)

type Site struct {
	gorm.Model
	UserId   uint   `json:"user_id"`
	Sitename string `json:"sitename"`
	Theme    string `json:"theme"`
}
