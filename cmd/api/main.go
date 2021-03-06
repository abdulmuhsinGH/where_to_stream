package main

import (
	"log"
	"os"
	"where_to_stream/pkg/server"
	"where_to_stream/pkg/unogsapi"
	"where_to_stream/pkg/utellyapi"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/* var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
) */

func main() {
	logger := log.New(os.Stdout, "stream_available_on ", log.LstdFlags|log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		logger.Fatalf(".env file not found %v", err)
	}

	utelly := utellyapi.NewHandlers(logger)
	unogs := unogsapi.NewHandlers(logger)

	router := mux.NewRouter()
	utelly.SetupRoutes(router)
	unogs.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := server.New(router, "localhost:"+port)
	logger.Println("server starting")
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
