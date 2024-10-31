package mconfig

import (
	"encoding/json"
)

type ConfigTenantData string

func (p *ConfigTenantData) Decode(tag string) (data interface{}, err error) {
	//空值返回默认值
	if *p == "" {
		data = getDefaultConfigTenant(tag)
		return
	}
	switch tag {
	case ConfigModeTagAccess:
		ret := ConfigTenantDataAccess{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagCommon:
		ret := ConfigTenantDataCommon{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagTouch:
		ret := ConfigTenantDataTouch{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagCore:
		ret := ConfigTenantDataCore{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagTouchDo:
		ret := ConfigTenantDataTouchDo{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagOneId:
		ret := ConfigTenantDataOneId{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagSaasTranspond:
		ret := ConfigTenantDataSaasTranspond{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	}
	return
}

func (p *ConfigTenantData) Check(tag string) (b bool) {
	switch tag {
	case ConfigModeTagAccess:
		ret := ConfigTenantDataAccess{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagCommon:
		ret := ConfigTenantDataCommon{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagTouch:
		ret := ConfigTenantDataTouch{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagCore:
		ret := ConfigTenantDataCore{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagTouchDo:
		ret := ConfigTenantDataTouchDo{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagOneId:
		ret := ConfigTenantDataOneId{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagSaasTranspond:
		ret := ConfigTenantDataSaasTranspond{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	}
	b = true
	return
}

func getDefaultConfigTenant(tag string) (data interface{}) {
	switch tag {
	case ConfigModeTagAccess:
		data = getDefaultConfigTenantAccess()
	case ConfigModeTagCommon:
		data = getDefaultConfigTenantCommon()
	case ConfigModeTagTouch:
		data = getDefaultConfigTenantTouch()
	case ConfigModeTagCore:
		data = getDefaultConfigTenantCore()
	case ConfigModeTagTouchDo:
		data = getDefaultConfigTenantTouchDo()
	case ConfigModeTagOneId:
		data = getDefaultConfigTenantOneId()
	case ConfigModeTagSaasTranspond:
		data = getDefaultConfigTenantSaasTranspond()
	}
	return
}

func getDefaultConfigTenantAccess() ConfigTenantDataAccess {
	ret := ConfigTenantDataAccess{}
	ret.OneIdUseNewService = true
	return ret
}

func getDefaultConfigTenantCommon() ConfigTenantDataCommon {
	ret := ConfigTenantDataCommon{}
	return ret
}

func getDefaultConfigTenantTouch() ConfigTenantDataTouch {
	ret := ConfigTenantDataTouch{}
	return ret
}

func getDefaultConfigTenantCore() ConfigTenantDataCore {
	ret := ConfigTenantDataCore{}
	return ret
}

func getDefaultConfigTenantTouchDo() ConfigTenantDataTouchDo {
	ret := ConfigTenantDataTouchDo{}
	return ret
}

func getDefaultConfigTenantOneId() ConfigTenantDataOneId {
	ret := ConfigTenantDataOneId{}
	return ret
}

func getDefaultConfigTenantSaasTranspond() ConfigTenantDataSaasTranspond {
	ret := ConfigTenantDataSaasTranspond{}
	return ret
}

//======================================================================================================================
//================================================tag 固定标记===========================================================
//======================================================================================================================
//================================================access 数据接入配置=====================================================
//======================================================================================================================
//============================================下面是diy 修改的model=======================================================
//======================================================================================================================

// ===================================================
// ====================== Common ======================
// ===================================================

type ConfigTenantDataCommon struct {
	WxAgentId uint64 `json:"wx_agent_id"` // 企微应用id 默认0 多个这个是默认的应用
	WxCorpId  string `json:"wx_corp_id"`  // 企微企业id 默认空白 一个租户仅此一个
}

// ===================================================
// ====================== Access ======================
// ===================================================

type ConfigTenantDataAccess struct {
	OneIdUseNewService  bool   `json:"one_id_use_new_service"` // 是否使用新的oneid服务,默认是true
	DefaultEnterpriseId string `json:"default_enterprise_id"`  // 默认企业id - 企微联系人使用
	DefaultAppId        string `json:"default_app_id"`         // 默认app id - 微信app id使用
	DefaultSubjectId    string `json:"default_subject_id"`     // 默认主体id - 微信union id使用
}

// ===================================================
// ====================== Touch ======================
// ===================================================

type ConfigTenantDataTouch struct {
	MiniProgramList         []MiniProgramInfo `json:"mini_program_list"`           // 小程序列表
	WorkWxList              []WorkWxListInfo  `json:"work_wx_list"`                // 企业微信列表
	CouponChannelType       uint32            `json:"coupon_channel_type"`         // 券渠道类型 1 自有渠道 2 孩子王
	DownloadBmcEmpRoleCodes []string          `json:"download_bmc_emp_role_codes"` // 下载bmc员工角色列表
	CommonConfig
}

type MiniProgramInfo struct {
	Name       string `json:"name"`        // 小程序名称
	AppId      string `json:"app_id"`      // 小程序id
	OriginalId string `json:"original_id"` // 小程序原始id
}

type WorkWxListInfo struct {
	Name  string `json:"name"`   // 企业微信名称
	AppId string `json:"app_id"` // 企业微信id
}

// ===================================================
// ====================== Core ======================
// ===================================================

type ConfigTenantDataCore struct{}

// ===================================================
// ====================== TouchDo ======================
// ===================================================

type ConfigTenantDataTouchDo struct {
	CommonConfig
}

// ===================================================
// ====================== OneId ======================
// ===================================================

type ConfigTenantDataOneId struct {
	ReportVCConfigs []ReportVirtualChannelConfig `json:"report_vc_configs"` // 上报虚拟渠道配置
	DataEncrypt     bool                         `json:"data_encrypt"`      // 数据加密,针对租户配置
}

type ReportVirtualChannelConfig struct {
	VirtualChannel    uint32   `json:"virtual_channel"`     // 虚拟渠道id
	ActualChannelList []uint32 `json:"actual_channel_list"` // 实际渠道id列表
}

// ===================================================
// ====================== SaasTranspond ======================
// ===================================================

type ConfigTenantDataSaasTranspond struct {
	LevelList []LevelList `json:"level_list"` // 层级配置
}

type LevelList struct {
	Name    string      `json:"name"`     // 名称
	Level   int8        `json:"level"`    // level
	OrgType string      `json:"org_type"` // 层级类型
	Ext     interface{} `json:"ext"`      // 扩展字段
}
