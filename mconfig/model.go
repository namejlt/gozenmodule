package mconfig

import (
	"encoding/json"
	"github.com/namejlt/gozenmodule/mconfig/mconst"
)

type ConfigModeData string

func (p *ConfigModeData) Decode(tag string) (data interface{}, err error) {
	//空值返回默认值
	if *p == "" {
		data = getDefaultConfig(tag)
		return
	}
	switch tag {
	case ConfigModeTagAccess:
		ret := ConfigModeDataAccess{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagCommon:
		ret := ConfigModeDataCommon{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagTouch:
		ret := ConfigModeDataTouch{}
		err = json.Unmarshal([]byte(*p), &ret)
		ret.setDefault()
		data = ret
	case ConfigModeTagCore:
		ret := ConfigModeDataCore{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagTouchDo:
		ret := ConfigModeDataTouchDo{}
		err = json.Unmarshal([]byte(*p), &ret)
		ret.setDefault()
		data = ret
	case ConfigModeTagOneId:
		ret := ConfigModeDataOneId{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	case ConfigModeTagSaasTranspond:
		ret := ConfigModeDataSaasTranspond{}
		err = json.Unmarshal([]byte(*p), &ret)
		data = ret
	}
	return
}

func (p *ConfigModeData) Check(tag string) (b bool) {
	switch tag {
	case ConfigModeTagAccess:
		ret := ConfigModeDataAccess{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagCommon:
		ret := ConfigModeDataCommon{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagTouch:
		ret := ConfigModeDataTouch{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagCore:
		ret := ConfigModeDataCore{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagTouchDo:
		ret := ConfigModeDataTouchDo{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	case ConfigModeTagOneId:
		ret := ConfigModeDataOneId{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}

		// 配置校验

		// 校验字段是否在配置列表内
		for _, field := range ret.GenOneIdFields {
			if _, exist := mconst.OneIdFieldMap[field]; !exist {
				return
			}
		}
		for _, field := range ret.ExtSaveDbFields {
			if _, exist := mconst.OneIdFieldMap[field]; !exist {
				return
			}
		}
		for _, field := range ret.OneIdMergeFields {
			if _, exist := mconst.OneIdFieldMap[field]; !exist {
				return
			}
		}

		if ret.OneIdMode == mconst.OneIdModeFieldsPriority {
			// 优先级模式, 生成one id的字段列表不能为空
			if len(ret.GenOneIdFields) == 0 {
				return
			}
		}
	case ConfigModeTagSaasTranspond:
		ret := ConfigModeDataSaasTranspond{}
		err := json.Unmarshal([]byte(*p), &ret)
		if err != nil {
			return
		}
	}
	b = true
	return
}

func getDefaultConfig(tag string) (data interface{}) {
	switch tag {
	case ConfigModeTagAccess:
		ret := ConfigModeDataAccess{}
		data = ret
	case ConfigModeTagCommon:
		ret := ConfigModeDataCommon{}
		data = ret
	case ConfigModeTagTouch:
		ret := ConfigModeDataTouch{}
		ret.WorkWxType = mconst.WorkWxTypeLt
		ret.CouponChannelType = mconst.RewardChannelTypeSelf
		data = ret
	case ConfigModeTagCore:
		ret := ConfigModeDataCore{}
		ret.IntelligentModelType = mconst.IntelligentModelTypeApply
		data = ret
	case ConfigModeTagTouchDo:
		ret := ConfigModeDataTouchDo{}
		ret.WorkWxType = mconst.WorkWxTypeLt
		data = ret
	case ConfigModeTagOneId:
		ret := ConfigModeDataOneId{}
		ret.OneIdMode = mconst.OneIdModeAllId // 默认是全id合并模式
		ret.IsUidMustBindChannel = true       // 默认是true
		data = ret
	case ConfigModeTagSaasTranspond:
		ret := ConfigModeDataSaasTranspond{}
		data = ret
	}
	return
}

//======================================================================================================================
//================================================tag 固定标记===========================================================
//======================================================================================================================
//================================================access 数据接入配置=====================================================
//======================================================================================================================
//============================================下面是diy 修改的model=======================================================
//======================================================================================================================

const (
	ConfigModeTagAccess        = "access"
	ConfigModeTagCommon        = "common"
	ConfigModeTagTouch         = "touch"
	ConfigModeTagCore          = "core"
	ConfigModeTagTouchDo       = "touchdo"
	ConfigModeTagOneId         = "oneid"
	ConfigModeTagSaasTranspond = "saastranspond"
)

// ===================================================
// ====================== common ======================
// ===================================================

type CommonConfig struct {
	OrgAuthSwitch bool `json:"org_auth_switch"` // 组织权限开关 默认关
}

// ===================================================
// ====================== Access ======================
// ===================================================

type ConfigModeDataAccess struct { //access 服务配置
	FieldEventSingle bool `json:"field_event_single"` //字段事件是否单独 默认非单独

	// 租户默认品牌id
	TenantDefaultBrandIdNameList []ConfigModeDataAccessTenantBrand `json:"tenant_default_brand_id_name_list"` //每次取循环获取
}

type ConfigModeDataAccessTenantBrand struct {
	TenantId  uint64 `json:"tenant_id"`
	BrandId   uint32 `json:"brand_id"`
	BrandName string `json:"brand_name"`
}

// ===================================================
// ====================== Common ======================
// ===================================================

type ConfigModeDataCommon struct { //公共配置
}

// ===================================================
// ====================== Touch ======================
// ===================================================

type ConfigModeDataTouch struct {
	WorkWxType               int          `json:"work_wx_type"`                // 企业微信服务类型 1 联童企微 2 ICC
	MerchantEnvironment      Compare[int] `json:"merchant_environment"`        // 商户环境
	CouponChannelType        uint32       `json:"coupon_channel_type"`         // 券渠道类型 1 自有渠道 2 孩子王
	ServiceChannelType       uint32       `json:"service_channel_type"`        // 服务渠道类型 1 自有渠道
	ServiceCouponChannelType uint32       `json:"service_coupon_channel_type"` // 服务券渠道类型 1 自有渠道
	CommonConfig
}

func (p *ConfigModeDataTouch) setDefault() {
	if p.WorkWxType == 0 {
		p.WorkWxType = mconst.WorkWxTypeLt
	}
	if p.CouponChannelType == 0 {
		p.CouponChannelType = mconst.RewardChannelTypeSelf
	}
	if p.ServiceChannelType == 0 {
		p.ServiceChannelType = mconst.RewardChannelTypeSelf
	}
	if p.ServiceCouponChannelType == 0 {
		p.ServiceCouponChannelType = mconst.RewardChannelTypeSelf
	}
}

// Compare 可对比公用类型自带对比方法
type Compare[T comparable] struct {
	Value T `json:"value"`
}

func (s Compare[T]) MatchSuccess(v T) bool {
	return s.Value == v
}

// ===================================================
// ====================== Core ======================
// ===================================================

type ConfigModeDataCore struct {
	IntelligentModelType           uint8  `json:"intelligent_model_type"`            //智能模型包类型 1 购买 2 申请 只能使用一种模式   ----- 弃用
	ShowMobilePlainText            bool   `json:"show_mobile_plain_text"`            //是否展示手机号明文 默认否
	ShoppingAssistantConfigDefault string `json:"shopping_assistant_config_default"` // 导购工作台默认配置
	BossWorkbenchConfigDefault     string `json:"boss_workbench_config_default"`     // boss工作台默认配置
}

// ===================================================
// ====================== TouchDo ======================
// ===================================================

type ConfigModeDataTouchDo struct {
	WorkWxType int `json:"work_wx_type"` // 企业微信服务类型 1 联童企微 2 ICC
	CommonConfig
}

func (p *ConfigModeDataTouchDo) setDefault() {
	if p.WorkWxType == 0 {
		p.WorkWxType = mconst.WorkWxTypeLt
	}
}

// ===================================================
// ====================== OneId ======================
// ===================================================

type ConfigModeDataOneId struct {
	IsOutageMaintain     bool                `json:"is_outage_maintain"`       // 是否停机维护
	OneIdMode            uint8               `json:"one_id_mode"`              // one id 模式 1:全id合并模式 2:优先级生成模式
	IsOneIdUseUidDirect  bool                `json:"is_one_id_use_uid_direct"` // one id是否直接使用uid,只有模式2才有效,并且uid必须能解析成数字
	IsUidMustBindChannel bool                `json:"is_uid_must_bind_channel"` // uid是否必须绑定渠道,只有模式2才有效,为true,db里存的渠道是0
	DefaultChannel       uint32              `json:"default_channel"`          // 默认渠道,只有uid必须绑定渠道才有效
	DataEncrypt          bool                `json:"data_encrypt"`             // 数据加密,针对环境配置
	GenOneIdFields       []mconst.OneIdField `json:"gen_one_id_fields"`        // 生成one id的字段,只有模式2才有效,模式2下有先后顺序
	ExtSaveDbFields      []mconst.OneIdField `json:"ext_save_db_fields"`       // 额外保存到db的字段,只有模式2才有效 -- 不用来生成one id,但需要存储到db
	OneIdMergeFields     []mconst.OneIdField `json:"one_id_merge_fields"`      // one id合并字段,只有模式2才有效,谨慎配置
}

type ConfigModeDataSaasTranspond struct {
}
