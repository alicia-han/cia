package alert_manager

import "testing"

func TestAlertManager_GetAlerts(t *testing.T) {
	am := NewAlertManager("https://alertmanager.staging-fr-onroad.com/api/v2/alerts")
	am.GetAlerts()
}
