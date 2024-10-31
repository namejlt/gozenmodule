package obs

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	hobs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"io"
	"os"
	"strings"
)

/**

下载地址

尽量内部系统使用
下载地址有相关生命周期限制


参考文档

https://support.huaweicloud.com/sdk-go-devg-obs/obs_23_0403.html

*/

// UploadFileByTenantPath 通过租户应用目录上传文件 返回地址可以直接下载
func (p *HuaweicloudObs) UploadFileByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (fileUrl string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)

	return p.UploadFile(fileName, filePath)
}

func (p *HuaweicloudObs) UploadFile(fileName string, filePath string) (fileUrl string, err error) {

	input := &hobs.PutFileInput{}
	// 指定存储桶名称
	input.Bucket = huaweicloudObsBucket
	// 指定上传对象，此处以 example/objectname 为例。
	input.Key = fileName
	// 指定本地文件，此处以localfile为例
	input.SourceFile = filePath
	// 文件上传
	output, err := obsClient.PutFile(input)
	if err != nil {
		return
	}

	if huaweicloudObsServiceUrl == "" {
		fileUrl = output.ObjectUrl
	} else {
		fileUrl = huaweicloudObsServiceUrl + "/" + fileName
	}

	return
}

/**

加密上传
加密下载

https://help.aliyun.com/zh/oss/developer-reference/download-objects-as-files-4?spm=a2c4g.11186623.0.0.fb1c7ae8edANvq

https://help.aliyun.com/zh/oss/user-guide/server-side-encryption-8#concept-lqm-fkd-5db

*/

func (p *HuaweicloudObs) UploadFileAES256ByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (key string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)

	return p.UploadFileAES256(fileName, filePath)
}

func (p *HuaweicloudObs) UploadFileAES256(fileName string, filePath string) (key string, err error) {

	input := &hobs.PutObjectInput{}
	// 指定存储桶名称
	input.Bucket = huaweicloudObsBucket
	// 指定上传对象名，此处以 example/objectname 为例。
	input.Key = fileName
	// 指定上传的内容
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	input.Body = f

	customerKeyMD5 := utilCryptoMd5(huaweicloudObsCustomerKey)
	input.SseHeader = hobs.SseCHeader{
		Encryption: "AES256",
		Key:        base64.StdEncoding.EncodeToString([]byte(huaweicloudObsCustomerKey)), // 32byteslongsecretkeymustprovided
		KeyMD5:     customerKeyMD5,
	}

	// 流式上传本地文件
	_, err = obsClient.PutObject(input)
	if err != nil {
		return
	}

	key = fileName

	return
}

func (p *HuaweicloudObs) DownloadFileAES256(key string) (bodyBytes []byte, err error) {
	input := &hobs.GetObjectInput{}
	// 指定存储桶名称
	input.Bucket = huaweicloudObsBucket
	// 指定下载对象，此处以 example/objectname 为例。
	input.Key = key
	// 指定服务端加密头信息，此处以 obs.SseCHeader为例
	customerKeyMD5 := utilCryptoMd5(huaweicloudObsCustomerKey)
	input.SseHeader = hobs.SseCHeader{
		Encryption: "AES256",
		Key:        base64.StdEncoding.EncodeToString([]byte(huaweicloudObsCustomerKey)), // 32byteslongsecretkeymustprovided
		KeyMD5:     customerKeyMD5,
	}
	// 流式下载对象
	output, err := obsClient.GetObject(input)
	if err != nil {
		return
	}
	// output.Body 在使用完毕后必须关闭，否则会造成连接泄漏。
	defer output.Body.Close()
	// 读取对象内容
	bodyBytes, _ = io.ReadAll(output.Body)

	return

}

func DownloadFileAES256Customer(key string, bucketUrl string, secretID string, secretKey string, customerKey string, endpoint string) (bodyBytes []byte, err error) {
	obsClient, err = hobs.New(secretID, secretKey, endpoint)
	if err != nil {
		return
	}
	input := &hobs.GetObjectInput{}
	// 指定存储桶名称
	input.Bucket = bucketUrl
	// 指定下载对象，此处以 example/objectname 为例。
	input.Key = key
	// 指定服务端加密头信息，此处以 obs.SseCHeader为例
	customerKeyMD5 := utilCryptoMd5(customerKey)
	input.SseHeader = hobs.SseCHeader{
		Encryption: "AES256",
		Key:        base64.StdEncoding.EncodeToString([]byte(customerKey)), // 32byteslongsecretkeymustprovided
		KeyMD5:     customerKeyMD5,
	}
	// 流式下载对象
	output, err := obsClient.GetObject(input)
	if err != nil {
		return
	}
	// output.Body 在使用完毕后必须关闭，否则会造成连接泄漏。
	defer output.Body.Close()
	// 读取对象内容
	bodyBytes, _ = io.ReadAll(output.Body)

	return
}

func utilCryptoMd5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return base64.StdEncoding.EncodeToString(cipherStr)
}
