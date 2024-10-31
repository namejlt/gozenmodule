package mcommon

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/namejlt/gozen"
)

func ErrorResponse(c *gin.Context) {
	path := c.Request.URL.Path
	// 默认返回1005
	code := CodeCommonUserNoLogin
	if strings.HasPrefix(path, "/app") {
		// app返回1024
		code = CodeCommonUserNoLoginApp
	}
	gozen.UtilResponseReturnJson(c, code, nil)
}
