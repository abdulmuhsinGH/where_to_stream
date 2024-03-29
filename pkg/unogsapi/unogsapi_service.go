package unogsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

// Service provides search adding operations.
type Service interface {
	GetAvailability(string) (UNOGSResponse, error)
	UNOGSAdvanceSearch(string) error
}

/*
NfInfo is netflix info of program
*/
type nfInfo struct {
	Image1        string `json:"image1" json:"image"`
	Image2        string `json:"image2" json:"large_image"`
	Image         string `json:"image"`
	LargeImage    string `json:"largeimage"`
	Title         string `json:"title"`
	Synopsis      string `json:"synopsis"`
	MatLevel      string `json:"matlevel"`
	MatLabel      string `json:"matlabel"`
	AverageRating string `json:"avgrating"`
	Rating        string `json:"rating"`
	Type          string `json:"type"`
	Updated       string `json:"updated"`
	UNOGSDate     string `json:"unogsdate"`
	Released      string `json:"released"`
	NetflixID     string `json:"netflixid"`
	Runtime       string `json:"runtime"`
	Download      string `json:"download"`
}

/*
ImdbInfo imdb info of program
*/
type imdbInfo struct {
	Runtime   string `json:"runtime"`
	Plot      string `json:"plot"`
	Language  string `json:"language"`
	IMDBID    string `json:"imdbid"`
	Country   string `json:"country"`
	Awards    string `json:"awards"`
	Genre     string `json:"genre"`
	Metascore string `json:"metascore"`
	Votes     string `json:"votes"`
	Rating    string `json:"rating"`
}

/*
Country details of program
*/
type country struct {
	Name          string   `json:"country"`
	CCode         string   `json:"ccode"`
	Seasons       string   `json:"seasons"`
	Expires       string   `json:"expires"`
	New           string   `json:"new"`
	CID           string   `json:"cid"`
	IsLive        string   `json:"islive"`
	SeasonDetails []string `json:"seasondet"`
	Audio         []string `json:"audio"`
	Subtiles      dynamic  `json:"subs"`
}
type dynamic interface{}

/*
People involved in the program
*/
type people struct {
	Actors    []string `json:"actor"`
	Creators  []string `json:"creator"`
	Directors []string `json:"director"`
}

// UNOGSResponse desribes the response body from UNOGS "Load Title Details" api
type UNOGSResponse struct {
	RESULT struct {
		GenreID []string `json:"Genreid"`
		MgName  []string `json:"mgname"`
		//ImdbInfo  imdbInfo  `json:"imdbinfo"`
		NfInfo    nfInfo    `json:"nfinfo"`
		Countries []country `json:"country"`
		Peoples   []people  `json:"people"`
	}
}

// UNOGSAdvanceSearchResponse describes the response body for UNOGS "Advance Search" api
type UNOGSAdvanceSearchResponse struct {
	COUNT string   `json:"COUNT"`
	ITEMS []nfInfo `json:"ITEMS"`
}

/*
GetNetflixDetails gets netflix details of a program from uNoGS api
*/
func GetNetflixDetails(netflixID string) (UNOGSResponse, error) {
	ugnosRes := UNOGSResponse{}
	netflixID = url.PathEscape(netflixID)

	unogsURL := fmt.Sprintf("https://unogs-unogs-v1.p.rapidapi.com/aaapi.cgi?t=loadvideo&q=%s", netflixID)

	body, err := sendUNOGSAPIRequest(unogsURL)
	if err != nil {
		return UNOGSResponse{}, err
	}

	err = json.Unmarshal(body, &ugnosRes)
	if err != nil {
		return UNOGSResponse{}, err
	}
	return ugnosRes, nil

}

/*
UNOGSAdvanceSearchRes contains the list of tv shows and movies searched
*/
var UNOGSAdvanceSearchRes UNOGSAdvanceSearchResponse

/*
UNOGSAdvanceSearch gets netflix details of a program from uNoGS api using the title of a movie/show
*/
func UNOGSAdvanceSearch(title string) error {

	//UNOGSAdvanceSearchRes := UNOGSAdvanceSearchResponse{}
	title = url.PathEscape(title)
	endYear := strconv.Itoa(time.Now().Year())

	unogsURL := fmt.Sprintf("https://unogs-unogs-v1.p.rapidapi.com/aaapi.cgi?q=%s-!1900,%s-!0,5-!0,10-!0-!Any-!Any-!Any-!gt0-!{downloadable}&t=ns&cl=all&st=adv&ob=Relevance&p=1&sa=or", title, endYear)

	body, err := sendUNOGSAPIRequest(unogsURL)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &UNOGSAdvanceSearchRes)
	if err != nil {
		return err
	}

	return nil

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

func loadTitleDetails(id string) (UNOGSResponse, error) {
	unogsResponse := UNOGSResponse{}
	id = url.PathEscape(id)

	titleDetailsURL := fmt.Sprintf("https://unogs-unogs-v1.p.rapidapi.com/aaapi.cgi?t=loadvideo&q=%s", id)

	body, err := sendUNOGSAPIRequest(titleDetailsURL)
	if err != nil {
		return UNOGSResponse{}, err
	}

	err = json.Unmarshal(body, &unogsResponse)
	if err != nil {
		return UNOGSResponse{}, err
	}
	return unogsResponse, nil

}

/*
GetNetflixDetailsForAdvanceSearchResults qwerty keyboard
*/
func GetNetflixDetailsForAdvanceSearchResults(skip int, limit int) ([]UNOGSResponse, error) {
	var netflixInfoList []UNOGSResponse
	var err error
	var result UNOGSResponse
	if len(UNOGSAdvanceSearchRes.ITEMS) <= 0 {
		return netflixInfoList, err
	}
	resultsRemainder := len(UNOGSAdvanceSearchRes.ITEMS) - skip
	if resultsRemainder < limit {
		limit = resultsRemainder
	}
	for index := 0; index < limit; index++ {
		item := UNOGSAdvanceSearchRes.ITEMS[skip+index]
		result, err = loadTitleDetails(item.NetflixID)
		if err != nil {
			break
		}
		netflixInfoList = append(netflixInfoList, result)
	}

	return netflixInfoList, err

}
