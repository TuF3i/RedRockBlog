package api

import (
	"RedRock/core/service/RedRockPage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (root *GinApi) RedRockHandle(c *gin.Context) {
	RedRockPage.RedRock(c).ShowPage()
}

func (root *GinApi) RedirectToRedRockPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/v1/blog/red-rock-page")
}
