package mcommon

type BaseJavaResp struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
}

type BaseGoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
