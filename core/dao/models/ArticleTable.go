package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	/* 标识段 */
	ArticleID  uint   `gorm:"autoIncrement;uniqueIndex"` //文章ID(用来索引评论)
	AuthorID   string `gorm:"type:varchar(255)"`         //GitHubID(标识文章所有者)(MD5值)
	LanguageID uint   //文章语言ID,目前仅支持中文和英文(0:中文, 1:英文)

	/* 数据段 */
	Title        string //文章标题
	ExtTitle     string `gorm:"type:varchar(255);uniqueIndex"` //别名
	Introduction string //文章简介S
	Content      string //正文内容(Base64)

	/* 属性段 */
	IfDraft   bool //是否为草稿(Ture: 是， False: 否)
	IfPrivate bool //是否私有(Ture: 是， False: 否)
	IfTop     bool //是否置顶(Ture: 是， False: 否)
}
