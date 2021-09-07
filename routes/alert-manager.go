package routes

import (
	"center-information-alicia/config"
	"center-information-alicia/controllers"
	"github.com/gin-gonic/gin"
)

func RouterAlertManager(router *gin.Engine, config config.Configuration) * gin.Engine {
	routerNew := router.Group("/alert-manager")
	routerNew.Use()
	{
		routerNew.GET("", func(context *gin.Context) {
			controllers.GetAlerts(context,config)
		})
	}

	return router
}

