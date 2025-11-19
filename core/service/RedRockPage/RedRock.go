package RedRockPage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedRock(c *gin.Context) *RedRockPage {
	return &RedRockPage{c: c}
}

func (root *RedRockPage) ShowPage() {
	root.c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "RedRock",
	})
}
