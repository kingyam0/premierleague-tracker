package tools

import (
	"io/ioutil"
	"log"
	"net/http"
)

func MyArtist() []byte {
	data1, e1 := http.Get("https://groupietrackers.herokuapp.com/api/artists") // requests data from the link server
	if e1 != nil {
		log.Fatal(e1)
		return nil // error reading
	}
	defer data1.Body.Close()
	data2, e2 := ioutil.ReadAll(data1.Body) // read the data from data1

	if e2 != nil {
		log.Fatal(e2)
		return nil // return the error
	}
	return data2 // else return the read data that had been requested
}

func MyDates() []byte {
	data1, e1 := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if e1 != nil {
		log.Fatal(e1)
		return nil
	}
	defer data1.Body.Close()
	data2, e2 := ioutil.ReadAll(data1.Body)

	if e2 != nil {
		log.Fatal(e2)
		return nil
	}
	return data2
}

func MyLocations() []byte {
	data1, e1 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if e1 != nil {
		log.Fatal(e1)
		return nil
	}
	defer data1.Body.Close()
	data2, e2 := ioutil.ReadAll(data1.Body)

	if e2 != nil {
		log.Fatal(e2)
		return nil
	}
	return data2
}

func MyDatesLocations() []byte {
	data1, e1 := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if e1 != nil {
		log.Fatal(e1)
		return nil
	}
	defer data1.Body.Close()
	data2, e2 := ioutil.ReadAll(data1.Body)

	if e2 != nil {
		log.Fatal(e2)
		return nil
	}
	return data2
}
