package cui

import (
	"center-information-alicia/models/Pingdom"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"os"
)

func GetListPingdom()  *widgets.List{

	l := widgets.NewList()
	l.Title = "Pingdom Monitor"

	//l.SetRect(0,0,100, 20)
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = true
	l.Rows = getPingdomStr()
	return l

}

func getPingdomStr () (result []string) {
	pa := Pingdom.NewPingdomAgent("https://api.pingdom.com/api/3.1/checks", os.Getenv("CIA_PINGDOM_TOKEN"))
	err := pa.GetChecks()
	if err != nil {
		return []string{
			err.Error(),
		}
	}

	result = append(result, "----------------Down Sites---------------")
	for _, i := range pa.Checks().Checks {
		if i.Status != "up"{
			result = append(result, i.Name + ":", i.Status)
		}
	}

	result = append(result, "-----------------------------------------")
	result = append(result, fmt.Sprintf("Total: %d, Limited: %d, Filtered: %d",
		pa.Checks().Counts.Total,
		pa.Checks().Counts.Limited,
		pa.Checks().Counts.Filtered))

	return
}
