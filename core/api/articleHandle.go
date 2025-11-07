package api

import (
	"RedRock/core"
	"RedRock/core/service/ArticleManager"

	"github.com/gin-gonic/gin"
)

func (root *GinApi) AddWorkHandle(c *gin.Context) {
	// 添加自己的文章（草稿和文章）
	if err := ArticleManager.CreateManager(c, core.DataBase).AddWork(); err != nil {
		return
	}
}

func (root *GinApi) DeleteWorkHandle(c *gin.Context) {
	//删除自己的文章（草稿和文章）
	if err := ArticleManager.CreateManager(c, core.DataBase).DeleteWork(); err != nil {
		return
	}
}

func (root *GinApi) GetMyWorkListHandle(c *gin.Context) {
	//获取自己的文章（草稿和文章）
	if err := ArticleManager.CreateManager(c, core.DataBase).GetMyWorkList(); err != nil {
		return
	}
}

func (root *GinApi) GetGlobalArticleListHandle(c *gin.Context) {
	//获取所有文章（只显示文章）
	if err := ArticleManager.CreateManager(c, core.DataBase).GetGlobalArticleList(); err != nil {
		return
	}
}

func (root *GinApi) SearchArticleByExtTitleHandle(c *gin.Context) {
	//别名索引
	if err := ArticleManager.CreateManager(c, core.DataBase).SearchArticleByExtTitle(); err != nil {
		return
	}
}

func (root *GinApi) SearchArticleMoHulyHandle(c *gin.Context) {
	//按名称模糊搜索
	if err := ArticleManager.CreateManager(c, core.DataBase).SearchArticleMoHuly(); err != nil {
		return
	}
}

func (root *GinApi) UpdateWorkHandle(c *gin.Context) {
	//更新自己的作品（可以更新作品和草稿）
	if err := ArticleManager.CreateManager(c, core.DataBase).UpdateWork(); err != nil {
		return
	}
}

func (root *GinApi) GetWorkDetailHandle(c *gin.Context) {
	//获取自己的作品内容，以用来修改和提交（只能看自己的）
	if err := ArticleManager.CreateManager(c, core.DataBase).GetWorkDetail(); err != nil {
		return
	}
}

func (root *GinApi) GetArticleHandle(c *gin.Context) {
	//获取用户文章的内容（可以看自己的以及别人公开的）
	if err := ArticleManager.CreateManager(c, core.DataBase).GetArticle(); err != nil {
		return
	}
}
