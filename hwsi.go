package main

import (
	log "github.com/sirupsen/logrus"
	"hwsi/config"
	"hwsi/handler"
	"net/http"
	"time"
)

func main() {
	if err := config.Data.Get(); err != nil {
		log.Error(err)
	}
	handler.Init()
	server := http.Server{
		Addr:              config.Data.Server.Addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Second,
	}
	log.Println("ListenAndServe:", config.Data.Server.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
