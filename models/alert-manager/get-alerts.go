package alert_manager

import (
	"center-information-alicia/utils/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (am * AlertManager) GetAlerts() error  {
	code, infoBytes, err := http.GetWithAuth(am.Url,"")
	if code < 200 || code >= 300 {
		log.Warn("code is:", code)
		return err
	}

	err = json.Unmarshal(infoBytes, &am.Alerts)
	if err != nil {
		return err
	}

	return nil
}
