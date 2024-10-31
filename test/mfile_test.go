package test

import (
	"fmt"
	"github.com/namejlt/gozenmodule/mfile"
	"testing"
)

func TestMfileUploadFile(t *testing.T) {
	tenantId := 134589

	appName := "touch"

	filePath := "./readme.md"

	fileName := fmt.Sprintf("%d/%s/%s", tenantId, appName, "readme.md")
	r, err := mfile.UploadFile(fileName, filePath)

	t.Log(r)
	t.Log(err)
}

func TestMfileAES256UploadFile(t *testing.T) {

	tenantId := 134589

	appName := "touch"

	filePath := "./readme.md"

	fileName := fmt.Sprintf("%d/%s/%s", tenantId, appName, "readme.md")
	r, err := mfile.UploadFileAES256(fileName, filePath)

	t.Log(r)
	t.Log(err)
}

func TestMfileAES256DownloadFileTencent(t *testing.T) {

	key := "6e4a8deaf0c63e2522ffe5410355dc80"
	r, err := mfile.DownloadFileAES256(key)
	t.Log(string(r))
	t.Log(err)
}

func TestMfileAES256DownloadFileByFileName(t *testing.T) {

	key := "134589/touch/readme.md"
	r, err := mfile.DownloadFileAES256(key)
	t.Log(string(r))
	t.Log(err)
}
