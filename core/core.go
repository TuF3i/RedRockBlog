package core

import (
	"RedRock/core/utils/config"
	llog "RedRock/core/utils/log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

var (
	GlobalConf config.Config
	Logger     *llog.Log
	DataBase   *gorm.DB
	Bundle     *i18n.Bundle
)

var (
	StatusOK           = 200
	StatusUnauthorized = 401
	StatusNotFound     = 404
	StatusServerError  = 500
)
