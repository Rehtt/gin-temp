package models

import "gorm.io/gorm"

type UserTables struct {
	*gorm.Model
	Name string `gorm:"not null"`
}
