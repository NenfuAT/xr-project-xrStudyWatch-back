package service

import (
	"encoding/json"
	"image"
	"io"
	"math/big"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/common"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/conf"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/model"
)

func CreateObject(uid string, req common.ObjectPost, fileheaders []*multipart.FileHeader) error {
	c := conf.GetProxyConfig()
	var object model.Object
	var undergraduate model.Undergraduate
	var university model.University
	var laboratory model.Laboratory
	var location model.Location
	var object_post_proxy common.ObjectPostProxy

	u := req.University
	l := req.Laboratory

	image_extension := filepath.Ext(fileheaders[0].Filename)

	//プロキシサーバー用
	object_post_proxy.UserID = uid
	object_post_proxy.Extension = image_extension
	object_post_proxy.SpotName = l.Name
	object_post_proxy.Floor = 0              //Todo入力あったら返すようにするとりあえず0
	object_post_proxy.LocationType = "indor" //Todoとりあえず(ry

	object_post_proxy.Latitude = l.Latitude
	object_post_proxy.Longitude = l.Longitude

	body, contentType, err := common.CreatePostObjectBody(object_post_proxy, fileheaders[1])
	if err != nil {
		panic(err)
	}

	send, err := http.NewRequest("POST", c.GetString("proxy.objectUpload"), body)
	if err != nil {
		panic(err)
	}
	send.Header.Set("Content-Type", contentType)

	// HTTPリクエストを実行します
	client := http.Client{}
	resp, err := client.Do(send)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// レスポンスのボディを読み取ります
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response common.ObjectPostProxyResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		panic(err)
	}

	//university
	university.Name = u.Name
	university.UniversityID = ""

	model.InsertUniversity(university)

	//undergraduate
	undergraduate.Name = u.Undergraduate
	universityId, err := model.GetUniversityIdByName(u.Name)
	if err != nil {
		panic(err)
	}
	undergraduate.UniversityID = universityId
	undergraduate.Department = u.Department
	undergraduate.Major = u.Major

	model.InsertUndergraduate(undergraduate)

	//location
	location.Building = l.Location
	location.Room = l.RoomNum

	model.InsertLocation(location)

	//laboratory
	laboratory.ID = response.Spot.ID
	laboratory.UserID = uid
	undergraduateId, err := model.GetUndergraduateIdByName(u.Name, universityId)
	if err != nil {
		panic(err)
	}
	laboratory.UndergraduateID = undergraduateId
	locationId, err := model.GetLocationIdByName(l.Location)
	if err != nil {
		panic(err)
	}
	laboratory.LocationID = locationId
	laboratory.Name = l.Name
	model.InsertLaboratory(laboratory)

	//object
	object.ID = response.ObjectID
	object.LabID = response.Spot.ID
	object.Height = 0

	file, err := fileheaders[1].Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	//画像サイズ
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	object.Size = strconv.Itoa(width) + "x" + strconv.Itoa(height)
	//アスペクト比
	bigA := big.NewInt(int64(width))
	bigB := big.NewInt(int64(height))
	gcd := bigA.GCD(nil, nil, bigA, bigB)
	aspectWidh := width / int(gcd.Int64())
	aspectHeight := height / int(gcd.Int64())
	object.Aspect = strconv.Itoa(aspectWidh) + ":" + strconv.Itoa(aspectHeight)
	model.InsertObject(object)
	return nil
}
