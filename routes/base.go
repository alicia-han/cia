package routes

import (
	"center-information-alicia/controllers"
	"github.com/gin-gonic/gin"
)

func RouterLivenessCheck(router *gin.Engine) *gin.Engine {

	routerNew := router.Group("/health")
	routerNew.Use()
	{
		routerNew.GET("", func(context *gin.Context) {
			controllers.LivenessCheckOK(context)
		})
	}
	return router
}

func RouterBaseRoot(router *gin.Engine) * gin.Engine {
	routerNew := router.Group("/")
	routerNew.Use()
	{
		routerNew.GET("", func(context *gin.Context) {
			controllers.LivenessCheckOK(context)
		})
	}

	return router
}

