package test

import (
	"fmt"
	"github.com/namejlt/gozenmodule/maliyun/oss"
	"testing"
)

func TestMtencentOssUploadFile(t *testing.T) {
	tenantId := 134589

	appName := "touch"

	filePath := "./readme.md"

	fileName := fmt.Sprintf("%d/%s/%s", tenantId, appName, "readme.md")
	m := oss.Maliyuncos{}
	r, err := m.UploadFile(fileName, filePath)

	t.Log(r)
	t.Log(err)
}

func TestMtencentOssAES256UploadFile(t *testing.T) {

	tenantId := 134589

	appName := "touch"

	filePath := "./readme.md"

	fileName := fmt.Sprintf("%d/%s/%s", tenantId, appName, "readme.md")
	m := oss.Maliyuncos{}
	r, err := m.UploadFileAES256(fileName, filePath)

	t.Log(r)
	t.Log(err)
}

func TestMtencentOssAES256DownloadFile(t *testing.T) {

	key := "134589/touch/readme.md"
	m := oss.Maliyuncos{}
	r, err := m.DownloadFileAES256(key)
	t.Log(string(r))
	t.Log(err)
}
