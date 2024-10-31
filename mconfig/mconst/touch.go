package mconst

const (
	MerchantEnvironmentSaas = 1 // saas 环境
	MerchantEnvironmentDR   = 2 // dr 环境
	MerchantEnvironmentBY   = 3 // 白药环境
	MerchantEnvironmentHzw  = 4 // 孩子王环境
)

const (
	WorkWxTypeLt  = 1 // 企微服务 联童
	WorkWxTypeICC = 2 // 企微服务 icc
)

const (
	RewardChannelTypeSelf = 1 // 自有渠道
)

const (
	IntelligentModelTypeBuy   = 1 // 智能模型包类型  购买
	IntelligentModelTypeApply = 2 // 智能模型包类型 申请
)

const (
	OneIdModeAllId          = 1 // 全id合并模式
	OneIdModeFieldsPriority = 2 // 字段优先级生成模式
)

// one id的字段

type OneIdField string

const (
	OneIdFieldMobile      OneIdField = "mobile"       // mobile
	OneIdFieldUid         OneIdField = "uid"          // uid
	OneIdFieldOpenId      OneIdField = "open_id"      // open_id
	OneIdFieldUnionId     OneIdField = "union_id"     // union_id
	OneIdFieldIdCard      OneIdField = "id_card"      // id_card
	OneIdFieldExternalUid OneIdField = "external_uid" // external_uid
	OneIdFieldDeviceNo    OneIdField = "device_no"    // device_no
	OneIdFieldSpExtId     OneIdField = "sp_ext_id"    // sp_ext_id
)

var OneIdFieldMap = map[OneIdField]struct{}{
	OneIdFieldMobile:      {},
	OneIdFieldUid:         {},
	OneIdFieldOpenId:      {},
	OneIdFieldUnionId:     {},
	OneIdFieldIdCard:      {},
	OneIdFieldExternalUid: {},
	OneIdFieldDeviceNo:    {},
	OneIdFieldSpExtId:     {},
}
