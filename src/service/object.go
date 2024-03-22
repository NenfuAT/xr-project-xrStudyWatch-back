package service

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/common"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/controller"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/model"
)

func CreateObject(uid string, req controller.ObjectPost, fileheaders []*multipart.FileHeader) error {
	var object model.Object
	var undergraduate model.Undergraduate
	var university model.University
	var laboratory model.Laboratory
	var location model.Location
	var object_post_proxy controller.ObjectPostProxy

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

	body, contentType, err := common.CreatePostObjectBody(object, fileheaders[1])
	if err != nil {
		panic(err)
	}

	send, err := http.NewRequest("POST", "https://example.com/upload", body)
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

	var response controller.ObjectPostProxyResponse
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

	//laboratory
	location.Building = l.Location
	location.Room = l.RoomNum

	return nil
}
