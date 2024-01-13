package cache

import (
	"github.com/goburrow/cache"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go-log-scanner/LogScanner/response"
	milog "hj_common/log"
	"sort"
	"time"
)

var GlobalCache = cache.New(
	cache.WithMaximumSize(100000),
	cache.WithExpireAfterWrite(90*time.Second),
)

type Handler func(ctx iris.Context) response.JsonResult

func HandleCache(handler Handler) context.Handler {
	irisHandlerFunc := func(ctx iris.Context) {
		cacheKey := GetCacheKey(ctx)
		rsp, b := GlobalCache.GetIfPresent(cacheKey)
		if b && rsp != nil {
			_, _ = ctx.JSON(rsp)
			milog.Debugf("use page cache : %v", cacheKey)
			return
		}

		ret := handler(ctx)
		GlobalCache.Put(cacheKey, ret)
		_, _ = ctx.JSON(ret)
		//milog.Debugf("new page cache : %v", cacheKey)
	}
	return irisHandlerFunc
}

func GetCacheKey(ctx iris.Context) string {
	cacheKey := ctx.RequestPath(true)
	mapData := ctx.URLParams()
	var keys []string
	for key, _ := range mapData {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, p1 := range keys {
		cacheKey += "_" + mapData[p1]
	}
	return cacheKey
}
