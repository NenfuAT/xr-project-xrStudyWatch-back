package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math/big"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

	image_extension := filepath.Ext(fileheaders[0].Filename)
	image_extension = strings.TrimPrefix(image_extension, ".")
	//プロキシサーバー用
	object_post_proxy.UserID = uid
	object_post_proxy.Extension = image_extension
	object_post_proxy.SpotName = req.Laboratory
	object_post_proxy.Floor = 3               //Todo入力あったら返すようにするとりあえず0
	object_post_proxy.LocationType = "indoor" //Todoとりあえず(ry

	object_post_proxy.Latitude = req.Latitude
	object_post_proxy.Longitude = req.Longitude

	body, contentType, err := common.CreatePostObjectBody(object_post_proxy, fileheaders[1])
	if err != nil {
		fmt.Println("CreateBodyError:", err)
	}

	send, err := http.NewRequest("POST", c.GetString("proxy.objectUpload")+"api/objects/upload", body)
	if err != nil {
		fmt.Println("SendError:", err)
	}
	// ベーシック認証の文字列を作成
	authString := c.GetString("proxy.ACCESS_KEY") + ":" + c.GetString("proxy.SECRET_KEY")

	// Base64エンコード
	authEncoded := base64.StdEncoding.EncodeToString([]byte(authString))
	send.Header.Set("Content-Type", contentType)
	send.Header.Set("Authorization", "Basic "+authEncoded)

	fmt.Println("Request Line:", send.Method, send.URL)

	// HTTPリクエストを実行します
	client := http.Client{}
	resp, err := client.Do(send)
	if err != nil {
		fmt.Println("RequestError:", err)
	}
	defer resp.Body.Close()

	// レスポンスのボディを読み取ります
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadError:", err)
	}
	var response common.ObjectUploadResponse
	fmt.Println("Response Body:", string(res))
	err = json.Unmarshal(res, &response)
	if err != nil {
		fmt.Println("JsonError:", err)
	}
	//university
	university.Name = req.University
	university.UniversityID = ""

	model.InsertUniversity(university)

	//undergraduate
	undergraduate.Name = req.Undergraduate
	universityId, err := model.GetUniversityIdByName(req.University)
	if err != nil {
		fmt.Println("Error:", err)
	}
	undergraduate.UniversityID = universityId
	undergraduate.Department = req.Department
	undergraduate.Major = req.Major

	model.InsertUndergraduate(undergraduate)

	//location
	location.Building = req.Location
	location.Room = req.RoomNum

	model.InsertLocation(location)

	//laboratory
	laboratory.ID = response.Spot.ID
	laboratory.UserID = uid
	undergraduateId, err := model.GetUndergraduateIdByName(req.Undergraduate, universityId)
	if err != nil {
		fmt.Println("Error:", err)
	}
	laboratory.UndergraduateID = undergraduateId
	locationId, err := model.GetLocationIdByName(req.Location)
	if err != nil {
		fmt.Println("Error:", err)
	}
	laboratory.LocationID = locationId
	laboratory.Name = req.Laboratory
	model.InsertLaboratory(laboratory)

	//object
	object.ID = response.ObjectID
	object.LabID = response.Spot.ID
	object.Height = 0

	file, err := fileheaders[0].Open()
	if err != nil {
		fmt.Println("FileOpenError:", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("ImageError:", err)
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
	// 大きい方を1として扱う
	if aspectWidh > aspectHeight {
		aspectHeight /= aspectWidh
		aspectWidh /= aspectWidh
	} else {
		aspectHeight /= aspectHeight
		aspectWidh /= aspectHeight
	}
	object.Height = float32(aspectHeight)
	object.Width = float32(aspectWidh)
	model.InsertObject(object)

	file, err = fileheaders[0].Open()
	if err != nil {
		fmt.Println("FileOpenError:", err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("ファイルの読み込みに失敗しました:", err)
		return err
	}

	fileReader := bytes.NewReader(fileBytes)

	imageupload, err := http.NewRequest("PUT", response.UploadURL, fileReader)
	if err != nil {
		fmt.Println("リクエストを作成できません:", err)
		return err
	}

	// ヘッダーの設定
	imageupload.Header.Set("Content-Type", "application/octet-stream")

	// リクエストの送信
	resp, err = client.Do(imageupload)
	if err != nil {
		fmt.Println("リクエストを送信できません:", err)
		return err
	}
	defer resp.Body.Close()

	// レスポンスの表示
	fmt.Println("imageupload:", resp.Status)
	return nil
}

func SearchObject(uid string, req common.SearchPost, fileheader *multipart.FileHeader) (common.SpotResult, error) {
	c := conf.GetProxyConfig()
	var search_object_proxy common.SearchPostProxy
	search_object_proxy.UserID = uid
	search_object_proxy.Latitude = req.Latitude
	search_object_proxy.Longitude = req.Longitude

	body, contentType, err := common.CreateSearchObjectBody(search_object_proxy, *fileheader)
	if err != nil {
		fmt.Println("CreateBodyError:", err)
		return common.SpotResult{}, err // エラーを返す
	}

	send, err := http.NewRequest("POST", c.GetString("proxy.objectUpload")+"api/objects/search/spot", body)
	if err != nil {
		fmt.Println("SendError:", err)
		return common.SpotResult{}, err // エラーを返す
	}

	// ベーシック認証の文字列を作成
	authString := c.GetString("proxy.ACCESS_KEY") + ":" + c.GetString("proxy.SECRET_KEY")

	// Base64エンコード
	authEncoded := base64.StdEncoding.EncodeToString([]byte(authString))
	send.Header.Set("Content-Type", contentType)
	send.Header.Set("Authorization", "Basic "+authEncoded)

	fmt.Println("Request Line:", send.Method, send.URL)

	// HTTPリクエストを実行します
	client := http.Client{}
	resp, err := client.Do(send)
	if err != nil {
		fmt.Println("RequestError:", err)
		return common.SpotResult{}, err // エラーを返す
	}
	defer resp.Body.Close()

	// レスポンスのボディを読み取ります
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadError:", err)
		return common.SpotResult{}, err // エラーを返す
	}
	var response common.SearchObjectResponse

	var spotresult common.SpotResult
	var arounds []common.AroundObject
	var arrivings []common.ArrivingObject
	for _, spotObj := range response.SpotObjects {
		var arriving common.ArrivingObject
		arriving.ID = spotObj.ID
		object, err := model.GetObjectByID(arriving.ID)
		if err != nil {
			fmt.Println("ReadError:", err)
			return common.SpotResult{}, err
		}
		arriving.Height = object.Height
		arriving.Width = object.Width
		arriving.Size = object.Size
		arriving.ViewURL = spotObj.ViewURL

		arrivings = append(arrivings, arriving)
	}

	for _, areaObj := range response.AreaObjects {
		var around common.AroundObject
		around.ID = areaObj.ID
		object, err := model.GetObjectByID(around.ID)
		if err != nil {
			fmt.Println("ReadError:", err)
			return common.SpotResult{}, err
		}

		laboratory, err := model.GetLaboratoryByID(object.LabID)
		if err != nil {
			fmt.Println("ReadError:", err)
			return common.SpotResult{}, err
		}

		location, err := model.GetLocationByID(laboratory.LocationID)
		if err != nil {
			fmt.Println("ReadError:", err)
			return common.SpotResult{}, err
		}

		undergraduate, err := model.GetUndergraduateByID(laboratory.UndergraduateID)
		if err != nil {
			fmt.Println("ReadError:", err)
			return common.SpotResult{}, err
		}
		university, err := model.GetUniversityByID(undergraduate.UniversityID)
		if err != nil {
			fmt.Println("ReadError:", err)
			return common.SpotResult{}, err
		}

		around.Laboratory.Name = laboratory.Name
		around.Laboratory.Location = location.Building
		around.Laboratory.RoomNum = location.Room

		around.University.Name = university.Name
		around.University.Undergraduate = undergraduate.Name
		around.University.Department = undergraduate.Department
		around.University.Major = undergraduate.Major
		arounds = append(arounds, around)
	}

	if arrivings == nil {
		spotresult.ArrivingObjects = []common.ArrivingObject{}
	} else {
		spotresult.ArrivingObjects = arrivings
	}

	if arounds == nil {
		spotresult.AroundObjects = []common.AroundObject{}
	} else {
		spotresult.AroundObjects = arounds
	}

	fmt.Println("Response Body:", string(res))
	err = json.Unmarshal(res, &response)
	if err != nil {
		fmt.Println("JsonError:", err)
		return common.SpotResult{}, err // エラーを返す
	}

	jsonStr, err := json.Marshal(spotresult)
	if err != nil {
		fmt.Println("JSON変換エラー:", err)
		return common.SpotResult{}, err
	}
	var result model.Result
	result.UserID = uid
	result.Json = string(jsonStr)
	model.InsertOrUpdateResult(result)
	return spotresult, nil
}

func MetaSearch(uid string) (common.SpotResult, error) {
	result, err := model.GetResultByID(uid)
	if err != nil {
		fmt.Println("JsonError:", err)
		return common.SpotResult{}, err // エラーを返す
	}
	var response common.SpotResult
	if err := json.Unmarshal([]byte(result), &response); err != nil {
		fmt.Println("JSONパースエラー:", err)
		return common.SpotResult{}, err // エラーを返す
	}
	return response, nil
}
