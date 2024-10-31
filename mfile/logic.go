package mfile

import (
	"fmt"
	"github.com/namejlt/gozenmodule/maliyun/oss"
	"github.com/namejlt/gozenmodule/mhuaweicloud/obs"
	"github.com/namejlt/gozenmodule/mtencent/cos"
	"strings"
)

var (
	MFileService MFile
)

func registerMode(mode string) {
	switch mode {
	case fileServiceModeTencent:
		MFileService = new(cos.Mtencentcos)
	case fileServiceModeAliyun:
		MFileService = new(oss.Maliyuncos)
	case fileServiceModeHuaweicloud:
		MFileService = new(obs.HuaweicloudObs)
	default:
		MFileService = new(Mdefault)
	}
}

//======================= 通用方法

func UploadFileByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (fileUrl string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)
	return MFileService.UploadFile(fileName, filePath)
}

func UploadFile(fileName string, filePath string) (fileUrl string, err error) {
	return MFileService.UploadFile(fileName, filePath)
}

func UploadFileAES256ByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (key string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)
	return MFileService.UploadFileAES256(fileName, filePath)
}

func UploadFileAES256(fileName string, filePath string) (key string, err error) {
	return MFileService.UploadFileAES256(fileName, filePath)
}

func DownloadFileAES256(key string) (bodyBytes []byte, err error) {
	return MFileService.DownloadFileAES256(key)
}
