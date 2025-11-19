package CommentManager

import (
	"RedRock/core"
	"RedRock/core/dao/models"
	"RedRock/core/utils/dataConv"
	"RedRock/core/utils/i18n"
	"RedRock/core/utils/ip2Location"
	"RedRock/core/utils/jwt"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateManager(db *gorm.DB, c *gin.Context) *Manager {
	return &Manager{db: db, c: c}
}

func (root *Manager) AddComment() error {
	var rootComment RootComment
	//开启数据库事务
	tx := root.db.Begin()

	//获取AuthorID
	token, _ := root.c.Cookie("sso_jwt")
	j, err := jwt.InitJWT()

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	AuthorID, ok := j.RecoverData(token)
	if !ok {
		tx.Rollback()
		i18n.GetI18n(root.c).TokenNotSupport()
		return fmt.Errorf("JWT Error - Unuseable Token")
	}

	//获取请求者的IP
	ip := root.c.ClientIP()
	location, err := ip2Location.IP2Location(ip).GetLocation()

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	//获取POST请求的JSON数据
	err = root.c.ShouldBindJSON(&rootComment)
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	var uname []string
	err = tx.Model(&models.Users{}).Where("id = ?", AuthorID).Pluck("name", &uname).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	data := models.Comment{
		AuthorID:  AuthorID,
		Level:     0,                     // 父级评论Level为0
		ArticleID: rootComment.ArticleID, // 评论对应的文章ID
		Content:   rootComment.Content,   // 评论内容
		IP:        ip,
		Location:  location,
		Author:    uname[0], //评论作者
		ParentID:  0,        // 父级评论
	}

	//创建父评论
	err = tx.Create(&data).Error

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	i18n.GetI18n(root.c).OperationSuccess()
	return nil
}

func (root *Manager) AddSubComment() error {
	var childComment ChildComment

	//开启数据库事务
	tx := root.db.Begin()

	token, _ := root.c.Cookie("sso_jwt")
	j, err := jwt.InitJWT()

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	AuthorID, ok := j.RecoverData(token)
	if !ok {
		tx.Rollback()
		i18n.GetI18n(root.c).TokenNotSupport()
		return fmt.Errorf("JWT Error - Unuseable Token")
	}

	//获取请求者的IP
	ip := root.c.ClientIP()
	location, err := ip2Location.IP2Location(ip).GetLocation()

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	//获取POST请求的JSON数据
	err = root.c.ShouldBindJSON(&childComment)
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	var uname []string
	err = tx.Model(&models.Users{}).Where("id = ?", AuthorID).Pluck("name", &uname).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	var fatherLevel uint

	//获取父评论
	err = tx.Model(&models.Comment{}).Where("id = ?", childComment.ParentID).Pluck("level", &fatherLevel).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		i18n.GetI18n(root.c).CanNotFindFatherComment()
		return fmt.Errorf("comment Error - Can`t find Father Comment")
	}

	if fatherLevel >= core.GlobalConf.MaxCommentLevel {
		tx.Rollback()
		i18n.GetI18n(root.c).ReachMaxLevel()
		return fmt.Errorf("comment Reach The Maxium Level")
	}

	data := models.Comment{
		Level:     fatherLevel + 1,
		ArticleID: childComment.ArticleID,
		Content:   childComment.Content,
		IP:        ip,
		Location:  location,
		Author:    uname[0],
		AuthorID:  AuthorID,
		ParentID:  childComment.ParentID, // 父级评论
	}

	err = tx.Create(&data).Error

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	i18n.GetI18n(root.c).OperationSuccess()
	return nil
}

func (root *Manager) GetComment() error {
	//router.GET("/article/:id"
	var AllComments []*models.Comment
	commentMap := make(map[uint]*models.Comment)

	//开启数据库事务
	tx := root.db.Begin()

	ArticleID := dataConv.InitConv().Str2uint(root.c.Param("id"))
	err := tx.Where("article_id = ?", ArticleID).Order("created_at DESC").Find(&AllComments).Error

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	for _, comment := range AllComments {
		commentMap[comment.ID] = comment
	}

	AllComments = []*models.Comment{}

	for _, comment := range commentMap {
		if comment.ParentID == 0 {
			AllComments = append(AllComments, comment)
			continue
		}

		commentMap[comment.ParentID].ChildComments = append(commentMap[comment.ParentID].ChildComments, comment)

	}

	tx.Commit()
	root.c.JSON(http.StatusOK, gin.H{"comments": AllComments})
	return nil
}

func (root *Manager) DeleteComments() error {
	//获取评论ID
	//router.GET("/article/:id"

	CommentID := dataConv.InitConv().Str2uint(root.c.Param("id"))

	//获取AuthorID
	token, _ := root.c.Cookie("sso_jwt")
	j, err := jwt.InitJWT()

	if err != nil {
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	AuthorID, ok := j.RecoverData(token)
	if !ok {
		i18n.GetI18n(root.c).TokenNotSupport()
		return fmt.Errorf("token not Support")
	}

	if row := root.db.Where("id = ?", CommentID).Find(&models.Comment{}).RowsAffected; row == 0 {
		i18n.GetI18n(root.c).CommentNotExist()
		return fmt.Errorf("comment Not Exist")
	}

	var thisComment models.Comment
	err = root.db.Where("id = ?", CommentID).First(&thisComment).Error
	if err != nil {
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	if thisComment.AuthorID != AuthorID {
		i18n.GetI18n(root.c).CommentNotBelongToYou()
		return fmt.Errorf("this do not belong to you")
	}

	err = DeleteCommentAndChildren(root.db, CommentID)
	if err != nil {
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}
	
	i18n.GetI18n(root.c).OperationSuccess()
	return nil
}

func getDescendantIDs(db *gorm.DB, targetID uint) ([]uint, error) {
	allIDs := []uint{targetID}
	currentIDs := []uint{targetID}

	for {
		var childIDs []uint
		err := db.Model(&models.Comment{}).Where("parent_id IN (?)", currentIDs).Pluck("id", &childIDs).Error
		if err != nil {
			return nil, err
		}

		if len(childIDs) == 0 {
			break
		}

		allIDs = append(allIDs, childIDs...)
		currentIDs = childIDs
	}

	return allIDs, nil
}

func DeleteCommentAndChildren(db *gorm.DB, targetID uint) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	ids, err := getDescendantIDs(tx, targetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id IN (?)", ids).Delete(&models.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
