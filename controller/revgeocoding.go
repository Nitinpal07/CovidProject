package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apikey = "-0nwQ1Cf1TQmwR9xQw_YzhGutZ499vY4WMLfF_9ejgs"
var address string

// struct type of json coming from gecoding api
type output struct {
	Items []struct {
		Title string `json:"title"`
		Address    struct {
			Label       string `json:"label"`
			CountryCode string `json:"countryCode"`
			CountryName string `json:"countryName"`
			StateCode   string `json:"stateCode"`
			State       string `json:"state"`
			County      string `json:"county"`
			City        string `json:"city"`
			District    string `json:"district"`
			PostalCode  string `json:"postalCode"`
		} `json:"address"`
	} `json:"items"`
}
// getState - method to getState from gps coordinates 
func getState(latitude string,longitude string) string {
	url := "https://revgeocode.search.hereapi.com/v1/revgeocode?apiKey=" + apikey + "&at=" + fmt.Sprint(latitude) + "," + fmt.Sprint(longitude)

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(string(body))
	var data output
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	var state string
	for _, add := range data.Items {
		state = add.Address.State
		//fmt.Printf(state)
	}
	return state
}

