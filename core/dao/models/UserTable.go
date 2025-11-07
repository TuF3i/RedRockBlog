package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	//gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	
	Name string //GitHub用户名(明文)
	ID   string `gorm:"type:varchar(255);uniqueIndex"` //GitHubID(唯一标识符)(MD5值)
}
