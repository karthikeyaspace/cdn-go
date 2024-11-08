package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthikeyaspace/cdn-go/internal/config"
	"github.com/karthikeyaspace/cdn-go/internal/handler"
)

func main() {
	cfg := config.LoadConfig()

	router := mux.NewRouter()

	router.HandleFunc("/upload", handler.UploadHandler).Methods("POST")

	router.HandleFunc("/view/{id}", handler.ViewHandler).Methods("GET")

	server := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	log.Printf("Server running on http://localhost%v", cfg.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server, %v", err.Error())
	}
}
