package i18n

import (
	"RedRock/core"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetI18n(c *gin.Context) *MyI18n {
	acceptLanguage := c.GetHeader("Accept-Language")
	bundle := core.Bundle
	ip := c.ClientIP()
	return &MyI18n{i18n: i18n.NewLocalizer(bundle, acceptLanguage), ipAddr: ip, c: c}
}

func (root *MyI18n) HelloWorld(userName string) (string, error) {
	message, err := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "hello",
		TemplateData: map[string]string{
			"Name": "User",
		},
	})
	return message, err
}

func (root *MyI18n) UserExist(userName string) {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "add_user_error_user_exist",
		TemplateData: map[string]string{
			"Name": userName,
		},
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) UserNotExist(userName string) {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "add_user_error_user_not_exist",
		TemplateData: map[string]string{
			"Name": userName,
		},
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusNotFound, gin.H{"Error": message})
}

func (root *MyI18n) ServerError(err error) {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "server_error",
		TemplateData: map[string]string{
			"ErrorMsg": err.Error(),
		},
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) DoNotLogin() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "not_login",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) TokenNotSupport() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "token_not_support",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) AddCommentError(err error) {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "add_comment_error",
		TemplateData: map[string]string{
			"ErrorMsg": err.Error(),
		},
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) ReachMaxLevel() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "reach_max_level",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) CommentNotExist() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "comment_not_exist",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) CommentNotBelongToYou() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "comment_not_belong_you",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) CanNotFindFatherComment() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "can_not_find_father_comment",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) WorkNotExist() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "work_not_exist",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) WorkNotBelongToYou() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "work_not_belong_you",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusServerError, gin.H{"Error": message})
}

func (root *MyI18n) YouAreNotAuthorized() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "you_are_not_authorized",
	})

	core.Logger.MyError(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusUnauthorized, gin.H{"Error": message})
}

func (root *MyI18n) OperationSuccess() {
	message, _ := root.i18n.Localize(&i18n.LocalizeConfig{
		MessageID: "operation_success",
	})

	core.Logger.MyDEBUG(fmt.Sprintf("%v - %v", root.ipAddr, message))
	root.c.JSON(core.StatusOK, gin.H{"Message": message})
}
