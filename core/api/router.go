package api

import (
	"github.com/gin-gonic/gin"
)

func InitGin() *GinApi {
	return &GinApi{r: gin.Default()}
}

func (root *GinApi) InitGinApi() *gin.Engine {

	root.r.GET("/callback", root.CallBackHandle)

	root.r.GET("/", root.RedirectToRedRockPage)

	root.r.LoadHTMLGlob("data/redrock/*")

	blogRoot := root.r.Group("/v1/blog")

	blogRoot.GET("/red-rock-page", root.RedRockHandle)

	articleSubRoot := blogRoot.Group("/article")
	commentSubRoot := blogRoot.Group("/comment")
	userSubRoot := blogRoot.Group("/user")

	userSubRoot.GET("/login", root.LoginHandle)
	userSubRoot.GET("/logout", AuthMiddleware(), root.LogoutHandle)
	userSubRoot.GET("/info", AuthMiddleware(), root.GetUserInfoHandle)
	userSubRoot.DELETE("/delete", AuthMiddleware(), root.DeleteUserHandle)

	commentSubRoot.POST("/add", AuthMiddleware(), root.AddCommentHandle)
	commentSubRoot.POST("/add-sub", AuthMiddleware(), root.AddSubCommentHandle)
	commentSubRoot.DELETE("/delete/:id", AuthMiddleware(), root.DeleteCommentsHandle)
	commentSubRoot.GET("/get/:id", root.GetCommentHandle)

	articleSubRoot.POST("/add", AuthMiddleware(), root.AddWorkHandle)                       //ok
	articleSubRoot.DELETE("/delete/:id", AuthMiddleware(), root.DeleteWorkHandle)           //ok
	articleSubRoot.GET("/my-work-list", AuthMiddleware(), root.GetMyWorkListHandle)         //ok
	articleSubRoot.GET("/articles", root.GetGlobalArticleListHandle)                        //ok
	articleSubRoot.GET("/search-ext-name/:extName", root.SearchArticleByExtTitleHandle)     //ok
	articleSubRoot.GET("/search-mohu-name/:name", root.SearchArticleMoHulyHandle)           //ok
	articleSubRoot.PUT("/update", AuthMiddleware(), root.UpdateWorkHandle)                  //ok
	articleSubRoot.GET("/get-work-content/:id", AuthMiddleware(), root.GetWorkDetailHandle) //ok
	articleSubRoot.GET("/get-article-content/:id", root.GetArticleHandle)

	return root.r
}
