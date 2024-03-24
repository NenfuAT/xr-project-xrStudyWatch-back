package common

// クライアント->アプリサーバ

type UserPost struct {
	UserName   string `json:"userName"`
	Mail       string `json:"mail"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Password   string `json:"password"`
}

type ObjectPost struct {
	University    string  `form:"university"`
	Undergraduate string  `form:"undergraduate"`
	Department    string  `form:"department"`
	Major         string  `form:"major"`
	Laboratory    string  `form:"laboratory"`
	Location      string  `form:"location"`
	RoomNum       string  `form:"roomNum"`
	Latitude      float64 `form:"latitude"`
	Longitude     float64 `form:"longitude"`
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

type ObjectUploadResponse struct {
	ObjectID  string `json:"objectId"`
	PosterID  string `json:"posterId"`
	Extension string `json:"extension"`
	Spot      struct {
		ID           string  `json:"id"`
		Name         string  `json:"name"`
		LocationType string  `json:"locationType"`
		Floor        int     `json:"floor"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
	} `json:"spot"`
	UploadURL string `json:"uploadUrl"`
}
