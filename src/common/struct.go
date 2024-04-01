package common

// クライアント->アプリサーバ

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

type SearchPost struct {
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

type UserUploadResponse struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Gender     string  `json:"gender"`
	Age        int     `json:"age"`
	Height     int     `json:"height"`
	Weight     float32 `json:"weight"`
	Occupation string  `json:"occupation"`
	Address    string  `json:"address"`
}

type SearchPostProxy struct {
	UserID    string  `form:"userId"`
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
}

type Spot struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Floor        int     `json:"floor"`
	LocationType string  `json:"locationType"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type ObjectUploadResponse struct {
	ObjectID  string `json:"objectId"`
	PosterID  string `json:"posterId"`
	Extension string `json:"extension"`
	Spot      Spot   `json:"spot"`
	UploadURL string `json:"uploadUrl"`
}

type Object struct {
	ID       string `json:"id"`
	PosterID string `json:"posterId"`
	Spot     Spot   `json:"spot"`
	ViewURL  string `json:"viewUrl"`
}

type SearchObjectResponse struct {
	UserID      string   `json:"userId"`
	SpotObjects []Object `json:"spotObjects"`
	AreaObjects []Object `json:"areaObjects"`
}

// ArrivingObject 構造体は、到着オブジェクトを表します。
type ArrivingObject struct {
	ID      string  `json:"id"`
	Width   float32 `json:"width"`
	Height  float32 `json:"height"`
	Size    string  `json:"size"`
	ViewURL string  `json:"viewUrl"`
}

// AroundObject 構造体は、周囲のオブジェクトを表します。
type AroundObject struct {
	ID         string `json:"id"`
	Laboratory struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		RoomNum  string `json:"roomNum"`
	} `json:"laboratory"`
	University struct {
		Name          string `json:"name"`
		Undergraduate string `json:"undergraduate"`
		Department    string `json:"department"`
		Major         string `json:"major"`
	} `json:"university"`
}

// Data 構造体は、提供されたJSONデータ全体を表します。
type SpotResult struct {
	ArrivingObjects []ArrivingObject `json:"arrivingObjects"`
	AroundObjects   []AroundObject   `json:"aroundObjects"`
}

// Error implements error.
func (s SearchObjectResponse) Error() string {
	panic("unimplemented")
}
