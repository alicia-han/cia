package server

import (
	"center-information-alicia/config"
	"center-information-alicia/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Start(config config.Configuration)  {
	router := gin.Default()
	router = routes.RouterBaseRoot(router)
	router = routes.RouterLivenessCheck(router)
	router = routes.RouterAlertManager(router, config)

	server := http.Server{
		Addr:              "0.0.0.0:"+config.Server.Port,
		Handler:           router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("[Error] Running Server", err)
	}
}

