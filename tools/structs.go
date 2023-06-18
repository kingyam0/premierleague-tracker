package tools

type errorData struct {
	Num  int
	Text string
}

// defining artist information struct for unmarshling
type artistInfo struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Concerts       string              `json:"concertDates"`
	Locations      string              `json:"locations"`
	Relations      string              `json:"Relations"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// struct to deal with info
// cannot be unmarshelled in one step hence why there are two structs
type artistInfo2 struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Concerts       []string            `json:"concertDates"`
	Locations      []string            `json:"locations"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// removes the word index from date list
type dateremoveindex struct {
	Index []date `json:"index"`
}

type locationremoveindex struct {
	Index []location `json:"index"`
}

type relationremoveindex struct {
	Index []relation `json:"index"`
}

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
