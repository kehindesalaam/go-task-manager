package main

import (
	"log"
	"net/http"
	"taskmanager/common"
	"taskmanager/routers"
)

func main() {
	common.StartUp()

	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
