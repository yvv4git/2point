package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"2point/app/infrastructure/components"
	"2point/app/infrastructure/services"
	"2point/app/interface/router"
)

func main() {
	// Graceful shutdown.
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	container := services.NewApp()
	logger := container.Get("logger").(*log.Logger)
	config := container.Get("config").(*components.Config)
	mux := router.GetMux(container)

	logger.Println("App started")
	logger.Println("App port: ", config.GetPort())

	server := http.Server{Handler: mux, Addr: config.GenerateServerAddress()}

	// Start server in gorutine.
	go func() {
		logger.Println(server.ListenAndServe())
	}()

	// Gracefull exit.
	<-exit
	logger.Println("Stopping app...")
	//container.DeleteWithSubContainers()

	logger.Println("App stopped")
}
