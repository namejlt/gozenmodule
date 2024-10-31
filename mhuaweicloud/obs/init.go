package obs

import (
	"fmt"
	hobs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/namejlt/gozen"
	"os"
)

/**

营销云专用


每一个客户环境申请一个新的bucket


仅用于内部上传下载，没有流量加速

*/

var (
	huaweicloudObsAccessKeyID     string
	huaweicloudObsSecretAccessKey string
	huaweicloudObsEndpoint        string
	huaweicloudObsBucket          string
	huaweicloudObsServiceUrl      string
	huaweicloudObsCustomerKey     string

	obsClient *hobs.ObsClient
)

func init() {
	huaweicloudObsAccessKeyID = gozen.ConfigAppGetString("HuaweicloudObsAccessKeyID", "")
	huaweicloudObsSecretAccessKey = gozen.ConfigAppGetString("HuaweicloudObsSecretAccessKey", "")
	huaweicloudObsEndpoint = gozen.ConfigAppGetString("HuaweicloudObsEndpoint", "")
	huaweicloudObsServiceUrl = gozen.ConfigAppGetString("HuaweicloudObsServiceUrl", "")
	huaweicloudObsBucket = gozen.ConfigAppGetString("HuaweicloudObsBucket", "")
	huaweicloudObsCustomerKey = gozen.ConfigAppGetString("HuaweicloudObsCustomerKey", "")

	var err error
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// yourAccessKeyId和yourAccessKeySecret填写RAM用户的访问密钥（AccessKey ID和AccessKey Secret）。
	obsClient, err = hobs.New(huaweicloudObsAccessKeyID, huaweicloudObsSecretAccessKey, huaweicloudObsEndpoint)
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
		os.Exit(-1)
	}

}
