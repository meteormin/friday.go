package entity

import "gorm.io/gorm"

type Site struct {
	gorm.Model
	Name string `gorm:"unique"`
	Host string `gorm:"unique"`
}
