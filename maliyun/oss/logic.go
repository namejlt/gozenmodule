package oss

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	aoss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"strings"
)

/**

下载地址

尽量内部系统使用
下载地址有相关生命周期限制


参考文档

https://help.aliyun.com/zh/oss/developer-reference/simple-upload-4?spm=a2c4g.11186623.0.0.71f56750Xvrmeb

*/

// UploadFileByTenantPath 通过租户应用目录上传文件 返回地址可以直接下载
func (p *Maliyuncos) UploadFileByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (fileUrl string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)

	return p.UploadFile(fileName, filePath)
}

func (p *Maliyuncos) UploadFile(fileName string, filePath string) (fileUrl string, err error) {

	// 填写存储空间名称，例如examplebucket。
	bucket, err := aliyunClient.Bucket(aliyunOssBucket)
	if err != nil {
		return
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile(fileName, filePath)
	if err != nil {
		return
	}

	fileUrl = aliyunOssServiceUrl + "/" + fileName
	return
}

/**

加密上传
加密下载

https://help.aliyun.com/zh/oss/developer-reference/download-objects-as-files-4?spm=a2c4g.11186623.0.0.fb1c7ae8edANvq

https://help.aliyun.com/zh/oss/user-guide/server-side-encryption-8#concept-lqm-fkd-5db

*/

func (p *Maliyuncos) UploadFileAES256ByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (key string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)

	return p.UploadFileAES256(fileName, filePath)
}

func (p *Maliyuncos) UploadFileAES256(fileName string, filePath string) (key string, err error) {

	// 初始化一个加密规则，加密方式以AES256为例。
	config := aoss.ServerEncryptionRule{SSEDefault: aoss.SSEDefaultRule{SSEAlgorithm: "AES256"}}
	err = aliyunClient.SetBucketEncryption(aliyunOssBucket, config)
	if err != nil {
		return
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := aliyunClient.Bucket(aliyunOssBucket)
	if err != nil {
		return
	}

	var opt []aoss.Option

	/*customerKeyMD5 := utilCryptoMd5(aliyunOssCustomerKey)

	opt = append(opt, aoss.ServerSideEncryption("AES256"))
	opt = append(opt, aoss.SSECAlgorithm("AES256"))
	opt = append(opt, aoss.SSECKey(aliyunOssCustomerKey))
	opt = append(opt, aoss.SSECKeyMd5(customerKeyMD5))*/

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile(fileName, filePath, opt...)
	if err != nil {
		return
	}

	key = fileName
	return
}

func (p *Maliyuncos) DownloadFileAES256(key string) (bodyBytes []byte, err error) {
	// 填写存储空间名称，例如examplebucket。
	bucket, err := aliyunClient.Bucket(aliyunOssBucket)
	if err != nil {
		return
	}

	var opt []aoss.Option

	/*customerKeyMD5 := utilCryptoMd5(aliyunOssCustomerKey)

	opt = append(opt, aoss.SSECAlgorithm("AES256"))
	opt = append(opt, aoss.SSECKey(aliyunOssCustomerKey))
	opt = append(opt, aoss.SSECKeyMd5(customerKeyMD5))*/
	// 下载文件到缓存。
	// yourObjectName填写不包含Bucket名称在内的Object的完整路径。
	body, err := bucket.GetObject(key, opt...)
	if err != nil {
		return
	}
	defer body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, body)
	if err != nil {
		return
	}
	bodyBytes = buf.Bytes()
	return
}

func DownloadFileAES256Customer(key string, bucketUrl string, secretID string, secretKey string, customerKey string, endpoint string) (bodyBytes []byte, err error) {

	client, err := aoss.New(endpoint, secretID, secretKey)
	if err != nil {
		return
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(bucketUrl)
	if err != nil {
		return
	}

	customerKeyMD5 := utilCryptoMd5(customerKey)

	var opt []aoss.Option

	opt = append(opt, aoss.SSECAlgorithm("AES256"))
	opt = append(opt, aoss.SSECKey(customerKey))
	opt = append(opt, aoss.SSECKeyMd5(customerKeyMD5))

	// 下载文件到缓存。
	// yourObjectName填写不包含Bucket名称在内的Object的完整路径。
	body, err := bucket.GetObject(key)
	if err != nil {
		return
	}
	defer body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, body)
	if err != nil {
		return
	}
	bodyBytes = buf.Bytes()

	return
}

func utilCryptoMd5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return base64.StdEncoding.EncodeToString(cipherStr)
}
