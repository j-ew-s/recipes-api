package main

import (
	"os"

	"github.com/j-ew-s/receipts-api/configs"
	httpserverhandler "github.com/j-ew-s/receipts-api/handlers/httpServer"
)

func main() {

	var envArg = setEnv()

	configs.Create(envArg)

	var httpServer = httpserverhandler.CreateHTTPServer()

	httpServer.ListenAndServe(configs.ServerConfig.APIPort)
}

func setEnv() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}
	return "dev"
}
