package controller

// クライアント->アプリサーバ
type ObjectPost struct {
	University  UniversityPost `form:"university"`
	Laboratory  LaboratoryPost `form:"laboratory"`
	ObjectFile  string         `form:"objectFile"`
	RawDataFile string         `form:"rawDataFile"`
}

type UniversityPost struct {
	Name          string `form:"name"`
	Undergraduate string `form:"undergraduate"`
	Department    string `form:"department"`
	Major         string `form:"major"`
}

type LaboratoryPost struct {
	Name      string  `form:"name"`
	Location  string  `form:"location"`
	RoomNum   string  `form:"roomNum"`
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
}

// /アプリサーバー -> プロキシサーバ
type ObjectPostProxy struct {
	UserID       string  `form:"userId"`
	Extension    string  `form:"extension"`
	SpotName     string  `form:"spotName"`
	Floor        int     `form:"floor"`
	LocationType string  `form:"locationType"`
	Latitude     float64 `form:"latitude"`
	Longitude    float64 `form:"longitude"`
}

type Spot struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	LocationType string  `json:"locationType"`
	Floors       int     `json:"floors"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type ObjectPostProxyResponse struct {
	ObjectID  string `json:"objectId"`
	PosterID  string `json:"posterId"`
	Extension string `json:"extension"`
	Spot      Spot   `json:"spot"`
	UploadURL string `json:"uploadUrl"`
}
