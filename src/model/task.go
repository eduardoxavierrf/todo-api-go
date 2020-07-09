package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Name string `json:"name"`
	Done uint   `json:"done"`
}
