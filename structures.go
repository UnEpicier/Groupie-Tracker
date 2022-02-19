package groupie

type ArtistsStruct struct {
	Tab       []Artists
	CreaDates []int
	AlbDates  []int
	Locations []string
}

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Final struct {
	A         Artist
	Dates     []string
	Locations []string
	Coords    map[string][]string
}

/******************/

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Map struct {
	Place_id     int      `json:"place_id"`
	Licence      string   `json:"licence"`
	Osm_type     string   `json:"osm_type"`
	Osm_id       int      `json:"osm_id"`
	Boundingbox  []string `json:"boundingbox"`
	Lat          string   `json:"lat"`
	Lon          string   `json:"lon"`
	Display_name string   `json:"display_name"`
	Class        string   `json:"class"`
	Type         string   `json:"type"`
	Importance   float64  `json:"importance"`
}
