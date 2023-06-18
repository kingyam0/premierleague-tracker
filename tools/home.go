package tools

import (
	"./tools"
	"html/template"
	"log"
	"net/http"
)

// handles all request to root URL
func HomePage(w http.ResponseWriter, r *http.Request) {
	// parses the HTML Files, and then Execute it
	temp, err := template.ParseFiles("docs/templates/index.html")
	if err != nil {
		errorPage(w, r, http.StatusInternalServerError)
		return
	}
	// if not homepage then return error status
	if r.URL.Path != "/" {
		errorPage(w, r, http.StatusNotFound)
		return
	}
	// ranging through struct artistInfo to gather data withins
	var result []artistInfo2
	result = tools.FullArtistInfo
	if err := temp.Execute(w, result); err != nil {
		log.Printf("Execute Error: %v", err)
		http.Error(w, "Error when Executing", http.StatusInternalServerError)
		return
	}
	// fmt.Println(result)
}
