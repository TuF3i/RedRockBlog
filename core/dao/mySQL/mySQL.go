package mySQL

import (
	"RedRock/core"
	"RedRock/core/dao/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *MySQLDataBase {
	return &MySQLDataBase{}
}

func (root *MySQLDataBase) GetConnection() error {
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		core.GlobalConf.MySQLUser,
		core.GlobalConf.MySQLPassword,
		core.GlobalConf.MySQLAddr,
		core.GlobalConf.MySQLPort,
		core.GlobalConf.MySQLDBName)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect to MySQL Fail - %v", err.Error())
	}

	root.DB = db
	return nil
}

func (root *MySQLDataBase) MigrateDataBase() error {
	err := root.DB.AutoMigrate(
		&models.Users{},
		&models.Article{},
		&models.Comment{},
	)

	if err != nil {
		return fmt.Errorf("migrate DataBase Fali - %v", err.Error())
	}

	return nil
}
