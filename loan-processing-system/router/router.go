package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	apiconstants "github.com/suhelz/loan-processing-system/constants/api-constants"
	"github.com/suhelz/loan-processing-system/model"
)

type Router struct {
	Endpoint string
	router   *gin.Engine
}

func NewRouter(config model.Config) *Router {
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Continer Status: OK",
		})
	})

	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set(apiconstants.AccessControlHeaderKey, apiconstants.AccessControlHeaderValue)
		ctx.Writer.Header().Set(apiconstants.AccessControlMethodKey, apiconstants.AccessControlHeaderValue)
		// Handle CORS
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			ctx.Next()
		}
	})

	endpoint := fmt.Sprintf(":%d", config.Service.Port)
	return &Router{
		Endpoint: endpoint,
		router:   r,
	}
}

func (r Router) Run() error {
	return r.router.Run(r.Endpoint)
}
