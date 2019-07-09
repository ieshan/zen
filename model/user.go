package model

import (
	"github.com/ieshan/zen/conn"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string  `gorm:"type:varchar(100);unique_index"`
}
type UserList []User

func (user *User) Create(db *conn.DBClient) {
	db.Create(user)
}

func (user *User) Update(db *conn.DBClient) {
	db.Save(user)
}

func (user *User) Get(db *conn.DBClient, query interface{}, args ...interface{}) {
	db.Where(query, args).Find(user)
}

func (users UserList) Find(db *conn.DBClient, offset int, limit int, query interface{}, args ...interface{}) {
	db.Where(query, args).Offset(offset).Limit(limit).Find(&users)
}
