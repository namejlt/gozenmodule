package mcommon

const (
	LoginUidKey = "login_uid"
	LoginEmpKey = "login_emp"

	DevAuthKey = "dev_auth_key" // 开发者认证key

	TenantIdKey       = "_platform_num" //租户id key
	TenantAliasKey    = "tenant_alias"  //租户alias key
	TenantIdGrpcKey   = "tenantId"      //租户id grpc key
	TenantDataKey     = "tenant_data"   //租户信息
	TenantIdMobileKey = "_spGnxhw8jd_platform_num"

	CodeOk                         = 0
	CodeNoAccessAuth               = 401 // 非法访问，无权限!
	CodeNoAccessNullSign           = 402 // 非法访问，缺少sign字段!
	CodeNoAccessAuthorizationError = 403 // Http头Authorization错误
	CodeCommonOk                   = 1001
	CodeCommonAccessFail           = 1002
	CodeCommonServerBusy           = 1003
	CodeCommonParamsIncomplete     = 1004
	CodeCommonUserNoLogin          = 1005 // 未登录
	CodeDisasterReadonlyErr        = 1006 // 灾难错误，系统只读
	CodeCommonUserNoLoginApp       = 1024 // app 未登录
	CodeCommonNoneTenantId         = 1101 // 租户号为空
	CodeCommonErrorTenantId        = 1102 // 错误的租户号
	CodeCommonErrorServiceId       = 1103 // 错误的服务号
	CodeCommonRouteDisabled        = 1104 // 无效路由
)
