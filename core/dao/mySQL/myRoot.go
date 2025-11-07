package mySQL

import "gorm.io/gorm"

type MySQLDataBase struct {
	DB *gorm.DB
}
