package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jhabc1314/applet_dinner/app/controllers/api"
	"github.com/jhabc1314/applet_dinner/app/middleware"
)

type ApiRoutes struct {
	prefix string //路由前缀
	domainPrefix string //二级域名前缀
	middleware map[string] middleware.Middleware
}

func NewApiRoutes() ApiRoutes {
	return ApiRoutes{}
}


func (h *ApiRoutes)SetRouteConfig()  {

}

func (h *ApiRoutes)GetRoutes(r *gin.Engine)  {
	
	r.GET("/search",api.Search)
	r.GET("/restaurant/:id/info", api.Restaurant)
}
