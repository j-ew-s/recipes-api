package main

import (
	httpserverhandler "github.com/j-ew-s/receipts-api/handlers/httpServer"
)

func main() {
	var httpServer = httpserverhandler.CreateHTTPServer()

	/*
		TODO : the port should be replaced by configuration parameter
	*/
	httpServer.ListenAndServe(":9098")
}
