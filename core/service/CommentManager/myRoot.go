package CommentManager

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Manager struct {
	db *gorm.DB
	c  *gin.Context
}

type ChildComment struct {
	ArticleID uint   `json:"articleID"` // 文章ID
	Content   string `json:"content"`   // 评论内容
	//Author    string `json:"author"`    // 谁发的
	ParentID uint `json:"parentID"` // 父评论ID
}

type RootComment struct {
	ArticleID uint   `json:"articleID"` // 文章ID
	Content   string `json:"content"`   // 评论内容
	//Author    string `json:"author"`    // 谁发的
}

type Comment struct {
	Level         uint       `json:"level"`                  //评论层级：0表示顶级评论，非0表示是这个ID的子评论
	ArticleID     uint       `json:"articleID"`              // 文章ID
	Content       string     `json:"content"`                // 评论内容
	Author        string     `json:"author"`                 // 谁发的
	ParentID      uint       `json:"parentID"`               // 父评论ID：0表示顶级评论，非0表示是这个ID的子评论
	ChildComments []*Comment `json:"childComments" gorm:"-"` // 子评论列表(***数据库不存，仅内存中用***)
}
