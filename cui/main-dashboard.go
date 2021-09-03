package cui

import (
	"center-information-alicia/config"
	ui "github.com/gizak/termui/v3"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

func CuiEntry(configPath string) {

	f, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Config File Opened Field: ", err)
	}
	var configuration config.Configuration
	err = yaml.Unmarshal(f, &configuration)
	if err != nil {
		log.Fatal("unmarshal failed : ", err)
	}

	log.Debug("read configuration:", configuration)

	// Initialize UI
	if err := ui.Init(); err != nil {
		log.Fatal("Error initializing the UI: ", err)
	}
	defer ui.Close()

	// generate a grid takes all the terminal screen
	grid := ui.NewGrid()
	termWidth, termHight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHight)

	dataDogList := GetListDatadog()
	pingDomList := GetListPingdom()
	alertManagerList := GetListAlertManagers(configuration)

	//Guid Configuration
	grid.Set(
		ui.NewCol(1.0/2,
			ui.NewRow(1.0/2, dataDogList),
			ui.NewRow(1.0/2, pingDomList),
		),
		ui.NewCol(1.0/2, alertManagerList),
	)

	ui.Render(grid)

	update := func() {
		dataDogList = GetListDatadog()
		pingDomList = GetListPingdom()
		alertManagerList = GetListAlertManagers(configuration)
		ui.Render(dataDogList,pingDomList,alertManagerList,grid)
	}

	ticker := time.NewTicker(time.Second * 5).C
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			update()
		}
	}
}