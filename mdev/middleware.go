package mdev

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/namejlt/gozen"
	"github.com/namejlt/gozenmodule/mcommon"
)

/**

开发调试开关

1、一些接口默认无法访问，可以开启开发者权限验证中间件，通过验证后访问，默认开启权限验证
2、一些操作，需要开启开发者展示，会展示多余信息，默认不展示
3、一些操作，需要开启开发者通过，会直接通过进行后续操作，默认不通过


*/

func DevAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if checkDevAuth(c) {
			c.Next()
		} else {
			mcommon.ErrorResponse(c)
			c.Abort()
		}
	}
}

func CheckConfigDevShow() bool {
	return isEnableDevShow()
}

func CheckConfigDevPass() bool {
	return isEnableDevPass()
}

func checkDevAuth(c *gin.Context) bool {
	if isEnableDevAuth() || gozen.ConfigEnvIsDev() || gozen.ConfigEnvIsBeta() {
		return true
	}
	inputKey, _ := c.Cookie(mcommon.DevAuthKey)

	realKey := getDevAuthSalt() + "saas" + time.Now().Format("2006-01-02") // saas加当天日期：saas2019-02-26
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(realKey))
	if inputKey == fmt.Sprintf("%x", md5Ctx.Sum(nil)) {
		return true
	}
	return false
}

func isEnableDevAuth() bool {
	return gozen.ConfigAppGetString("EnableDevAuth", "true") == "true" //默认开启
}

func isEnableDevShow() bool {
	return gozen.ConfigAppGetString("EnableDevShow", "false") == "true" //默认关闭
}

func isEnableDevPass() bool {
	return gozen.ConfigAppGetString("EnableDevPass", "false") == "true" //默认关闭
}

func getDevAuthSalt() string {
	return gozen.ConfigAppGetString("DevAuthSalt", "")
}
