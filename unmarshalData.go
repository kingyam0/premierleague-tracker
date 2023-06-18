package main

import (
	"encoding/json"
	"log"
)

func gatherDataUp() []artistInfo {
	data1 := myArtist() // data1 has a value of getData function
	// slices the artistData from any value, assign that to artist
	er := json.Unmarshal(data1, &Artists) // convert the data1 and pointed artists through json into a variable
	if er != nil {
		log.Fatal(er)
		return nil
	}
	for i := 0; i < len(Artists); i++ { // iterates through the length of artist
		r := artistInfo{}         // assigns any typed relation to r
		json.Unmarshal(data1, &r) // converts sliced artist and relation? and pointed r variable thr json to
		// links artists to concerts that's pointed from r
	}
	for i := 0; i < 52; i++ {
		// fmt.Println(Artists[i].Name)
	}
	return Artists // return artist result
}

func gatherDataUpRelation() []relation {
	data1 := myDatesLocations()

	er := json.Unmarshal(data1, &dataLocation)
	if er != nil {
		log.Fatal(er)
		return nil
	}
	for i := 0; i < len(dataLocation.Index); i++ {
		r := relation{}
		json.Unmarshal(data1, &r)
	}
	for i := 0; i < 52; i++ {
		// fmt.Println(dataLocation.Index[i])
	}
	return dataLocation.Index
}

func gatherDataUpDates() []date {
	data1 := myDates()

	er := json.Unmarshal(data1, &eventDates)
	if er != nil {
		log.Fatal(er)
		return nil
	}
	for i := 0; i < len(eventDates.Index); i++ {
		// r := .dateindex
		json.Unmarshal(data1, &eventDates)
	}

	// fmt.Println(eventDates.Index)

	return eventDates.Index
}

func gatherDataUpLocations() []location {
	data1 := myLocations()

	er := json.Unmarshal(data1, &eventLocations)
	if er != nil {
		log.Fatal(er)
		return nil
	}
	for i := 0; i < len(eventLocations.Index); i++ {
		// r := .dateindex
		json.Unmarshal(data1, &eventLocations)
	}

	// fmt.Println(eventLocations.Index)

	return eventLocations.Index
}
