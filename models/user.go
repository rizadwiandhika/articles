package models

import (
	"time"

	"gorm.io/gorm"
)

/* with gorm.Model there will be addition fields
type User struct {
	ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
} */

type User struct {
	// gorm.Model        // equivalent: ID, created_at, updated_at, deleted_at
	ID       uint   `json:"id" gorm:"primary_key"` // kalo primary key -> auto increment
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
