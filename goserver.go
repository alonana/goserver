package main

import (
	"github.com/alonana/goserver/logging"
	"github.com/alonana/goserver/webserver"
)

func main() {
	logging.AppLoggerInit("goserver.log")
	logging.AppLogger.Info("Starting")
	webserver.Start()
}
