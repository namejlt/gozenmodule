package mmiddleware

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/namejlt/gozen"
	"github.com/namejlt/gozenmodule/mcommon"
	"strings"
	"time"
)

func AuthPass() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 通过header中获取密码，验证密码准确性

		pass := getAuthPassword()

		if c.GetHeader(mcommon.ToolApiAuthPassword) == pass {
			c.Next()
			return
		} else {
			gozen.UtilResponseReturnJsonNoP(c, mcommon.CodeCommonAccessFail, nil)
			c.Abort()
			return
		}
	}
}

// getAuthPassword 根据时间和密钥获取动态密码
func getAuthPassword() (str string) {
	origin := authPassToken + getCommonDateYdm_(time.Now())
	return utilCryptoMd5Lower(origin)
}

func utilCryptoMd5Lower(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return strings.ToLower(hex.EncodeToString(cipherStr))
}

func getCommonDateYdm_(t time.Time) string {
	return t.Format(mcommon.TIME_FORMAT_Y_M_D_)
}
