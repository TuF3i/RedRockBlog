package i18n

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type MyI18n struct {
	i18n   *i18n.Localizer
	ipAddr string
	c      *gin.Context
}
