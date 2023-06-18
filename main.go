package main

import (
	"log"
	"net/http"
)

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var (
	FullArtistInfo    []artistInfo2
	Articles          []artistInfo
	locations         []location
	Dates             []date
	DatesLocations    []relation
	DateswithIndex    dateremoveindex
	RelationwithIndex relationremoveindex
	LocationWithIndex locationremoveindex
)

var (
	eventLocations locationremoveindex
	eventDates     dateremoveindex
	Artists        []artistInfo
	dataLocation   relationremoveindex
)

// Entry point to our REST server
func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/response", response)
	// http.HandleFunc("/search", searchBar)

	// declaring variables for fucntions to be called
	Articles = gatherDataUp()
	Dates = gatherDataUpDates()
	locations = gatherDataUpLocations()
	DatesLocations = gatherDataUpRelation()
	appendalldata()

	// Serves and targets files that are within folder
	fs := http.FileServer((http.Dir("docs")))
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	// starting our server
	log.Println("Starting Server :10001")
	err := http.ListenAndServe(":10001", nil)
	log.Fatal(err)
}
