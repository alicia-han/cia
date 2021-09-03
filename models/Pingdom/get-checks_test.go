package Pingdom

import (
	"os"
	"testing"
)

func TestPingdomAgent_GetChecks(t *testing.T) {
	pa := NewPingdomAgent("https://api.pingdom.com/api/3.1/checks", os.Getenv("CIA_PINGDOM_TOKEN"))

	err := pa.GetChecks()
	if err != nil {
		t.Log("err:", err)
		t.Fail()
	}
	t.Log("Checks:", pa.Checks())
}
