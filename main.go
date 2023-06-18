package main

import (
	"./tools"
	"log"
	"net/http"
)

// Entry point to our REST server
func main() {
	http.HandleFunc("/", tools.HomePage)
	http.HandleFunc("/response", tools.Response)

	// declaring variables for fucntions to be called
	tools.Articles = tools.GatherDataUp()
	tools.Dates = tools.GatherDataUpDates()
	tools.locations = tools.GatherDataUpLocations()
	tools.DatesLocations = tools.GatherDataUpRelation()
	tools.AppendAllData()

	// Serves and targets files that are within folder
	fs := http.FileServer((http.Dir("docs")))
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	// starting our server
	log.Println("Starting Server :10001")
	err := http.ListenAndServe(":10001", nil)
	log.Fatal(err)
}
