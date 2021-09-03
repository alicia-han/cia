package Pingdom

import (
	"center-information-alicia/utils/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)


func (pa * PingdomAgent) GetChecks() error {

	code, bodyBytes, err := http.GetWithAuth(pa.url, pa.apiToken)
	if code < 200 || code >= 300 || err != nil{
		log.Warn("Return Code: ", code, " Err:", err)
		return err
	}

	err = json.Unmarshal(bodyBytes, &pa.checks)
	if err != nil {
		return err
	}

	return nil
}