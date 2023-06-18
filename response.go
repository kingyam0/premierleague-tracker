package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func response(w http.ResponseWriter, r *http.Request) {
	temp, er := template.ParseFiles("docs/templates/response.html")
	if er != nil {
		log.Fatal(er)
		errorPage(w, r, http.StatusInternalServerError)
		return
	}
	idvalue := r.FormValue("id")

	idvalueint, err := strconv.Atoi(idvalue)
	if err != nil {
		fmt.Println("error")
	}

	v := getArtistID(idvalueint)

	temp.Execute(w, v)
}
