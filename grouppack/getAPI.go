package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Unmarshalling API JSON
//ref: https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get

const api = "https://groupietrackers.herokuapp.com/api"

func ReadURL(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body) // response body is []byte
	return body
}

// Artists collection from API
type ArtistsAPI struct {
	// note even though Id, Logo.. start with upper case to be exportable
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	
	// Locations    string   `json:"locations"`
	// ConcertDates string   `json:"concertDates"`
	// Relations    string   `json:"relations"` 
}

// Embedded JSON
type LocationsAPI struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}
type DatesAPI struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}
type RelationsAPI struct {
	Index []struct {
		ID        int                 `json:"id"`
		Relations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func unmarchAPI(url string) {
	var Artists []ArtistsAPI
	var Locations LocationsAPI
	var Dates DatesAPI
	var Relations RelationsAPI

	a := ReadURL(url + "/artists")
	l := ReadURL(url + "/locations")
	d := ReadURL(url + "/dates")
	r := ReadURL(url + "/relation")

	err1 := json.Unmarshal(a, &Artists)
	err2 := json.Unmarshal(l, &Locations)
	err3 := json.Unmarshal(d, &Dates)
	err4 := json.Unmarshal(r, &Relations)

	FullList := make(map[string]interface{})

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Can not unmarshal JSON")
	 } else {
	 	FullList = map[string]interface{}{
	 		"Artists": Artists, "Locations": Locations, "Dates": Dates, "Relations": Relations}
	 }
	fmt.Printf("%+v\n" , FullList["Artists"] )
}

func main() {
	unmarchAPI(api)
}

// func OrganiseAPI(url string) string {

// 	var orgArt []ArtistsAPI
// 	var orgLoc Locations
// 	var orgDates Dates
// 	var orgRel Relations
// 	//unmarshaling JSON pointing to orgArt-orgRel variable
// 	art := UnmarshAPI(url+"artists", orgArt)
// 	loc := UnmarshAPI(url+"locations", orgArt)
// 	dates := UnmarshAPI(url+"dates", orgArt)
// 	rel := UnmarshAPI(url+"relation", orgRel)

// 	if art == "no problem on unmarshing" && dates == "no problem on unmarshing" && loc == "no problem on unmarshing" && rel == "no problem on unmarshing" {
// 		a := make(FullList, len(orgArt))
// 		for i := range orgArt {
// 			var a ArtistsAPI
// 			a.Id = orgArt[i].Id
// 			a.Logo = orgArt[i].Logo
// 			a.Name = orgArt[i].Name
// 			a.Members = orgArt[i].Members
// 			a.Establish = orgArt[i].Establish
// 			a.FirstAlbum = orgArt[i].FirstAlbum
// 			for j := range orgLoc.Index[i].Locations {
// 				a.Locations = orgLoc.Index[i].Locations[j]
// 			}
// 			for k := range orgDates.Index[i].ConcertDates {
// 				a.ConcertDates = orgDates.Index[i].ConcertDates[k]
// 			}
// 		}
// 		b, _ := json.Marshal(a)
// 		f, _ := os.Create("grouppack/base.json")
// 		ioutil.WriteFile("grouppack/base.json", b, 0644)
// 		defer f.Close()
// 		fmt.Println(art)

// 	} else {
// 		fmt.Println("problem on unmarshing, check JSON")
// 	}
// 	return "API organised"
// }
// func main() {
// 	print := OrganiseAPI("https://groupietrackers.herokuapp.com/api/")
// 	fmt.Println(print)
// }
