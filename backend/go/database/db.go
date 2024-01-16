package database

import "gorm.io/gorm"

type Connection struct {
	DB *gorm.DB
}
