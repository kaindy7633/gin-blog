package v1

import (
	"fmt"
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"gin-blog/service/article_service"
	"gin-blog/service/tag_service"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取所有文章列表
func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	tagId := -1
	if arg := c.PostForm("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// 获取指定id的文章
func GetArticle(c *gin.Context) {
	// 获取id
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	// 验证
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

// 新建文章结构体
type AddArticleForm struct {
	TagID     int    `form:"tag_id" json:"tag_id" valid:"Required;Min(1)"`
	Title     string `form:"title" json:"title" valid:"Required;MaxSize(100)"`
	Desc      string `form:"desc" json:"desc" valid:"Required;MaxSize(255)"`
	Content   string `form:"content" json:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy string `form:"created_by" json:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" json:"state" valid:"Range(0,1)"`
}

// 新建文章
func AddArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddArticleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{ID: form.TagID}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	articleService := article_service.Article{
		TagID:     form.TagID,
		Title:     form.Title,
		Desc:      form.Desc,
		Content:   form.Content,
		State:     form.State,
		CreatedBy: form.CreatedBy,
	}

	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// 更新文章结构体
type EditArticleForm struct {
	ID         int    `form:"id" json:"id" valid:"Required;Min(1)"`
	TagID      int    `form:"tag_id" json:"tag_id" valid:"Required;Min(1)"`
	Title      string `form:"title" json:"title" valid:"Required;MaxSize(100)"`
	Desc       string `form:"desc" json:"desc" valid:"Required;MaxSize(255)"`
	Content    string `form:"content" json:"content" valid:"Required;MaxSize(65535)"`
	ModifiedBy string `form:"modified_by" json:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" json:"state" valid:"Range(0,1)"`
}

// 更新文章
func EditArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditArticleForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	fmt.Printf("%#v %#v\n", httpCode, errCode)
	fmt.Println(form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleService := article_service.Article{
		ID:         form.ID,
		TagID:      form.TagID,
		Title:      form.Title,
		Desc:       form.Desc,
		Content:    form.Content,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tagService := tag_service.Tag{ID: form.TagID}
	exists, err = tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	// 获取要删除的文章的id
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
