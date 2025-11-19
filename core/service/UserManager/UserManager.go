package UserManager

import (
	"RedRock/core/dao/models"
	"RedRock/core/utils/OAuth"
	"RedRock/core/utils/i18n"
	"RedRock/core/utils/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* 初始化UserManager对象 - Start */
func CreateManager(db *gorm.DB, c *gin.Context) *Manager {
	//创建UserManager对象
	return &Manager{DB: db, c: c}
}

/* 初始化UserManager对象 - End */

/* AddUser方法 - Start */
func (root *Manager) AddUser(userInfo OAuth.UserInfo) error {
	//开启数据库事务
	tx := root.DB.Begin()

	//检查用户是否存在
	if row := tx.Where("id = ?", md5.GenMD5(userInfo.ID)).Find(&models.Users{}).RowsAffected; row > 0 {
		// i18n.GetI18n(root.c).UserExist(md5.GenMD5(userInfo.ID))
		// return fmt.Errorf("user Exists")
		tx.Rollback()
		return fmt.Errorf("user Exists")
	}

	//创建用户
	err := tx.Create(&models.Users{
		Name: userInfo.Name,
		ID:   md5.GenMD5(userInfo.ID),
	}).Error

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	i18n.GetI18n(root.c).OperationSuccess()
	return nil
}

/* AddUser方法 - End */

/* DeleteUserByID方法 - Start */
func (root *Manager) DeleteUserByID(userID string) error {
	//开启数据库事务
	tx := root.DB.Begin()

	//检查用户是否存在
	if row := tx.Where("id = ?", userID).Find(&models.Users{}).RowsAffected; row == 0 {
		tx.Rollback()
		i18n.GetI18n(root.c).UserNotExist(userID)
		return fmt.Errorf("user Not Exist")
	}

	//获取用户名下文章
	var articleID []uint
	err := tx.Model(&models.Article{}).Where("author_id = ?", userID).Pluck("article_id", &articleID).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	//获取用户文章下的评论
	var ArticleCommentID []uint
	err = tx.Model(&models.Comment{}).Where("article_id IN (?)", articleID).Pluck("id", &ArticleCommentID).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	//删除用户名下的文章
	if len(articleID) != 0 {
		err = tx.Where("author_id = ?", userID).Delete(&models.Article{}).Error
		if err != nil {
			tx.Rollback()
			i18n.GetI18n(root.c).ServerError(err)
			return err
		}
	}

	//删除用户文章下的评论
	if len(ArticleCommentID) != 0 {
		err = tx.Where("id IN (?)", ArticleCommentID).Delete(&models.Comment{}).Error
		if err != nil {
			tx.Rollback()
			i18n.GetI18n(root.c).ServerError(err)
			return err
		}
	}

	//删除用户自己的评论以及其子评论
	var userCommentsID []uint
	err = tx.Model(&models.Comment{}).Where("author_id = ?", userID).Pluck("id", &userCommentsID).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	//删除子评论
	for _, id := range userCommentsID {
		err = DeleteCommentAndChildren(tx, id)
		if err != nil {
			tx.Rollback()
			i18n.GetI18n(root.c).ServerError(err)
			return err
		}
	}

	//删除用户
	err = tx.Where("id = ?", userID).Unscoped().Delete(&models.Users{}).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	i18n.GetI18n(root.c).OperationSuccess()
	return nil
}

func getDescendantIDs(db *gorm.DB, targetID uint) ([]uint, error) {
	allIDs := []uint{targetID}     //所有评论(父评论 + 子评论)
	currentIDs := []uint{targetID} //要遍历的节点ID

	for {
		//查询当前遍历的节点下所有的子评论
		var childIDs []uint
		err := db.Model(&models.Comment{}).Where("parent_id IN (?)", currentIDs).Pluck("id", &childIDs).Error
		if err != nil {
			return nil, err
		}

		//没有子评论时跳出循环
		if len(childIDs) == 0 {
			break
		}

		//追加ID
		allIDs = append(allIDs, childIDs...)
		currentIDs = childIDs
	}

	return allIDs, nil
}

func DeleteCommentAndChildren(db *gorm.DB, targetID uint) error {
	// 开启事务，防止删到一半出错，数据不一致
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 找到所有要删除的ID
	ids, err := getDescendantIDs(tx, targetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 批量删除这些ID的评论
	if err := tx.Where("id IN (?)", ids).Delete(&models.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

/* DeleteUserByID方法 - End */

/* GetUserInfo方法 - Start */
func (root *Manager) GetUserInfo(sso string) error {
	var userInfo models.Users

	//开启数据库事务
	tx := root.DB.Begin()

	err := tx.Where("id = ?", sso).First(&userInfo).Error

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	root.c.JSON(http.StatusOK, userInfo)
	return nil
}

/* GetUserInfo方法 - End */
