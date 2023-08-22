package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suhelz/loan-processing-system/controller"
	"github.com/suhelz/loan-processing-system/model"
)

func CreateApplicationRouter(config model.Config, controller controller.Controller) *gin.Engine {
	router := gin.Default()
	// TODO: add all controller to router
	return router
}
