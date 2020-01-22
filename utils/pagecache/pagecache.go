package pagecache

import (
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var (
	BasePath  string              = ""
	ExpireSec int64               = 0
	store     *cache.FileCache    = nil
	cacheMap  map[string]bool     = nil
	paramMap  map[string][]string = nil
)

func InitCache() {
	store = &cache.FileCache{CachePath: BasePath}
	pagecacheList := beego.AppConfig.Strings("pagecache_list")

	//初始化静态化配置列表
	cacheMap = make(map[string]bool)
	for _, v := range pagecacheList {
		cacheMap[strings.ToLower(v)] = true
	}

	paramMap = make(map[string][]string)
	pagecacheMap, _ := beego.AppConfig.GetSection("pagecache_param")
	for k, v := range pagecacheMap {
		sv := strings.Split(v, ";")
		paramMap[k] = sv
	}
}

func InCacheList(controllerName, actionName string) bool {
	keyname := cacheKey(controllerName, actionName)
	if f := cacheMap[keyname]; f {
		return f
	}
	return false
}

func NeedWrite(controllerName, actionName string, params map[string]string) bool {
	if InCacheList(controllerName, actionName) {
		keyname := cacheKey(controllerName, actionName, params)
		if len(store.Get(keyname).(string)) > 0 {
			return false
		} else {
			beego.Debug("need write :" + keyname)
			return true
		}
	}
	return false
}

func Write(controllerName, actionName string, content *string, params map[string]string) error {
	keyname := cacheKey(controllerName, actionName, params)
	if len(keyname) == 0 {
		return errors.New("未找到缓存key")
	}

	err := store.Put(keyname, *content, time.Duration(ExpireSec)*time.Second)

	return err
}

func Read(controllerName, actionName string, params map[string]string) (*string, error) {
	keyname := cacheKey(controllerName, actionName, params)
	if len(keyname) == 0 {
		return nil, errors.New("未找到缓存key")
	}

	content := store.Get(keyname).(string)

	return &content, nil
}

func cacheKey(controllerName, actionName string, paramArray ...map[string]string) string {
	if len(controllerName) > 0 && len(actionName) > 0 {
		rtnstr := strings.ToLower(controllerName + "_" + actionName)
		if len(paramArray) > 0 {
			for _, v := range paramMap[rtnstr] {
				rtnstr = rtnstr + "_" + strings.ReplaceAll(v, ":", "") + "_" + paramArray[0][v]
			}
		}
		return rtnstr
	}

	return ""
}

func ClearExpiredFiles() {
	for k, _ := range cacheMap {
		if store.IsExist(k) {
			content := store.Get(k).(string)
			if len(content) == 0 {
				store.Delete(k)
			}
		}
	}
}
