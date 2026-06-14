package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	Email    string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string    `json:"password"`
	RoleID   uint      `json:"role_id"`
	Role     Role      `gorm:"foreignKey:RoleID"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
