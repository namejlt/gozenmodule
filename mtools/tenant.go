package mtools

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namejlt/gozenmodule/mcommon"
)

func GetTenantId(c *gin.Context) (tenantId uint64) {
	var param string
	if c.Request.Method == "GET" || c.Request.Method == "DELETE" {
		param = c.Query(mcommon.TenantIdKey)
	} else {
		param = c.PostForm(mcommon.TenantIdKey)
	}
	if param == "" {
		param, _ = c.Cookie(mcommon.TenantIdKey)
	}
	pInt, _ := strconv.Atoi(param)
	return uint64(pInt)
}
