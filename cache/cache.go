package cache

import (
	"github.com/gogf/gf/container/gmap"
	structModel "modules/struct"
	"reflect"
)

type cacheStruct struct {
	cache *gmap.Map
}

var ModelCache *cacheStruct

func init() {
	ModelCache = &cacheStruct{
		cache: gmap.New(true),
	}
}

func (c *cacheStruct) AddSlice(key interface{},structName reflect.Type, value interface{}) interface{} {
	if value == nil {
		return nil
	}
	if _, ok := value.(structModel.A); ok {
		originValue, ok := c.cache.Get(key).([]structModel.A)
		if !ok {
			originValue = make([]structModel.A, 0)
		}
		originValue = append(originValue, value.(structModel.A))
		return originValue
	} else if _, ok := value.(structModel.B); ok {
		originValue, ok := c.cache.Get(key).([]structModel.B)
		if !ok {
			originValue = make([]structModel.B, 0)
		}
		originValue = append(originValue, value.(structModel.B))
		return originValue
	} else if _, ok := value.(structModel.C); ok {
		originValue, ok := c.cache.Get(key).([]structModel.C)
		if !ok {
			originValue = make([]structModel.C, 0)
		}
		originValue = append(originValue, value.(structModel.C))
		return originValue
	}
	return nil
}
