package cui

import (
	"center-information-alicia/models/datadog"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func GetListDatadog()  *widgets.List{

	l := widgets.NewList()
	l.Title = "DataDog Monitor"

	//l.SetRect(0,0,100, 20)
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = true
	l.Rows = datadog.GetActiveAlertsFromDataDog()
	return l

}
