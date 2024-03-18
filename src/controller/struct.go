package controller

type ObjectPost struct {
	University  UniversityPost `json:"university"`
	Laboratory  LaboratoryPost `json:"laboratory"`
	ObjectFile  string         `json:"objectFile"`
	RawDataFile string         `json:"rawDataFile"`
}

type UniversityPost struct {
	Name          string `json:"name"`
	Undergraduate string `json:"undergraduate"`
	Department    string `json:"department"`
	Major         string `json:"major"`
}

type LaboratoryPost struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	RoomNum  string `json:"roomNum"`
}
