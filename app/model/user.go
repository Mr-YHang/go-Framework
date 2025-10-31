package model

import "time"

type User struct {
	ID       int64     `json:"id" gorm:"column:id"`
	UserName string    `json:"user_name" gorm:"column:user_name"`
	Password string    `json:"password" gorm:"column:password"`
	Role     int       `json:"role" gorm:"column:role"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}

func (m *User) TableName() string {
	return "user"
}
