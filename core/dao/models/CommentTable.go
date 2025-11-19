package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Level         uint       //评论层级：0表示顶级评论，非0表示是这个ID的子评论
	ArticleID     uint       // 文章ID
	Content       string     // 评论内容
	IP            string     // 评论者的IP
	Location      string     // 评论者的地理位置
	Author        string     // 谁发的
	AuthorID      string     // 评论作者ID(标识符)(MD5值)
	ParentID      uint       // 父评论ID
	ChildComments []*Comment `json:"childComments" gorm:"-"`
}
