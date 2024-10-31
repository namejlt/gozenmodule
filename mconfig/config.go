package mconfig

// ===================================================
// ====================== 模式配置 ======================
// ===================================================

func GetConfigAccess() (data ConfigModeDataAccess) {
	l.RLock()
	data = accessConfig
	defer l.RUnlock()
	return
}

func GetConfigCommon() (data ConfigModeDataCommon) {
	l.RLock()
	data = commonConfig
	defer l.RUnlock()
	return
}

func GetConfigTouch() (data ConfigModeDataTouch) {
	l.RLock()
	data = touchConfig
	defer l.RUnlock()
	return
}

func GetConfigCore() (data ConfigModeDataCore) {
	l.RLock()
	data = coreConfig
	defer l.RUnlock()
	return
}

func GetConfigTouchDo() (data ConfigModeDataTouchDo) {
	l.RLock()
	data = touchDoConfig
	defer l.RUnlock()
	return
}

func GetConfigOneId() (data ConfigModeDataOneId) {
	l.RLock()
	data = oneIdConfig
	defer l.RUnlock()
	return
}

func GetConfigSaasTranspond() (data ConfigModeDataSaasTranspond) {
	l.RLock()
	data = saasTranspondConfig
	defer l.RUnlock()
	return
}

// ===================================================
// ====================== 租户配置 ======================
// ===================================================

func GetConfigTenantAccess(tenantId uint64) (data ConfigTenantDataAccess) {
	value, ok := accessConfigTenant.Load(tenantId)
	if !ok {
		// 不存在,返回默认值
		data = getDefaultConfigTenantAccess()
		return
	}

	data, _ = value.(ConfigTenantDataAccess)
	return
}

func GetConfigTenantCommon(tenantId uint64) (data ConfigTenantDataCommon) {
	value, ok := commonConfigTenant.Load(tenantId)
	if !ok {
		// 不存在,返回默认值
		data = getDefaultConfigTenantCommon()
		return
	}

	data, _ = value.(ConfigTenantDataCommon)
	return
}

func GetConfigTenantTouch(tenantId uint64) (data ConfigTenantDataTouch) {
	value, ok := touchConfigTenant.Load(tenantId)
	if !ok {
		// 不存在,返回默认值
		data = getDefaultConfigTenantTouch()
		return
	}

	data, _ = value.(ConfigTenantDataTouch)
	return
}

func GetConfigTenantCore(tenantId uint64) (data ConfigTenantDataCore) {
	value, ok := coreConfigTenant.Load(tenantId)
	if !ok {
		// 不存在,返回默认值
		data = getDefaultConfigTenantCore()
		return
	}

	data, _ = value.(ConfigTenantDataCore)
	return
}

func GetConfigTenantTouchDo(tenantId uint64) (data ConfigTenantDataTouchDo) {
	value, ok := touchDoConfigTenant.Load(tenantId)
	if !ok {
		// 不存在,返回默认值
		data = getDefaultConfigTenantTouchDo()
		return
	}

	data, _ = value.(ConfigTenantDataTouchDo)
	return
}

func GetConfigTenantOneId(tenantId uint64) (data ConfigTenantDataOneId) {
	value, ok := oneIdConfigTenant.Load(tenantId)
	if !ok {
		// 不存在,返回默认值
		data = getDefaultConfigTenantOneId()
		return
	}

	data, _ = value.(ConfigTenantDataOneId)
	return
}

func GetConfigTenantSaasTranspond(tenantId uint64) (data ConfigTenantDataSaasTranspond) {
	value, ok := saasTranspondConfigTenant.Load(tenantId)
	if !ok {
		// 不存在,返回默认值
		data = getDefaultConfigTenantSaasTranspond()
		return
	}

	data, _ = value.(ConfigTenantDataSaasTranspond)
	return
}
