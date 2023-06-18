package tools

import (
	"./tools"
)

// appends all info in our 2nd struct into one variable
func AppendAllData() []artistInfo2 {
	for i := range Articles {
		var appendingdata artistInfo2
		appendingdata.ID = i + 1
		appendingdata.Name = Artists[i].Name
		appendingdata.Members = Artists[i].Members
		appendingdata.Image = Artists[i].Image
		appendingdata.CreationDate = Artists[i].CreationDate
		appendingdata.FirstAlbum = Artists[i].FirstAlbum
		appendingdata.DatesLocations = dataLocation.Index[i].DatesLocations
		appendingdata.Locations = eventLocations.Index[i].Locations
		appendingdata.Concerts = eventDates.Index[i].Dates
		tools.FullArtistInfo = append(FullArtistInfo, appendingdata)

	}
	// fmt.Println("----------------", FullArtistInfo[0])
	return FullArtistInfo
}
