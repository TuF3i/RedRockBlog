package ArticleManager

import (
	"RedRock/core/dao/models"
	"RedRock/core/utils/dataConv"
	"RedRock/core/utils/i18n"
	"RedRock/core/utils/jwt"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateManager(c *gin.Context, db *gorm.DB) *ArticleManager {
	return &ArticleManager{c: c, db: db}
}

func (root *ArticleManager) AddWork() error {
	var Article ReceivedArticleJson

	//开启数据库事务
	tx := root.db.Begin()

	err := root.c.ShouldBindJSON(&Article)
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	j, err := jwt.InitJWT()
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	session, _ := root.c.Cookie("sso_jwt")
	UserID, ok := j.RecoverData(session)
	if !ok {
		tx.Rollback()
		i18n.GetI18n(root.c).TokenNotSupport()
		return fmt.Errorf("token Unuseable")
	}

	NewArticle := models.Article{
		AuthorID:     UserID,
		LanguageID:   Article.LanguageID,
		Title:        Article.Title,
		ExtTitle:     Article.ExtTitle,
		Introduction: Article.Introduction,
		Content:      Article.Content,
		IfDraft:      Article.IfDraft,
		IfPrivate:    Article.IfPrivate,
		IfTop:        Article.IfTop,
	}

	err = tx.Create(&NewArticle).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	i18n.GetI18n(root.c).OperationSuccess()
	return nil
}

func (root *ArticleManager) DeleteWork() error {
	//router.DELETE("/article/:id"

	//开启数据库事务
	tx := root.db.Begin()

	ArticleID := root.c.Param("id")

	j, err := jwt.InitJWT()
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	session, _ := root.c.Cookie("sso_jwt")
	AuthorID, ok := j.RecoverData(session)
	if !ok {
		tx.Rollback()
		i18n.GetI18n(root.c).TokenNotSupport()
		return fmt.Errorf("token Unuseable")
	}

	err = tx.Where("author_id = ? AND article_id = ?", AuthorID, dataConv.InitConv().Str2uint(ArticleID)).Delete(&models.Article{}).Error

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	err = tx.Where("article_id = ?", dataConv.InitConv().Str2uint(ArticleID)).Delete(&models.Comment{}).Error

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	i18n.GetI18n(root.c).OperationSuccess()
	return nil
}

func (root *ArticleManager) GetMyWorkList() error {
	//开启数据库事务
	tx := root.db.Begin()

	j, err := jwt.InitJWT()
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	session, _ := root.c.Cookie("sso_jwt")
	AuthorID, ok := j.RecoverData(session)
	if !ok {
		tx.Rollback()
		i18n.GetI18n(root.c).TokenNotSupport()
		return fmt.Errorf("token Unuseable")
	}

	var topWork []MyWorkList
	var drafts []MyWorkList
	var normalWork []MyWorkList

	myWork := tx.Model(&models.Article{}).
		Where("author_id = ?", AuthorID).
		Select("language_id, " +
			"article_id, " +
			"title, " +
			"ext_title, " +
			"introduction, " +
			"if_draft, " +
			"if_private, " +
			"if_top, " +
			"updated_at")

	err = myWork.Where("if_top = ? AND if_draft = ?", true, false).
		Order("updated_at desc").Find(&topWork).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	myWork = tx.Model(&models.Article{}).
		Where("author_id = ?", AuthorID).
		Select("language_id, " +
			"article_id, " +
			"title, " +
			"ext_title, " +
			"introduction, " +
			"if_draft, " +
			"if_private, " +
			"if_top, " +
			"updated_at")

	err = myWork.Where("if_draft = ?", true).
		Order("updated_at desc").Find(&drafts).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	myWork = tx.Model(&models.Article{}).
		Where("author_id = ?", AuthorID).
		Select("language_id, " +
			"article_id, " +
			"title, " +
			"ext_title, " +
			"introduction, " +
			"if_draft, " +
			"if_private, " +
			"if_top, " +
			"updated_at")

	err = myWork.Where("if_top = ? AND if_draft = ?", false, false).
		Order("updated_at desc").Find(&normalWork).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	data := ResponseAllMyWork{
		TopWork:    topWork,
		Drafts:     drafts,
		NormalWork: normalWork,
	}

	tx.Commit()
	root.c.JSON(http.StatusOK, data)

	return nil
}

func (root *ArticleManager) GetGlobalArticleList() error {
	var topWork []WorkList
	var normalWork []WorkList

	//开启数据库事务
	tx := root.db.Begin()

	Work := tx.Model(&models.Article{}).
		Where("if_draft = ?", false).
		Select("language_id, " +
			"article_id, " +
			"title, " +
			"ext_title, " +
			"introduction, " +
			"if_private, " +
			"if_top, " +
			"updated_at")

	err := Work.Where("if_top = ?", true).Order("updated_at desc").Find(&topWork).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	Work = tx.Model(&models.Article{}).
		Where("if_draft = ?", false).
		Select("language_id, " +
			"article_id, " +
			"title, " +
			"ext_title, " +
			"introduction, " +
			"if_private, " +
			"if_top, " +
			"updated_at")

	err = Work.Where("if_top = ?", false).Order("updated_at desc").Find(&normalWork).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	data := ResponseAllWork{
		TopWork:    topWork,
		NormalWork: normalWork,
	}

	tx.Commit()
	root.c.JSON(http.StatusOK, data)

	return nil
}

func (root *ArticleManager) SearchArticleByExtTitle() error {
	//router.GET("/article/:extName"
	var work WorkWithExtName

	//开启数据库事务
	tx := root.db.Begin()

	ArticleExtName := root.c.Param("extName")
	err := tx.Model(&models.Article{}).Where("ext_title = ?", ArticleExtName).Select(
		"language_id, " +
			"article_id, " +
			"title, " +
			"ext_title, " +
			"introduction, " +
			"if_private, " +
			"if_top").First(&work).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	data := struct {
		SoloWork []WorkWithExtName
	}{SoloWork: []WorkWithExtName{work}}

	tx.Commit()
	root.c.JSON(http.StatusOK, data)

	return nil
}

func (root *ArticleManager) SearchArticleMoHuly() error {
	//router.GET("/article/:name"
	var top []WorkWithMoHuName
	var normal []WorkWithMoHuName

	//开启数据库事务
	tx := root.db.Begin()

	name := root.c.Param("name")
	//work := root.db.Where("name LIKE ?", "%"+name+"%")
	Work := tx.Model(&models.Article{}).
		Where("title LIKE ?", "%"+name+"%").
		Select(
			"language_id, " +
				"article_id, " +
				"title, " +
				"ext_title, " +
				"introduction, " +
				"if_private, " +
				"if_top, " +
				"updated_at")

	err := Work.Where("if_top = ?", 1).
		Order("updated_at desc").First(&top).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	Work = tx.Model(&models.Article{}).
		Where("title LIKE ?", "%"+name+"%").
		Select(
			"language_id, " +
				"article_id, " +
				"title, " +
				"ext_title, " +
				"introduction, " +
				"if_private, " +
				"if_top, " +
				"updated_at")

	err = Work.Where("if_top = ?", 0).
		Order("updated_at desc").First(&normal).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	data := ResponseWorkWithMoHuName{
		TopWork:    top,
		NormalWork: normal,
	}

	tx.Commit()
	root.c.JSON(http.StatusOK, data)

	return nil
}

func (root *ArticleManager) UpdateWork() error {
	var Article UpdateWork

	//开启数据库事务
	tx := root.db.Begin()

	err := root.c.ShouldBindJSON(&Article)
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	j, err := jwt.InitJWT()
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	session, _ := root.c.Cookie("sso_jwt")
	UserID, ok := j.RecoverData(session)

	if !ok {
		tx.Rollback()
		i18n.GetI18n(root.c).TokenNotSupport()
		return fmt.Errorf("token Unuseable")
	}

	var checker models.Article

	err = tx.Where("article_id = ?", Article.ArticleID).First(&checker).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		i18n.GetI18n(root.c).WorkNotExist()
		return fmt.Errorf("work Dont Exist")
	}

	if checker.AuthorID != UserID {
		tx.Rollback()
		i18n.GetI18n(root.c).WorkNotBelongToYou()
		return fmt.Errorf("this Work Dont Belong To You")
	}

	NewArticle := models.Article{
		AuthorID:     UserID,
		LanguageID:   Article.LanguageID,
		Title:        Article.Title,
		ExtTitle:     Article.ExtTitle,
		Introduction: Article.Introduction,
		Content:      Article.Content,
		IfDraft:      Article.IfDraft,
		IfPrivate:    Article.IfPrivate,
		IfTop:        Article.IfTop,
	}

	err = tx.Model(&models.Article{}).Where("article_id = ?", Article.ArticleID).Omit("created_at").Updates(&NewArticle).Error
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	i18n.GetI18n(root.c).OperationSuccess()
	return nil
}

func (root *ArticleManager) GetWorkDetail() error {
	//router.DELETE("/article/:id"
	var Work models.Article

	//开启数据库事务
	tx := root.db.Begin()

	ArticleID := root.c.Param("id")

	j, err := jwt.InitJWT()
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	session, _ := root.c.Cookie("sso_jwt")
	AuthorID, ok := j.RecoverData(session)
	if !ok {
		tx.Rollback()
		i18n.GetI18n(root.c).TokenNotSupport()
		return fmt.Errorf("token Unuseable")
	}

	err = tx.Where("author_id = ? AND article_id = ?",
		AuthorID, dataConv.InitConv().Str2uint(ArticleID)).First(&Work).Error

	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	tx.Commit()
	root.c.JSON(http.StatusOK, Work)

	return nil
}

func (root *ArticleManager) GetArticle() error {
	var Article models.Article

	//开启数据库事务
	tx := root.db.Begin()

	//router.GET("/article/:id"
	ArticleID := root.c.Param("id")

	j, err := jwt.InitJWT()
	if err != nil {
		tx.Rollback()
		i18n.GetI18n(root.c).ServerError(err)
		return err
	}

	session, _ := root.c.Cookie("sso_jwt")
	if session != "" {
		AuthorID, ok := j.RecoverData(session)
		if !ok {
			tx.Rollback()
			i18n.GetI18n(root.c).TokenNotSupport()
			return fmt.Errorf("token Unuseable")
		}

		err := tx.Where("article_id = ?", dataConv.InitConv().Str2uint(ArticleID)).First(&Article).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			i18n.GetI18n(root.c).WorkNotExist()
			return fmt.Errorf("artical Not Exist")
		}

		if Article.IfPrivate == false {
			tx.Commit()
			root.c.JSON(http.StatusOK, Article)
			return nil
		}

		if Article.IfPrivate == true && Article.AuthorID == AuthorID {
			tx.Commit()
			root.c.JSON(http.StatusOK, Article)
			return nil
		} else {
			tx.Rollback()
			i18n.GetI18n(root.c).YouAreNotAuthorized()
			return fmt.Errorf("you are not Authorized")
		}
	}

	err = tx.Where("article_id = ?", dataConv.InitConv().Str2uint(ArticleID)).First(&Article).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		i18n.GetI18n(root.c).WorkNotExist()
		return fmt.Errorf("artical Not Exist")
	}

	if Article.IfPrivate == true {
		tx.Rollback()
		i18n.GetI18n(root.c).YouAreNotAuthorized()
		return fmt.Errorf("you are not Authorized")
	}

	tx.Commit()
	root.c.JSON(http.StatusOK, Article)
	return nil
}
