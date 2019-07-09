package model

import "github.com/ieshan/zen/conn"

func Migrate(db *conn.DBClient)  {
	db.AutoMigrate(User{})
}
