package dao

import (
	"gorm.io/gorm"
)

// User Dao
type User struct {
	DB *gorm.DB
}
