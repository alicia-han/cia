package controllers

import (
	"github.com/gin-gonic/gin"
)

func LivenessCheckOK (context *gin.Context) {
	context.JSON(200,gin.H{
		"Message": "OK",
	})
}
