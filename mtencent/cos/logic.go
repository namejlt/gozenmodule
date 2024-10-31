package cos

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/namejlt/gozen"
	tcos "github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/**

下载地址

尽量内部系统使用
下载地址有相关生命周期限制


参考文档

https://www.tencentcloud.com/zh/document/product/436/44063

*/

// UploadFileByTenantPath 通过租户应用目录上传文件 返回地址可以直接下载
func (p *Mtencentcos) UploadFileByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (fileUrl string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)

	return p.UploadFile(fileName, filePath)
}

func (p *Mtencentcos) UploadFile(fileName string, filePath string) (fileUrl string, err error) {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.tencentcloud.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.tencentcloud.com/ ，关于地域的详情见 https://www.tencentcloud.com/document/product/436/6224?from_cn_redirect=1 。
	u, _ := url.Parse(tencentCloudCosBucketUrl)
	b := &tcos.BaseURL{BucketURL: u}
	client := tcos.NewClient(b, &http.Client{
		Transport: &tcos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretID: tencentCloudCosSecretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretKey: tencentCloudCosSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
		},
	})

	/**

	key 对象键（Key）是对象在存储桶中的唯一标识。例如，在对象的访问域名examplebucket-1250000000.cos.ap-guangzhou.myqcloud.com/doc/pic.jpg中，对象键为 doc/pic.jpg
	filepath 本地文件名

	*/
	key := fileName
	result, _, err := client.Object.Upload(
		context.Background(), key, filePath, nil,
	)
	/**
	type CompleteMultipartUploadResult struct {
	  Location string URL 地址
	  Bucket   string 存储桶名称，格式：BucketName-APPID。例如 examplebucket-1250000000
	  Key      string 对象键（Key）是对象在存储桶中的唯一标识。例如，在对象的访问域名 examplebucket-1250000000.cos.ap-guangzhou.myqcloud.com/doc/pic.jpg 中，对象键为 doc/pic.jpg
	  ETag     string 合并后对象的唯一标签值，该值不是对象内容的 MD5 校验值，仅能用于检查对象唯一性。如需校验文件内容，可以在上传过程中校验单个分块的 ETag 值
	}

	*/
	if err != nil {
		return
	}
	if tencentCloudCosServiceUrl == "" {
		fileUrl = result.Location
	} else {
		//加速访问   加速访问域名/上传路径
		fileUrl = tencentCloudCosServiceUrl + "/" + fileName
	}

	return
}

/**

加密上传
加密下载

https://www.tencentcloud.com/zh/document/product/436/38120

该加密所运行的服务需要使用 HTTPS 请求。
customerKey：用户提供的密钥，传入一个32字节的字符串，支持数字、字母、字符的组合，不支持中文。
如果上传的源文件调用了该方法，那么在使用 GET（下载）、HEAD（查询）时对源对象操作的时候也要调用该方法。

加密上传必须用加密下载函数进行下载

*/

func (p *Mtencentcos) UploadFileAES256ByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (key string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)

	return p.UploadFileAES256(fileName, filePath)
}

func (p *Mtencentcos) UploadFileAES256(fileName string, filePath string) (key string, err error) {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.tencentcloud.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.tencentcloud.com/ ，关于地域的详情见 https://www.tencentcloud.com/document/product/436/6224?from_cn_redirect=1 。
	u, _ := url.Parse(tencentCloudCosBucketUrl)
	b := &tcos.BaseURL{BucketURL: u}
	client := tcos.NewClient(b, &http.Client{
		Transport: &tcos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretID: tencentCloudCosSecretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretKey: tencentCloudCosSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
		},
	})

	customerKeyMD5 := utilCryptoMd5(tencentCloudCosCustomerKey)

	opt := &tcos.ObjectPutOptions{
		ObjectPutHeaderOptions: &tcos.ObjectPutHeaderOptions{
			ContentType:           "text/html",
			XCosSSECustomerAglo:   "AES256",
			XCosSSECustomerKey:    base64.StdEncoding.EncodeToString([]byte(tencentCloudCosCustomerKey)),
			XCosSSECustomerKeyMD5: customerKeyMD5,
		},
		ACLHeaderOptions: &tcos.ACLHeaderOptions{},
	}

	/**

	key 对象键（Key）是对象在存储桶中的唯一标识。例如，在对象的访问域名examplebucket-1250000000.cos.ap-guangzhou.myqcloud.com/doc/pic.jpg中，对象键为 doc/pic.jpg
	filepath 本地文件名

	*/
	key = gozen.UtilCryptoMd5Lower(fileName)

	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = client.Object.Put(
		context.Background(), key, f, opt,
	)
	if err != nil {
		return
	}

	return
}

func (p *Mtencentcos) DownloadFileAES256(key string) (bodyBytes []byte, err error) {
	u, _ := url.Parse(tencentCloudCosBucketUrl)
	b := &tcos.BaseURL{BucketURL: u}
	client := tcos.NewClient(b, &http.Client{
		Transport: &tcos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretID: tencentCloudCosSecretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretKey: tencentCloudCosSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
		},
	})

	customerKeyMD5 := utilCryptoMd5(tencentCloudCosCustomerKey)

	getopt := &tcos.ObjectGetOptions{
		XCosSSECustomerAglo:   "AES256",
		XCosSSECustomerKey:    base64.StdEncoding.EncodeToString([]byte(tencentCloudCosCustomerKey)),
		XCosSSECustomerKeyMD5: customerKeyMD5,
	}

	var resp *tcos.Response
	resp, err = client.Object.Get(context.Background(), key, getopt)
	if err != nil {
		return
	}
	bodyBytes, _ = io.ReadAll(resp.Body)

	return
}

func DownloadFileAES256Customer(key string, bucketUrl string, secretID string, secretKey string, customerKey string) (bodyBytes []byte, err error) {
	u, err := url.Parse(bucketUrl)
	if err != nil {
		return
	}
	b := &tcos.BaseURL{BucketURL: u}
	client := tcos.NewClient(b, &http.Client{
		Transport: &tcos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretID: secretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretKey: secretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
		},
	})

	customerKeyMD5 := utilCryptoMd5(customerKey)

	getopt := &tcos.ObjectGetOptions{
		XCosSSECustomerAglo:   "AES256",
		XCosSSECustomerKey:    base64.StdEncoding.EncodeToString([]byte(customerKey)),
		XCosSSECustomerKeyMD5: customerKeyMD5,
	}

	var resp *tcos.Response
	resp, err = client.Object.Get(context.Background(), key, getopt)
	if err != nil {
		return
	}
	bodyBytes, _ = io.ReadAll(resp.Body)

	return
}

func utilCryptoMd5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return base64.StdEncoding.EncodeToString(cipherStr)
}
