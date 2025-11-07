package UserManager

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Manager struct {
	DB *gorm.DB
	c  *gin.Context
}
