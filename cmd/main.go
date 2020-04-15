package main

import (
	"os"

	"github.com/j-ew-s/recipes-api/configs"
	httpserverhandler "github.com/j-ew-s/recipes-api/handlers/httpServer"
)

func main() {

	var envArg = setEnv()

	configs.Create(envArg)

	var httpServer = httpserverhandler.CreateHTTPServer()

	httpServer.ListenAndServe(configs.ServerConfig.APIPort)
}

func setEnv() string {

	var env = os.Getenv("ENV")

	if len(env) <= 0 {
		env = "dev"
	}

	return env
}
