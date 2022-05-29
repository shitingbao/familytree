package http

import (
	"fmt"
	"net/http"
	"path"

	"familytree/family/conf"

	"familytree/family/service"

	"github.com/gin-gonic/gin"
)

var (
	Svc                *service.Service
	vueAssetsRoutePath = "./dist" // dist 所在路径
)

//初始化路由
func Init(c *conf.Config, s *service.Service) {
	e := gin.Default()
	Svc = s
	setupInnerEngine(e)

	go e.Run(fmt.Sprintf("%s:%d", c.HTTPSvc.Host, c.HTTPSvc.Port))
}

func setupInnerEngine(e *gin.Engine) {

	e.StaticFile("/", path.Join(vueAssetsRoutePath, "index.html"))             // 指定资源文件 127.0.0.1/ 这种
	e.StaticFile("/favicon.ico", path.Join(vueAssetsRoutePath, "favicon.ico")) // 127.0.0.1/favicon.ico
	e.StaticFS("/assets", http.Dir(path.Join(vueAssetsRoutePath, "assets")))   // 以 assets 为前缀的 url

	g := e.Group("/v1").Use(origin)
	{
		g.POST("/member/create", FamilytreeCreate)
		g.POST("/member/update", FamilytreeUpdate)
		g.POST("/member/delete", FamilytreeDelete)
		g.POST("/member/list", FamilytreeList)
		g.POST("/member/last", FamilytreeLast)
		g.POST("/member/search", SearchMember)
	}
}

func origin(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")
}
