package utellyapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Service provides search adding operations.
type Service interface {
	GetAvailability(string) (UtellyResponse, error)
}

/*
Location describes the streaming location of a program
*/
type Location struct {
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	ID          string `json:"id"`
	Icon        string `json:"icon"`
}

/*
Result describes the response result for a program search
*/
type Result struct {
	Name      string `json:"name"`
	Weight    int    `json:"weight"`
	ID        string `json:"id"`
	Picture   string `json:"picture"`
	Locations []Location
}

// UtellyResponse desribes the responses body from utelly
type UtellyResponse struct {
	Variant    string   `json:"variant"`
	StatusCode int      `json:"status_code"`
	Term       string   `json:"term"`
	Results    []Result `json:"results"`
	Updated    string   `json:"updated"`
}

/*
GetAvailability searches the utelly api for show availability
*/
func GetAvailability(program string) (UtellyResponse, error) {
	utellyRes := UtellyResponse{}

	program = url.PathEscape(program)
	utellyURL := fmt.Sprintf("https://utelly-tv-shows-and-movies-availability-v1.p.rapidapi.com/lookup?term=%s", program)

	body, err := sendUtellyAPIRequest(utellyURL)
	if err != nil {
		return UtellyResponse{}, err
	}
	err = json.Unmarshal(body, &utellyRes)
	if err != nil {
		return UtellyResponse{}, err
	}
	return utellyRes, nil

}

func sendUtellyAPIRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPI_API_UTELLY_HOST"))
	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_UTELLY_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("e")
		return nil, err
	}

	return body, nil
}
