package controllers

import (
	"center-information-alicia/config"
	alert_manager "center-information-alicia/models/alert-manager"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetAlerts(ctx *gin.Context, config config.Configuration) {

	var result [] alert_manager.Alert
	for _, am := range config.AlertManagers {
		alertManagerInstance := alert_manager.NewAlertManager(am.Url+"/api/v2/alerts")
		err := alertManagerInstance.GetAlerts()
		if err != nil {
			log.Warn("alertmanager gets info err: ", err)
		}
		result = append(result, alertManagerInstance.Alerts...)
	}
	ctx.JSON(http.StatusOK, result)
	return
}
