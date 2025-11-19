package api

import (
	"RedRock/core/utils/i18n"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("sso_jwt")

		if err != nil || token == "" {
			i18n.GetI18n(c).DoNotLogin()
			//c.Redirect(http.StatusFound, "/v1/blog/user/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
