package mmiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/namejlt/gozen"
	"github.com/namejlt/gozenmodule/mcommon"
)

func DisabledRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := mcommon.CodeCommonRouteDisabled
		gozen.UtilResponseReturnJson(c, code, nil)
		c.Abort()
	}
}
