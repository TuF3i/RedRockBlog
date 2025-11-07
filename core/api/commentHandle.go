package api

import (
	"RedRock/core"
	"RedRock/core/service/CommentManager"

	"github.com/gin-gonic/gin"
)

func (root *GinApi) AddCommentHandle(c *gin.Context) {
	if err := CommentManager.CreateManager(core.DataBase, c).AddComment(); err != nil {
		return
	}
}

func (root *GinApi) AddSubCommentHandle(c *gin.Context) {
	if err := CommentManager.CreateManager(core.DataBase, c).AddSubComment(); err != nil {
		return
	}
}

func (root *GinApi) DeleteCommentsHandle(c *gin.Context) {
	if err := CommentManager.CreateManager(core.DataBase, c).DeleteComments(); err != nil {
		return
	}
}

func (root *GinApi) GetCommentHandle(c *gin.Context) {
	if err := CommentManager.CreateManager(core.DataBase, c).GetComment(); err != nil {
		return
	}
}
