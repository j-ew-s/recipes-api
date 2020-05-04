package httpserverhandler

import (
	"github.com/buaazp/fasthttprouter"
	recipescontroller "github.com/j-ew-s/recipes-api/api"
	"github.com/j-ew-s/recipes-api/configs"
	"github.com/valyala/fasthttp"
)

// CreateHTTPServer Creates new Server using fasthttp
// Returns a instance from fasthttp Server
func CreateHTTPServer() {

	router := fasthttprouter.New()

	setRoutes(router)

	fasthttp.ListenAndServe(configs.ServerConfig.APIPort, CORS(router.Handler))

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

var (
	corsAllowHeaders     = "*"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

// CORS handles CORS
func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}
