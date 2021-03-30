package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DatabaseConnection struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

func New(connection DatabaseConnection) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", connection.User, connection.Password, connection.Database))
	return db, err
}
