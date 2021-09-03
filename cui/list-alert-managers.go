package cui

import (
	"center-information-alicia/config"
	alert_manager "center-information-alicia/models/alert-manager"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	log "github.com/sirupsen/logrus"
)

func GetListAlertManagers(configuration config.Configuration)  *widgets.List{

	l := widgets.NewList()
	l.Title = "Alert Manager Monitor"

	//l.SetRect(0,0,100, 20)
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = true
	l.Rows = getAlertManagerStr(configuration)
	return l

}

func getAlertManagerStr (configuration config.Configuration) (result []string) {
	Counter := 1

	result = append(result, "---------------------------------------------")
	for _, am := range configuration.AlertManagers {
		alertManagerInstance := alert_manager.NewAlertManager(am.Url+"/api/v2/alerts")
		err := alertManagerInstance.GetAlerts()
		if err != nil {
			log.Warn("alertmanager gets info err: ", err)
			result = append(result, "Error Reaching "+am.Name)
		}
		for _, a := range alertManagerInstance.Alerts {
			result = append (result, fmt.Sprintf("[<%d>](fg:red) [<%s>](fg:red) [<%s>](fg:red) %s",
				Counter, am.Name, a.Status.State, a.Annotations.Description))
			Counter ++
		}
		result = append(result, "---------------------------------------------")

	}
	return
}
