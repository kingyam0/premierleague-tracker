package tools

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// functions to check & deal with all errors
func errorPage(w http.ResponseWriter, r *http.Request, errorPage int) {
	temp, er := template.ParseFiles("docs/templates/error.html")
	if er != nil {
		log.Fatal(er)
		return
	}
	w.WriteHeader(errorPage)
	errData := errorData{Num: errorPage}
	if errorPage == 404 {
		errData.Text = "Page Not Found"
	} else if errorPage == 400 {
		errData.Text = "Bad Request"
	} else if errorPage == 500 {
		errData.Text = "Internal Server Error"
	}
	fmt.Println(errData)
	temp.Execute(w, errData)
}
