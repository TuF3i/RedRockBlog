package ArticleManager

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleManager struct {
	c  *gin.Context
	db *gorm.DB
}

/* 功能块(AddWork) - Strat */

type ReceivedArticleJson struct {
	/* 标识段 */
	//ArticleID  uint   `json:"articleID"`  //文章ID(用来索引评论)
	//AuthorID   string `json:"authorID"`   //GitHubID(标识文章所有者)(MD5值)
	LanguageID uint `json:"languageID"` //文章语言ID,目前仅支持中文和英文(0:中文, 1:英文)(我太菜了, i18n不会用，用这个做替代方案)
	//UpdatedAt  time.Time `json:"updatedAt"`

	/* 数据段 */
	Title        string `json:"title"`        //文章标题
	ExtTitle     string `json:"extTitle"`     //别名
	Introduction string `json:"introduction"` //文章简介S
	Content      string `json:"content"`      //正文内容(Base64)(发Base64!!!)

	/* 属性段 */
	IfDraft   bool `json:"ifDraft"`   //是否为草稿(Ture: 是， False: 否)
	IfPrivate bool `json:"ifPrivate"` //是否私有(Ture: 是， False: 否)
	IfTop     bool `json:"ifTop"`     //是否置顶(Ture: 是， False: 否)
}

/* 功能块(AddWork) - End */

/* 功能块(GetMyWorkList) - Start */

type MyWorkList struct {
	/* 标识段 */
	LanguageID uint      `json:"languageID"`
	ArticleID  uint      `json:"articleID"`
	UpdatedAt  time.Time `json:"updatedAt"`

	/* 数据段 */
	Title        string `json:"title"`
	ExtTitle     string `json:"extTitle"`
	Introduction string `json:"introduction"`

	/* 标识段 */
	IfDraft   bool `json:"ifDraft"`
	IfPrivate bool `json:"ifPrivate"`
	IfTop     bool `json:"ifTop"`
}

type ResponseAllMyWork struct {
	TopWork    []MyWorkList `json:"topWork"`
	Drafts     []MyWorkList `json:"drafts"`
	NormalWork []MyWorkList `json:"normalWork"`
}

/* 功能块(GetMyWorkList) - End */

/* 功能块(GetGlobalArticleList) - Start */

type WorkList struct {
	/* 标识段 */
	LanguageID uint      `json:"languageID"`
	ArticleID  uint      `json:"articleID"`
	UpdatedAt  time.Time `json:"updatedAt"`

	/* 数据段 */
	Title        string `json:"title"`
	ExtTitle     string `json:"extTitle"`
	Introduction string `json:"introduction"`

	/* 标识段 */
	IfPrivate bool `json:"ifPrivate"`
	IfTop     bool `json:"ifTop"`
}

type ResponseAllWork struct {
	TopWork    []WorkList `json:"topWork"`
	NormalWork []WorkList `json:"normalWork"`
}

/* 功能块(GetGlobalArticleList) - End */

/* 功能块(SearchArticleByExtTitle) - Start */

type WorkWithExtName struct {
	/* 标识段 */
	LanguageID uint      `json:"languageID"`
	ArticleID  uint      `json:"articleID"`
	UpdatedAt  time.Time `json:"updatedAt"`

	/* 数据段 */
	Title        string `json:"title"`
	ExtTitle     string `json:"extTitle"`
	Introduction string `json:"introduction"`

	/* 标识段 */
	IfPrivate bool `json:"ifPrivate"`
	IfTop     bool `json:"ifTop"`
}

/* 功能块(SearchArticleByExtTitle) - End */

/* 功能块(SearchArticleMoHuly) - Start */

type WorkWithMoHuName struct {
	/* 标识段 */
	LanguageID uint      `json:"languageID"`
	ArticleID  uint      `json:"articleID"`
	UpdatedAt  time.Time `json:"updatedAt"`

	/* 数据段 */
	Title        string `json:"title"`
	ExtTitle     string `json:"extTitle"`
	Introduction string `json:"introduction"`

	/* 标识段 */
	IfPrivate bool `json:"ifPrivate"`
	IfTop     bool `json:"ifTop"`
}

type ResponseWorkWithMoHuName struct {
	TopWork    []WorkWithMoHuName `json:"topWork"`
	NormalWork []WorkWithMoHuName `json:"normalWork"`
}

/* 功能块(SearchArticleMoHuly) - End */

/* 功能块(UpdateWork) - Start */

type UpdateWork struct {
	/* 标识段 */
	ArticleID uint `json:"articleID"` //文章ID(用来索引评论)
	//AuthorID   string `json:"authorID"`   //GitHubID(标识文章所有者)(MD5值)
	LanguageID uint `json:"languageID"` //文章语言ID,目前仅支持中文和英文(0:中文, 1:英文)(我太菜了, i18n不会用，用这个做替代方案)
	//UpdatedAt  time.Time `json:"updatedAt"`

	/* 数据段 */
	Title        string `json:"title"`        //文章标题
	ExtTitle     string `json:"extTitle"`     //别名
	Introduction string `json:"introduction"` //文章简介S
	Content      string `json:"content"`      //正文内容(Base64)(发Base64!!!)

	/* 属性段 */
	IfDraft   bool `json:"ifDraft"`   //是否为草稿(Ture: 是， False: 否)
	IfPrivate bool `json:"ifPrivate"` //是否私有(Ture: 是， False: 否)
	IfTop     bool `json:"ifTop"`     //是否置顶(Ture: 是， False: 否)
}

/* 功能块(UpdateWork) - End */

/* 功能块(GetWorkDetail) - Start */

type GetWorkDetail struct {
	/* 标识段 */
	ArticleID  uint   `json:"articleID"`  //文章ID(用来索引评论)
	AuthorID   string `json:"authorID"`   //GitHubID(标识文章所有者)(MD5值)
	LanguageID uint   `json:"languageID"` //文章语言ID,目前仅支持中文和英文(0:中文, 1:英文)(我太菜了, i18n不会用，用这个做替代方案)
	//UpdatedAt  time.Time `json:"updatedAt"`

	/* 数据段 */
	Title        string `json:"title"`        //文章标题
	ExtTitle     string `json:"extTitle"`     //别名
	Introduction string `json:"introduction"` //文章简介S
	Content      string `json:"content"`      //正文内容(Base64)(发Base64!!!)

	/* 属性段 */
	IfDraft   bool `json:"ifDraft"`   //是否为草稿(Ture: 是， False: 否)
	IfPrivate bool `json:"ifPrivate"` //是否私有(Ture: 是， False: 否)
	IfTop     bool `json:"ifTop"`     //是否置顶(Ture: 是， False: 否)
}

/* 功能块(GetWorkDetail) - End */
