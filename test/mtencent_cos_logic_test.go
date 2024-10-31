package test

import (
	"fmt"
	"github.com/namejlt/gozenmodule/mtencent/cos"
	"testing"
)

func TestMtencentCosUploadFile(t *testing.T) {

	tenantId := 134589

	appName := "touch"

	filePath := "./readme.md"

	fileName := fmt.Sprintf("%d/%s/%s", tenantId, appName, "readme.md")
	m := cos.Mtencentcos{}
	r, err := m.UploadFile(fileName, filePath)

	t.Log(r)
	t.Log(err)
}

func TestMtencentCosAES256UploadFile(t *testing.T) {

	tenantId := 134589

	appName := "touch"

	filePath := "./readme.md"

	fileName := fmt.Sprintf("%d/%s/%s", tenantId, appName, "readme.md")
	m := cos.Mtencentcos{}
	r, err := m.UploadFileAES256(fileName, filePath)

	t.Log(r)
	t.Log(err)
}

func TestMtencentCosAES256DownloadFile(t *testing.T) {

	key := "6e4a8deaf0c63e2522ffe5410355dc80"
	m := cos.Mtencentcos{}
	r, err := m.DownloadFileAES256(key)
	t.Log(string(r))
	t.Log(err)
}
