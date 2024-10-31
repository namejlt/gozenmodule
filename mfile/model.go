package mfile

import (
	"fmt"
	"strings"
)

type Mdefault struct {
}

// UploadFileByTenantPath 通过租户应用目录上传文件 返回地址可以直接下载
func (p *Mdefault) UploadFileByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (fileUrl string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)

	return p.UploadFile(fileName, filePath)
}

func (p *Mdefault) UploadFile(fileName string, filePath string) (fileUrl string, err error) {

	return
}

func (p *Mdefault) UploadFileAES256ByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (key string, err error) {
	fileName = fmt.Sprintf("%d/%s/%s", tenantId, strings.ReplaceAll(appName, " ", ""), fileName)

	return p.UploadFileAES256(fileName, filePath)
}

func (p *Mdefault) UploadFileAES256(fileName string, filePath string) (key string, err error) {

	return
}

func (p *Mdefault) DownloadFileAES256(key string) (bodyBytes []byte, err error) {

	return
}
