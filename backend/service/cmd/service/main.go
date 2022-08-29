package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"gettingLogs/internal/clients"
	logHandle "gettingLogs/internal/http/log"
	"gettingLogs/internal/logger"
	queue "gettingLogs/internal/queue/log"
	logRepository "gettingLogs/internal/repository/log"
	logService "gettingLogs/internal/service/log"
)

func main() {
	logger := logger.New()
	db := clients.New(logger)
	logRepository := logRepository.New(db, logger)
	logService := logService.New(logRepository, logger)
	queue := queue.New(logService, logger)
	logHandle := logHandle.New(logService, logger, queue)
	go queue.TakeFromTheQueue()
	router := mux.NewRouter()
	router.HandleFunc("/log", logHandle.Add).Methods(http.MethodPost)
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
