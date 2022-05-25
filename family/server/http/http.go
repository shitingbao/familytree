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

	g := e.Group("/v1")
	{

		g.POST("/detail/create", DyProjectCaseCreate)
		g.GET("/detail/list", DyProjectCaseCreate)

	}

}
