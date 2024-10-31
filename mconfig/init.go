package mconfig

import (
	"github.com/namejlt/gozen"
	"math/rand"
	"sync"
	"time"
)

var (
	// ======================= 私有化部署配置 - 跨租户
	l                   sync.RWMutex                // 控制全局读写
	commonConfig        ConfigModeDataCommon        // common 配置
	accessConfig        ConfigModeDataAccess        // access 配置
	touchConfig         ConfigModeDataTouch         // touch 配置
	coreConfig          ConfigModeDataCore          // core 配置
	touchDoConfig       ConfigModeDataTouchDo       // touchdo 配置
	oneIdConfig         ConfigModeDataOneId         // oneid 配置
	saasTranspondConfig ConfigModeDataSaasTranspond // saasTranspond 配置

	// ======================= 租户配置
	commonConfigTenant        sync.Map // common配置
	accessConfigTenant        sync.Map // access配置
	touchConfigTenant         sync.Map // touch配置
	coreConfigTenant          sync.Map // core配置
	touchDoConfigTenant       sync.Map // touch_do配置
	oneIdConfigTenant         sync.Map // one_id配置
	saasTranspondConfigTenant sync.Map // saastranspond配置
)

const (
	syncTime = 5 * time.Minute
)

func setConfig(tag string, data ConfigModeData) {
	ret, err := data.Decode(tag)
	if err != nil {
		gozen.LogErrorw(gozen.LogNameLogic, "mconfig setConfig",
			gozen.LogKNameCommonErr, err)
		return
	}

	switch tag {
	case ConfigModeTagAccess:
		if tmp, ok := ret.(ConfigModeDataAccess); ok {
			accessConfig = tmp
		}
	case ConfigModeTagCommon:
		if tmp, ok := ret.(ConfigModeDataCommon); ok {
			commonConfig = tmp
		}
	case ConfigModeTagTouch:
		if tmp, ok := ret.(ConfigModeDataTouch); ok {
			touchConfig = tmp
		}
	case ConfigModeTagCore:
		if tmp, ok := ret.(ConfigModeDataCore); ok {
			coreConfig = tmp
		}
	case ConfigModeTagTouchDo:
		if tmp, ok := ret.(ConfigModeDataTouchDo); ok {
			touchDoConfig = tmp
		}
	case ConfigModeTagOneId:
		if tmp, ok := ret.(ConfigModeDataOneId); ok {
			oneIdConfig = tmp
		}
	case ConfigModeTagSaasTranspond:
		if tmp, ok := ret.(ConfigModeDataSaasTranspond); ok {
			saasTranspondConfig = tmp
		}
	}
	return
}

var (
	initConfigTagM sync.Map
)

func InitConfig(tag string) {
	if _, ok := initConfigTagM.LoadOrStore(tag, struct{}{}); !ok { //未配置 则初始化
		apiToLocal(tag, true)
		//周期更新配置
		go func() { //强制退出
			rSecond := time.Duration(rand.Intn(60))
			t := time.NewTicker(syncTime + rSecond*time.Second)
			for {
				select {
				case <-t.C:
					apiToLocal(tag, false)
				}
			}
		}()
	}
}

func apiToLocal(tag string, panicBool bool) {
	l.Lock()
	defer l.Unlock()
	//初始加载失败则panic
	ret, err := configModeData(tag)
	if panicBool {
		if err != nil {
			panic("初始加载配置失败" + tag)
			return
		}
	} else {
		if err != nil {
			gozen.LogErrorw(gozen.LogNameApi, "mconfig apiToLocal",
				gozen.LogKNameCommonErr, err)
			return
		}
	}
	setConfig(tag, ret)
}

func setConfigTenant(tag string, data []ConfigTenantInfo) {
	tenantIdM := make(map[uint64]interface{})
	for _, v := range data {
		tenantIdM[v.TenantId] = struct{}{}
		ret, err := v.Config.Decode(tag)
		if err != nil {
			gozen.LogErrorw(gozen.LogNameLogic, "mconfig setConfigTenant",
				gozen.LogKNameCommonData, v,
				gozen.LogKNameCommonErr, err)
			return
		}
		switch tag {
		case ConfigModeTagAccess:
			if tmp, ok := ret.(ConfigTenantDataAccess); ok {
				accessConfigTenant.Store(v.TenantId, tmp)
			}
		case ConfigModeTagCommon:
			if tmp, ok := ret.(ConfigTenantDataCommon); ok {
				commonConfigTenant.Store(v.TenantId, tmp)
			}
		case ConfigModeTagTouch:
			if tmp, ok := ret.(ConfigTenantDataTouch); ok {
				touchConfigTenant.Store(v.TenantId, tmp)
			}
		case ConfigModeTagCore:
			if tmp, ok := ret.(ConfigTenantDataCore); ok {
				coreConfigTenant.Store(v.TenantId, tmp)
			}
		case ConfigModeTagTouchDo:
			if tmp, ok := ret.(ConfigTenantDataTouchDo); ok {
				touchDoConfigTenant.Store(v.TenantId, tmp)
			}
		case ConfigModeTagOneId:
			if tmp, ok := ret.(ConfigTenantDataOneId); ok {
				oneIdConfigTenant.Store(v.TenantId, tmp)
			}
		case ConfigModeTagSaasTranspond:
			if tmp, ok := ret.(ConfigTenantDataSaasTranspond); ok {
				saasTranspondConfigTenant.Store(v.TenantId, tmp)
			}
		}
	}

	// 删除多余的数据
	switch tag {
	case ConfigModeTagAccess:
		accessConfigTenant.Range(func(key, value any) bool {
			tenantId, ok := key.(uint64)
			if ok {
				if _, exist := tenantIdM[tenantId]; !exist {
					// 需要删除
					accessConfigTenant.Delete(key)
				}
			}
			return true
		})
	case ConfigModeTagCommon:
		commonConfigTenant.Range(func(key, value any) bool {
			tenantId, ok := key.(uint64)
			if ok {
				if _, exist := tenantIdM[tenantId]; !exist {
					// 需要删除
					commonConfigTenant.Delete(key)
				}
			}
			return true
		})
	case ConfigModeTagTouch:
		touchConfigTenant.Range(func(key, value any) bool {
			tenantId, ok := key.(uint64)
			if ok {
				if _, exist := tenantIdM[tenantId]; !exist {
					// 需要删除
					touchConfigTenant.Delete(key)
				}
			}
			return true
		})
	case ConfigModeTagCore:
		coreConfigTenant.Range(func(key, value any) bool {
			tenantId, ok := key.(uint64)
			if ok {
				if _, exist := tenantIdM[tenantId]; !exist {
					// 需要删除
					coreConfigTenant.Delete(key)
				}
			}
			return true
		})
	case ConfigModeTagTouchDo:
		touchDoConfigTenant.Range(func(key, value any) bool {
			tenantId, ok := key.(uint64)
			if ok {
				if _, exist := tenantIdM[tenantId]; !exist {
					// 需要删除
					touchDoConfigTenant.Delete(key)
				}
			}
			return true
		})
	case ConfigModeTagOneId:
		oneIdConfigTenant.Range(func(key, value any) bool {
			tenantId, ok := key.(uint64)
			if ok {
				if _, exist := tenantIdM[tenantId]; !exist {
					// 需要删除
					oneIdConfigTenant.Delete(key)
				}
			}
			return true
		})
	case ConfigModeTagSaasTranspond:
		saasTranspondConfigTenant.Range(func(key, value any) bool {
			tenantId, ok := key.(uint64)
			if ok {
				if _, exist := tenantIdM[tenantId]; !exist {
					// 需要删除
					saasTranspondConfigTenant.Delete(key)
				}
			}
			return true
		})

	}
	return
}

var (
	initTenantConfigTagM sync.Map
)

func InitTenantConfig(tag string) {
	if _, ok := initTenantConfigTagM.LoadOrStore(tag, struct{}{}); !ok { //未配置 则初始化
		apiTenantConfigListToLocal(tag, true)
		//周期更新配置
		go func() { //强制退出
			rSecond := time.Duration(rand.Intn(60))
			t := time.NewTicker(syncTime + rSecond*time.Second)
			for {
				select {
				case <-t.C:
					apiTenantConfigListToLocal(tag, false)
				}
			}
		}()
	}
}

func apiTenantConfigListToLocal(tag string, panicBool bool) {
	//初始加载失败则panic
	ret, err := configTenantList(tag)
	if panicBool {
		if err != nil {
			panic("租户配置初始加载配置失败" + tag)
			return
		}
	} else {
		if err != nil {
			gozen.LogErrorw(gozen.LogNameApi, "mconfig apiTenantConfigListToLocal",
				gozen.LogKNameCommonErr, err)
			return
		}
	}
	setConfigTenant(tag, ret)
}
