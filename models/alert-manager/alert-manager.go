package alert_manager

type AlertManager struct {
	Alerts 		[] Alert
	Url 		string
}

func NewAlertManager(url string) *AlertManager {
	return &AlertManager{Url: url}
}
