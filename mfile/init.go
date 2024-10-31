package mfile

import "github.com/namejlt/gozen"

/**

营销云专用


每一个客户环境申请一个新的bucket


仅用于内部上传下载，没有流量加速

*/

const (
	fileServiceModeTencent     = "tencent"
	fileServiceModeAliyun      = "aliyun"
	fileServiceModeHuaweicloud = "huaweicloud"
)

var (
	fileServiceMode string // 文件服务模式  tencent  aliyun huaweicloud  支持三种
)

func init() {
	fileServiceMode = gozen.ConfigAppGetString("FileServiceMode", "tencent")

	registerMode(fileServiceMode)
}
