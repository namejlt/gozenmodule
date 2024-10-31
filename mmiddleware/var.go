package mmiddleware

import "github.com/namejlt/gozen"

var (
	authPassToken string // 密码验证密钥 秒

)

func init() {
	authPassToken = gozen.ConfigAppGetString("AuthPassToken", "")
}
