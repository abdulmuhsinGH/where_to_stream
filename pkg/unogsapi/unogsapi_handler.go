package unogsapi

import (
	"fmt"
	"log"
	"net/http"
	"where_to_stream/pkg/format"
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
	programID := url.Query().Get("program_id")
	fmt.Println(programID)
	programDetails, err := GetNetflixDetails(programID)
	if err != nil {
		h.logger.Printf("Error getting program details: %v", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error getting program details", UNOGSResponse{}))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Program Details", programDetails.RESULT))

}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/unogs/search", h.Logger(h.HandleSearch)).Methods("GET")
}

/*
NewHandlers initiates user handler
*/
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
