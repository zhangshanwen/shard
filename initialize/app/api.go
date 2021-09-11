package app

import (
	"fmt"
	"strings"

	"github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
)

var R = &gin.Engine{}

func init() {
	gin.ForceConsoleColor()
	R = gin.Default()
}

func CheckUrl(Method, Path string) (isExited bool) {
	// `backend` and `api` is in one project,so need to adjust router belong `backend`
	prefix := common.Backlash + common.BackendPrefix
	if !strings.HasPrefix(Path, prefix) {
		return
	}
	for _, route := range R.Routes() {
		if route.Method == Method && route.Path == Path {
			return true
		}
	}
	return
}

func InitRoute() {
	var routers []model.Route
	db.G.Find(&routers)
	routerMap := map[string]int64{}
	var dbRouters = mapset.NewSet()
	var newRouters = mapset.NewSet()
	for _, router := range routers {
		fullRouter := fmt.Sprintf("%s%s%s", router.Method, common.RouteSeparator, router.Path)
		dbRouters.Add(fullRouter)
		routerMap[fullRouter] = router.Id
	}
	for _, router := range R.Routes() {
		prefix := common.Backlash + common.BackendPrefix
		if !strings.HasPrefix(router.Path, prefix) {
			continue
		}
		newRouters.Add(fmt.Sprintf("%s||%s", router.Method, router.Path))
	}
	var delIds []int64
	for _, v := range dbRouters.Difference(newRouters).ToSlice() {
		fullRoute := v.(string)
		delIds = append(delIds, routerMap[fullRoute])
	}
	if len(delIds) > 0 {
		db.G.Delete(model.Route{}, delIds)
	}
	var createRouters []model.Route
	for _, v := range newRouters.Difference(dbRouters).ToSlice() {
		fullRoute := v.(string)
		split := strings.Split(fullRoute, common.RouteSeparator)
		createRouters = append(createRouters, model.Route{
			Method: split[0],
			Path:   split[1],
		})
	}
	if len(createRouters) > 0 {
		db.G.Create(createRouters)
	}
}
