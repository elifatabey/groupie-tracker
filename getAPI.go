package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Unmarshalling API JSON
//ref: https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
const api = "https://groupietrackers.herokuapp.com/api"

var tpl *template.Template

type Content struct {
	FullList interface{}
}

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

type FullList struct {
	Artists   []ArtistsAPI
	Relations  RelationsAPI
}

var returno FullList

func unmarchAPI(url string) FullList {
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


	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Can not unmarshal JSON")
	} else {
		returno =	FullList { Artists: Artists,  Relations: Relations}
	}

	return returno;
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // handling the CSS
	tpl, _ = template.ParseGlob("static/*.html")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about/", About)

	fmt.Printf("Starting server at port 3000\n")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "GET":
		template, _ := template.ParseFiles("./static/index.html")

		page := Content{FullList: unmarchAPI(api).Artists}
		template.Execute(writer, page)

	default:
		fmt.Fprintf(writer, "Sorry, only GET methods are supported.")
	}
}

func About(writer http.ResponseWriter, request *http.Request) {
	
	id, _ := strconv.Atoi(strings.Split(request.URL.Path, "/")[len(strings.Split(request.URL.Path, "/"))-1])
	//fmt.Print(unmarchAPI(api).Artists[id])
	switch request.Method {
	case "GET":
		template, _ := template.ParseFiles("./static/about.html")

		//page =  {unmarchAPI(api).Artists[id], unmarchAPI(api).Relations.Index[id] }
		page := Content{FullList: unmarchAPI(api).Artists[id]}
		template.Execute(writer, page)

	default:
		fmt.Fprintf(writer, "Sorry, only GET methods are supported.")
	}
}
