package unogsapi

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"where_to_stream/pkg/format"

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
		defer h.logger.Printf("request processed in %s\n", time.Since(startTime))
		next(w, r)
	}
}

/*
HandleSearchByNetflixID gets movies/tv shows from http request and sends to GetAvailability
*/
func (h *Handlers) HandleSearchByNetflixID(response http.ResponseWriter, request *http.Request) {
	//var program string
	var url = *request.URL
	netflixID := url.Query().Get("netflix_id")

	ID, err := strconv.Atoi(netflixID)
	if err != nil {
		h.logger.Printf("Error getting program details: %v", err.Error())
		format.Send(response, http.StatusBadRequest, format.Message(false, "Error getting program details", nil))
		return
	}

	searchDetails, err := GetNetflixTitleDetails(ID)
	if err != nil {
		h.logger.Printf("Error getting program details: %v", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error getting program details", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Program Details", searchDetails))

}

/*
HandleSearchByTitle searches for movies/tv shows by name
*/
func (h *Handlers) HandleSearchByTitle(response http.ResponseWriter, request *http.Request) {
	//var program string
	var url = *request.URL
	var err error
	title := url.Query().Get("title")
	skip, _ := strconv.Atoi(url.Query().Get("skip"))
	limit, _ := strconv.Atoi(url.Query().Get("limit"))

	searchResults, err := TitleDetailsSearch(title, skip, limit)
	if err != nil {
		h.logger.Printf("Error getting program details: %v", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error getting program details", nil))
		return
	}

	format.Send(response, http.StatusOK, format.Message(true, "Program Details", searchResults))

}

/*
SetupRoutes links routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/unogs/title/details", h.Logger(h.HandleSearchByNetflixID)).Methods("GET")
	mux.HandleFunc("/api/unogs/search", h.Logger(h.HandleSearchByTitle)).Methods("GET")
}

/*
NewHandlers initiates user handler
*/
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
