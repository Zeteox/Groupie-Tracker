package DataStruct

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	AllLocations Locations
	AllDates     Dates
	AllRelations Relations
}

type Locations struct {
	Id             int      `json:"id"`
	GroupLocations []string `json:"locations"`
}

type Relations struct {
	Id    int                 `json:"id"`
	Dates map[string][]string `json:"datesLocations"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Form struct {
	SearchbarContent string
	CareerYears      []string
	AlbumYears       []string
	MemberNumber     string
}

type ArtistsAndForm struct {
	Artists  []Artist
	FormData Form
}
