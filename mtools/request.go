package mtools

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/namejlt/gozen"
)

const (
	POST_JSON = "POST_JSON"
	POST_FORM = "POST_FORM"
	GET       = "GET"
)

func NewRequest() *Request {
	return &Request{}
}

type Request struct {
}

func (p *Request) Call(methodStr string, url string, paramByte []byte, headerSet map[string]string) (req *http.Request, body []byte, err error) {
	return p.CallWithContext(methodStr, url, paramByte, headerSet, nil)
}

func (p *Request) CallWithContext(methodStr string, url string, paramByte []byte, headerSet map[string]string, ctx context.Context) (
	req *http.Request, body []byte, err error) {
	methodStr = strings.ToUpper(methodStr)
	if headerSet == nil {
		headerSet = make(map[string]string)
	}
	method := "GET"
	switch methodStr {
	case "POST_FORM":
		headerSet["Content-Type"] = "application/x-www-form-urlencoded; param=value"
		method = "POST"
	case "POST_JSON":
		headerSet["Content-Type"] = "application/json"
		method = "POST"
	}
	req, err = http.NewRequest(method, url, bytes.NewBuffer(paramByte))
	if err != nil {
		gozen.LogErrorw(gozen.LogNameApi, "http.NewRequest",
			"err", err)
		return
	}
	for key, value := range headerSet {
		req.Header.Set(key, value)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	body, err = p.call(req)
	return req, body, err
}

func (p *Request) call(req *http.Request) (body []byte, err error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
