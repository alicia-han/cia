package main

import (
	"center-information-alicia/cui"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

func setLog(){
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
	})

	logFile, err := os.OpenFile("./cia.log", os.O_CREATE| os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to create log file: ",err)
	}
	//log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	log.SetOutput(logFile)
	log.RegisterExitHandler(func() {
		if logFile == nil {
			return
		}
		logFile.Close()
	})

	switch os.Getenv("CIA_LOG_LEVEL") {
	case "warning","WARN","WARNING", "Warning":
		log.SetLevel(log.WarnLevel)
	case "debug","DEBUG","Debug":
		log.SetLevel(log.DebugLevel)
	case "trace","TRACE","Trace":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}



func main() {
	setLog()
	mainApp := &cli.App{
		Name:                   "center-info-alicia",
		Usage:                  "cia",
		Version:                "1.0.0",
		Commands:               [] *cli.Command{
			{
				Name: "cui",
				Usage: "launch CUI for monitoring",
				Action: func(c *cli.Context) error {
					cui.CuiEntry(c.String("config"))
					return nil
				},
			},
		},
		Flags:                  []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{
					"c",
				},
				Usage:       "load configuration file, default is ~/cia.yaml",
				Required:    false,
				Value:       os.Getenv("HOME")+"/cia.yaml",
			},
		},
		EnableBashCompletion:   true,
		BashComplete:           nil,
		Action:                 nil,
		CommandNotFound:        nil,
	}

	err := mainApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
