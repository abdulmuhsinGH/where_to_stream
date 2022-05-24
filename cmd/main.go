package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	spa := spaHandler{staticPath: "web/dist", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := server.New(router, "localhost:"+port)
	logger.Println("server starting :" + port)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
