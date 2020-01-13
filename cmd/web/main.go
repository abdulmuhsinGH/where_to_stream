package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("where_to_stream/web/dist"))
  	http.Handle("/", fs)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}
