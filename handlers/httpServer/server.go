package httpserverhandler

import (
	"github.com/buaazp/fasthttprouter"
	recipescontroller "github.com/j-ew-s/recipes-api/api"
	"github.com/j-ew-s/recipes-api/configs"
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

	router.GET("/ping/", recipescontroller.Ping)
	router.POST("/recipes/", recipescontroller.Create)
	router.GET("/recipes/", recipescontroller.Get)
	router.GET("/recipes/tags/", recipescontroller.GetByTags)
	router.GET("/recipes/id/:id", recipescontroller.GetByID)
	router.PUT("/recipes/:id", recipescontroller.Put)
	router.DELETE("/recipes/:id", recipescontroller.Delete)
	/*router.GET("/recipes/terms/:search", controller.Search)*/
}
