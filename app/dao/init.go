package dao

import "go-Framework/global"

type Dao struct {
	User *User
}

func NewDao() *Dao {
	return &Dao{
		User: NewUser(global.App.DB, global.App.Redis),
	}
}
