package util

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	// page, _ := com.StrTo(c.Query("page")).Int()
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
