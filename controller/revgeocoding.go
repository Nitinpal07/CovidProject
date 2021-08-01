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

func getstate(latitude string,longitude string) string {
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
		state = add.Address.City
		//fmt.Printf(state)
	}
	return state
}

//json format of reverse geocoding api result
// {
//   "items": [
//     {
//       "title": "Creativity",
//       "id": "here:pds:place:356jx7ps-6b18784695c70f5001df83da965f2570",
//       "resultType": "place",
//       "address": {
//         "label": "Creativity, Shivajinagar, Bengaluru 560001, India",
//         "countryCode": "IND",
//         "countryName": "India",
//         "stateCode": "KA",
//         "state": "Karnataka",
//         "county": "Bengaluru",
//         "city": "Bengaluru",
//         "district": "Shivajinagar",
//         "postalCode": "560001"
//       },
//       "position": {
//         "lat": 12.98023,
//         "lng": 77.60094
//       },
//       "access": [
//         {
//           "lat": 12.98022,
//           "lng": 77.60093
//         }
//       ],
//       "distance": 11,
//       "categories": [
//         {
//           "id": "800-8200-0174",
//           "name": "Школа",
//           "primary": true
//         }
//       ]
//     }
//   ]
// }
