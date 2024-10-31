package mfile

type MFile interface {
	UploadFileByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (fileUrl string, err error)
	UploadFile(fileName string, filePath string) (fileUrl string, err error)
	UploadFileAES256ByTenantPath(tenantId uint64, appName string, fileName string, filePath string) (key string, err error)
	UploadFileAES256(fileName string, filePath string) (key string, err error)
	DownloadFileAES256(key string) (bodyBytes []byte, err error)
}
