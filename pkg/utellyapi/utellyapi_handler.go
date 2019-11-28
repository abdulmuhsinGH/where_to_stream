package utellyapi

import (
	"fmt"
	"log"
	"net/http"
	"stream_available_on/pkg/format"
	"time"

	"github.com/gorilla/mux"
)

/*
Handlers define user
*/
type Handlers struct {
	logger *log.Logger
}

//var searchService Service

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

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
	format.Send(response, http.StatusOK, format.Message(true, "Progra Details", programDetails))

}

/*
SetupRoutes sets up routes to respective handlers
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
