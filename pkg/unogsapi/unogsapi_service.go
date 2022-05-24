package unogsapi

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
	GetAvailability(string) (TitleResponse, error)
	UNOGSAdvanceSearch(string) error
}

/*
NfInfo is netflix info of program
*/
type TitleInfo struct {
	Title     string `json:"title"`
	Img       string `json:"img"`
	TitleType string `json:"title_type"`
	NetflixID int    `json:"netflix_id"`
	Synopsis  string `json:"synopsis"`
	Rating    string `json:"rating"`
	Year      string `json:"year"`
	Runtime   string `json:"runtime"`
	ImdbID    string `json:"imdb_id"`
	Poster    string `json:"poster"`
	Top250    int    `json:"top250"`
	Top250Tv  int    `json:"top250tv"`
	TitleDate string `json:"title_date"`
}

/*
Country details of program
*/
type Country struct {
	CountryCode  string `json:"country_code"`
	Country      string `json:"country"`
	SeasonDetail string `json:"season_detail"`
	ExpireDate   string `json:"expire_date"`
	NewDate      string `json:"new_date"`
	Audio        string `json:"audio"`
	Subtitle     string `json:"subtitle"`
}

type CountryResponse struct {
	Object struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	} `json:"object"`
	Results []Country `json:"results"`
}

type TitleDetails struct {
	Title         string `json:"title"`
	MaturityLabel string `json:"maturity_label"`
	MaturityLevel string `json:"maturity_level"`
	Synopsis      string `json:"synopsis"`
	TitleType     string `json:"title_type"`
	DefaultImage  string `json:"default_image"`
	LargeImage    string `json:"large_image"`
	NetflixID     string `json:"netflix_id"`
	StartDate     string `json:"start_date"`
	LatestDate    string `json:"latest_date"`
	Year          string `json:"year"`
	Poster        string `json:"poster"`
	Runtime       string `json:"runtime"`
	Awards        string `json:"awards"`
	OriginCountry string `json:"origin_country"`
	Rating        string `json:"rating"`
	AltID         string `json:"alt_id"`
	AltPlot       string `json:"alt_plot"`
	AltMetascore  string `json:"alt_metascore"`
	AltVotes      string `json:"alt_votes"`
	AltRuntime    string `json:"alt_runtime"`
	AltImage      string `json:"alt_image"`
}

type TitleInfoResponse struct {
	Object struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	} `json:"object"`
	Results []TitleInfo `json:"results"`
}

type TitleResponse struct {
	Details   TitleDetails `json:"details"`
	Countries []Country    `json:"countries"`
}

/*
GetNetflixDetails gets netflix details of a program from uNoGS api
*/
func GetNetflixDetails(netflixID string) (TitleDetails, error) {
	titleDetails := TitleDetails{}
	netflixID = url.PathEscape(netflixID)
	UNOGSBaseURL := os.Getenv("RAPI_API_UNOGS_HOST")
	unogsURL := fmt.Sprintf("https://%s/title/details?netflix_id=%s", UNOGSBaseURL, netflixID)

	body, err := sendUNOGSAPIRequest(unogsURL)
	if err != nil {
		return TitleDetails{}, err
	}

	err = json.Unmarshal(body, &titleDetails)
	if err != nil {
		return TitleDetails{}, err
	}
	return titleDetails, nil

}

/*
UNOGSAdvanceSearch gets netflix details of a program from uNoGS api using the title of a movie/show
*/
func TitleDetailsSearch(title string, skip int, limit int) (TitleInfoResponse, error) {

	var titleInfoResponse TitleInfoResponse
	title = url.PathEscape(title)
	//endYear := strconv.Itoa(time.Now().Year())

	UNOGSBaseURL := os.Getenv("RAPI_API_UNOGS_HOST")

	unogsAdvanceSearchURL := fmt.Sprintf("https://%s/search/titles??limit=%d&offset=%d&order_by=date&title=%s", UNOGSBaseURL, limit, skip, title)

	body, err := sendUNOGSAPIRequest(unogsAdvanceSearchURL)
	if err != nil {
		return TitleInfoResponse{}, err
	}

	err = json.Unmarshal(body, &titleInfoResponse)
	if err != nil {
		return TitleInfoResponse{}, err
	}

	return titleInfoResponse, nil

}

func sendUNOGSAPIRequest(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", os.Getenv("RAPI_API_UNOGS_HOST"))
	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_UNOGS_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func loadTitleDetails(id int) (TitleResponse, error) {
	titleResponse := TitleResponse{}
	countryResponse := CountryResponse{}
	UNOGSBaseURL := os.Getenv("RAPI_API_UNOGS_HOST")
	titleDetailsURL := fmt.Sprintf("https://%s/title/details?netflix_id=%d", UNOGSBaseURL, id)
	titleCountriesURL := fmt.Sprintf("https://%s/title/countries?netflix_id=%d", UNOGSBaseURL, id)

	body, err := sendUNOGSAPIRequest(titleDetailsURL)
	if err != nil {
		return TitleResponse{}, err
	}
	err = json.Unmarshal(body, &titleResponse.Details)
	if err != nil {
		return TitleResponse{}, err
	}

	countriesRes, err := sendUNOGSAPIRequest(titleCountriesURL)
	if err != nil {
		return TitleResponse{}, err
	}

	err = json.Unmarshal(countriesRes, &countryResponse)
	if err != nil {
		return TitleResponse{}, err
	}
	titleResponse.Countries = countryResponse.Results

	return titleResponse, nil

}

/*
GetNetflixDetailsForAdvanceSearchResults qwerty keyboard
*/
func GetNetflixTitleDetails(netflixID int) (TitleResponse, error) {

	result, err := loadTitleDetails(netflixID)
	if err != nil {
		return TitleResponse{}, err
	}

	return result, nil

}
