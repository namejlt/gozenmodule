package mconfig

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/namejlt/gozen"
	"net/url"
	"time"
)

var (
	configCustomerTag = gozen.ConfigAppGetString("SaasConfigCustomerTag", "")
	configNeibuUrl    = gozen.ConfigAppGetString("SaasConfigNeibuUrl", "")
)

const (
	timeOutHttpDefault = 10 * time.Second
	codeOk             = 1001
)

type configModeDataReq struct {
	CustomerTag string `json:"customer_tag" form:"customer_tag" binding:"required"` //客户标记
	Tag         string `json:"tag" form:"tag" binding:"required"`                   //模式标记
}

type configModeDataRes struct {
	Code    int            `json:"code" example:"1001"`  // code
	Message string         `json:"message" example:"成功"` // message
	Data    ConfigModeData `json:"data"`                 // data
}

func configModeData(tag string) (data ConfigModeData, err error) {
	param := configModeDataReq{}
	param.CustomerTag = configCustomerTag
	param.Tag = tag
	reqUrl := configNeibuUrl + "i/v1/config/mode/data/info?" +
		fmt.Sprintf("customer_tag=%s&tag=%s", url.QueryEscape(param.CustomerTag), url.QueryEscape(param.Tag))
	ctx := context.Background()
	ret, err := gozen.CurlGet(ctx, reqUrl, []string{}, timeOutHttpDefault)
	if err != nil {
		gozen.LogInfow(gozen.LogNameApi, "configModeData curl error",
			gozen.LogKNameCommonUrl, reqUrl,
			gozen.LogKNameCommonData, param,
			gozen.LogKNameCommonErr, err,
		)
		return
	}
	var retData configModeDataRes
	err = json.Unmarshal(ret, &retData)
	if err != nil {
		gozen.LogErrorw(gozen.LogNameApi, "configModeData json  ret unmarshal error",
			gozen.LogKNameCommonData, string(ret),
			gozen.LogKNameCommonErr, err,
		)
		return
	}
	if retData.Code != codeOk {
		gozen.LogInfow(gozen.LogNameApi, "configModeData code fail",
			gozen.LogKNameCommonUrl, reqUrl,
			gozen.LogKNameCommonData, param,
			gozen.LogKNameCommonRes, string(ret),
		)
		err = errors.New("configModeData code fail:" + string(ret))
		return
	}
	data = retData.Data
	return
}

type configTenantListReq struct {
	CustomerTag string `json:"customer_tag" form:"customer_tag" binding:"required"` //客户标记
	Tag         string `json:"tag" form:"tag" binding:"required"`                   //模式标记
}

type configTenantListRes struct {
	Code    int                `json:"code" example:"1001"`  // code
	Message string             `json:"message" example:"成功"` // message
	Data    []ConfigTenantInfo `json:"data"`                 // data
}

type ConfigTenantInfo struct {
	TenantId uint64           `json:"tenant_id"` // 租户id
	Config   ConfigTenantData `json:"config"`    // config
}

func configTenantList(tag string) (data []ConfigTenantInfo, err error) {
	param := configTenantListReq{}
	param.CustomerTag = configCustomerTag
	param.Tag = tag
	reqUrl := configNeibuUrl + "i/v1/config/tenant/data/list?" +
		fmt.Sprintf("customer_tag=%s&tag=%s", url.QueryEscape(param.CustomerTag), url.QueryEscape(param.Tag))
	ctx := context.Background()
	ret, err := gozen.CurlGet(ctx, reqUrl, []string{}, timeOutHttpDefault)
	if err != nil {
		gozen.LogInfow(gozen.LogNameApi, "configTenantList curl error",
			gozen.LogKNameCommonUrl, reqUrl,
			gozen.LogKNameCommonData, param,
			gozen.LogKNameCommonErr, err,
		)
		return
	}
	var retData configTenantListRes
	err = json.Unmarshal(ret, &retData)
	if err != nil {
		gozen.LogErrorw(gozen.LogNameApi, "configTenantList json  ret unmarshal error",
			gozen.LogKNameCommonData, string(ret),
			gozen.LogKNameCommonErr, err,
		)
		return
	}
	if retData.Code != codeOk {
		gozen.LogInfow(gozen.LogNameApi, "configTenantList code fail",
			gozen.LogKNameCommonUrl, reqUrl,
			gozen.LogKNameCommonData, param,
			gozen.LogKNameCommonRes, string(ret),
		)
		err = errors.New("configTenantList code fail:" + string(ret))
		return
	}
	data = retData.Data
	return
}
