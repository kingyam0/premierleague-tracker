package main

func getArtistID(id int) artistInfo2 {
	var artistnotfound artistInfo2

	for _, item := range FullArtistInfo {
		if item.ID == id {
			return item
		}
	}
	return artistnotfound
}