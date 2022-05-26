package http

import (
	"fmt"

	"familytree/family/conf"

	"familytree/family/service"

	"github.com/gin-gonic/gin"
)

var (
	Svc *service.Service
)

//初始化路由
func Init(c *conf.Config, s *service.Service) {
	e := gin.Default()
	Svc = s
	setupInnerEngine(e)

	go e.Run(fmt.Sprintf("%s:%d", c.HTTPSvc.Host, c.HTTPSvc.Port))
}

func setupInnerEngine(e *gin.Engine) {
	g := e.Group("/v1").Use(origin)
	{
		g.POST("/member/create", FamilytreeCreate)
		g.POST("/member/update", FamilytreeUpdate)
		g.POST("/member/delete", FamilytreeDelete)
		g.POST("/member/list", FamilytreeList)
		g.POST("/member/last", FamilytreeLast)

	}
}

func origin(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")
}
