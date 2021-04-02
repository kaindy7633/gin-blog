package api

import (
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"gin-blog/service/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `form:"username" json:"username" valid:"Required;MaxSize(50)"`
	Password string `form:"password" json:"password" valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var form auth
	appG := app.Gin{C: c}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{Username: form.Username, Password: form.Password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

}
