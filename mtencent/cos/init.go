package cos

import (
	"github.com/namejlt/gozen"
)

/**

营销云专用


每一个客户环境申请一个新的bucket


仅用于内部上传下载，没有流量加速

*/

var (
	tencentCloudCosSecretID    string
	tencentCloudCosSecretKey   string
	tencentCloudCosBucketUrl   string
	tencentCloudCosServiceUrl  string
	tencentCloudCosCustomerKey string
)

func init() {
	tencentCloudCosSecretID = gozen.ConfigAppGetString("TencentCloudCosSecretID", "")
	tencentCloudCosSecretKey = gozen.ConfigAppGetString("TencentCloudCosSecretKey", "")
	tencentCloudCosBucketUrl = gozen.ConfigAppGetString("TencentCloudCosBucketUrl", "")
	tencentCloudCosServiceUrl = gozen.ConfigAppGetString("TencentCloudCosServiceUrl", "")
	tencentCloudCosCustomerKey = gozen.ConfigAppGetString("TencentCloudCosCustomerKey", "")
}
