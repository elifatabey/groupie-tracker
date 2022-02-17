package grouppack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Unmarshalling API JSON
// ref: https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get

func UnmarshAPI(url string, listStruct interface{}) string {
	resp, err1 := http.Get(url)
	body, err2 := ioutil.ReadAll(resp.Body)
	err3 := json.Unmarshal(body, &listStruct)
	if err1 != nil || err2 != nil || err3 != nil {
		return "Error with unmarshing of URL: " + url
	}
	return "no problem"
}

// Artists collection from API
type ArtistsAPI struct {
	Id         int      `json:"id"`
	Logo       string   `json:"image"`
	Name       string   `json:"name"`
	Members    []string `json:"members"`
	Establish  int      `json:"creationDate"`
	FirstAlbum string   `json:"firstAlbum"`
}

// Relation and Location collection from API
type RelLocAPI struct {
	Index []struct {
		Rel_LD map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// to reorginised
type FullList []struct {
	Id         int
	Name       string
	Logo       string
	Members    []string
	Establish  int
	FirstAlbum int64
	Concerts   []struct {
		Location string
		// Coords []float64
		Dates []int64
	}
}

var ListAPI FullList

func OrganiseAPI(url string) string {

	var orgArt ArtistsAPI
	var orgRel RelLocAPI

	art := UnmarshAPI(url+"artists", orgArt)
	rel := UnmarshAPI(url+"relation", orgRel)

	

}