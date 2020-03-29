package httpserverhandler

import (
	"github.com/buaazp/fasthttprouter"
	receiptscontroller "github.com/j-ew-s/receipts-api/api"
	"github.com/j-ew-s/receipts-api/configs"
	"github.com/valyala/fasthttp"
)

// CreateHTTPServer Creates new Server using fasthttp
// Returns a instance from fasthttp Server
func CreateHTTPServer() *fasthttp.Server {

	router := fasthttprouter.New()

	setRoutes(router)

	http := &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: configs.ServerConfig.MaxRequestBodySize,
	}

	return http
}

// Configure all routes for controllers
func setRoutes(router *fasthttprouter.Router) {

	router.POST("/receipts/", receiptscontroller.Create)
	router.GET("/receipts/", receiptscontroller.Get)
	router.GET("/receipts/tags/", receiptscontroller.GetByTags)
	router.GET("/receipts/id/:id", receiptscontroller.GetByID)
	router.PUT("/receipts/:id", receiptscontroller.Put)
	router.DELETE("/receipts/:id", receiptscontroller.Delete)
	/*router.GET("/receipts/terms/:search", controller.Search)*/
}
