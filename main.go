package main

import (
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
	"user-generator/infra/database"
	"user-generator/infra/server"
)

func main() {

	// Starts server application
	go server.StartServer()

	// starts meta application
	go server.StartMetaServer()

	// run database migrations
	go database.MySqlMigrations()

	// Listen for system signals to gracefully stop the application
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		log.Info("Received SIGINT, stopping...", nil)
	case syscall.SIGTERM:
		log.Info("Received SIGTERM, stopping...", nil)
	}
}
