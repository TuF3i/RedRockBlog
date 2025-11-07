package api

import (
	"RedRock/core"
	"RedRock/core/service/UserManager"
	"RedRock/core/utils/OAuth"
	"RedRock/core/utils/i18n"
	"RedRock/core/utils/jwt"
	"RedRock/core/utils/md5"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (root *GinApi) LogoutHandle(c *gin.Context) {
	c.SetCookie(
		"sso_jwt",
		"",
		-1,
		"/",
		core.GlobalConf.Domain,
		false,
		true,
	)
	c.Redirect(http.StatusFound, "/")
}

func (root *GinApi) LoginHandle(c *gin.Context) {
	o, err := OAuth.InitOAuth()

	if err != nil {
		i18n.GetI18n(c).ServerError(err)
		return
	}

	o.OAuth2LoginLogic(c)
}

func (root *GinApi) DeleteUserHandle(c *gin.Context) {
	session, _ := c.Cookie("sso_jwt")

	if session == "" {
		i18n.GetI18n(c).DoNotLogin()
		return
	}

	j, err := jwt.InitJWT()
	if err != nil {
		i18n.GetI18n(c).ServerError(err)
		return
	}

	userID, ok := j.RecoverData(session)
	if !ok {
		i18n.GetI18n(c).TokenNotSupport()
		return
	}

	err = UserManager.CreateManager(core.DataBase, c).DeleteUserByID(userID)
	if err != nil {
		return
	}

	//i18n.GetI18n(c).OperationSuccess()
}

func (root *GinApi) CallBackHandle(c *gin.Context) {
	o, err := OAuth.InitOAuth()
	if err != nil {
		i18n.GetI18n(c).ServerError(err)
		return
	}

	j, err := jwt.InitJWT()
	if err != nil {
		i18n.GetI18n(c).ServerError(err)
		return
	}

	err, userinfo := o.CallBackLogic(c)
	if err != nil {
		i18n.GetI18n(c).ServerError(err)
		return
	}

	token, err := j.GenJWT(md5.GenMD5(userinfo.ID))
	if err != nil {
		//fmt.Println("sign")
		i18n.GetI18n(c).ServerError(err)
		return
	}

	c.SetCookie(
		"sso_jwt",
		token,
		int(j.JwtExpiry.Seconds()),
		"/",
		core.GlobalConf.Domain,
		false,
		true,
	)

	err = UserManager.CreateManager(core.DataBase, c).AddUser(userinfo)

	if err.Error() == "user Exists" {
		c.Redirect(http.StatusFound, "/")
		return
	}

	if err != nil {
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func (root *GinApi) GetUserInfoHandle(c *gin.Context) {
	session, _ := c.Cookie("sso_jwt")

	if session == "" {
		i18n.GetI18n(c).DoNotLogin()
		return
	}

	j, err := jwt.InitJWT()
	if err != nil {
		i18n.GetI18n(c).ServerError(err)
		return
	}

	userID, ok := j.RecoverData(session)
	if !ok {
		i18n.GetI18n(c).TokenNotSupport()
		return
	}

	err = UserManager.CreateManager(core.DataBase, c).GetUserInfo(userID)
	if err != nil {
		return
	}
}
