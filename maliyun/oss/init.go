package oss

import (
	"fmt"
	aoss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/namejlt/gozen"
	"os"
)

/**

营销云专用


每一个客户环境申请一个新的bucket


仅用于内部上传下载，没有流量加速

*/

var (
	aliyunOssEndpoint        string
	aliyunOssServiceUrl      string
	aliyunOssAccessKeyId     string
	aliyunOssAccessKeySecret string
	aliyunOssBucket          string
	aliyunOssCustomerKey     string

	aliyunClient *aoss.Client
)

func init() {
	aliyunOssAccessKeyId = gozen.ConfigAppGetString("AliyunOssAccessKeyId", "")
	aliyunOssAccessKeySecret = gozen.ConfigAppGetString("AliyunOssAccessKeySecret", "")
	aliyunOssEndpoint = gozen.ConfigAppGetString("AliyunOssEndpoint", "")
	aliyunOssServiceUrl = gozen.ConfigAppGetString("AliyunOssServiceUrl", "")
	aliyunOssBucket = gozen.ConfigAppGetString("AliyunOssBucket", "")
	aliyunOssCustomerKey = gozen.ConfigAppGetString("AliyunOssCustomerKey", "")

	var err error
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// yourAccessKeyId和yourAccessKeySecret填写RAM用户的访问密钥（AccessKey ID和AccessKey Secret）。
	aliyunClient, err = aoss.New(aliyunOssEndpoint, aliyunOssAccessKeyId, aliyunOssAccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

}
