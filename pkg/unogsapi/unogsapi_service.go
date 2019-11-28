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
	GetAvailability(string) (UNOGSResponse, error)
}

/*
NfInfo is netflix info of program
*/
type nfInfo struct {
	Image1        string `json:"image1"`
	Image2        string `json:"image2"`
	Title         string `json:"title"`
	Synopsis      string `json:"synopsis"`
	MatLevel      string `json:"matlevel"`
	MatLabel      string `json:"matlabel"`
	AverageRating string `json:"avgrating"`
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
	Subtiles      []string `json:"subs"`
}

/*
People involved in the program
*/
type people struct {
	Actors    []string `json:"actor"`
	Creators  []string `json:"creator"`
	Directors []string `json:"director"`
}

// UNOGSResponse desribes the responses body from utelly
type UNOGSResponse struct {
	RESULT struct{
		GenreID   []string  `json:"Genreid"`
		MgName    []string  `json:"mgname"`
		ImdbInfo  imdbInfo  `json:"imdbinfo"`
		NfInfo    nfInfo    `json:"nfinfo"`
		Countries []country `json:"country"`
		Peoples []people  `json:"people"`
	}
}

/*
GetNetflixDetails gets netflix details of a program from uNoGS api
*/
func GetNetflixDetails(programID string) (UNOGSResponse, error) {
	ugnosRes := UNOGSResponse{}
	programID = url.PathEscape(programID)

	unogsURL := fmt.Sprintf("https://unogs-unogs-v1.p.rapidapi.com/aaapi.cgi?t=loadvideo&q=%s", programID)

	req, err := http.NewRequest("GET", unogsURL, nil)
	if err != nil {
		return ugnosRes, err
	}
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPI_API_UNOGS_HOST"))
	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_UNOGS_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ugnosRes, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("e")
		return UNOGSResponse{}, err
	}

	err = json.Unmarshal(body, &ugnosRes)
	if err != nil {
		fmt.Println("r")
		return UNOGSResponse{}, err
	}
	return ugnosRes, nil

}
