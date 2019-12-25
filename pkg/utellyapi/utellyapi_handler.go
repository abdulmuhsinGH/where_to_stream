package utellyapi

import (
	"fmt"
	"log"
	"net/http"
	"where_to_stream/pkg/format"
	"time"

	"github.com/gorilla/mux"
)

/*
Handlers with logger
*/
type Handlers struct {
	logger *log.Logger
}

/*
Logger handles logs
*/
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

/*
HandleSearch gets data from http request and sends to GetAvailability
*/
func (h *Handlers) HandleSearch(response http.ResponseWriter, request *http.Request) {
	//var program string
	var url = *request.URL
	program := url.Query().Get("program")
	fmt.Println(program)
	programDetails, err := GetAvailability(program)
	if err != nil {
		h.logger.Printf("Error getting program details: %v", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error getting program details", UtellyResponse{}))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Program Details", programDetails))

}

/*
SetupRoutes links routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/utelly/search", h.Logger(h.HandleSearch)).Methods("GET")
}

/*
NewHandlers initiates user handler
*/
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
