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
	object.Aspect = strconv.Itoa(aspectWidh) + ":" + strconv.Itoa(aspectHeight)
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

func SearchObject(uid string, req common.SearchPost, fileheader *multipart.FileHeader) (common.SearchObjectResponse, error) {
	c := conf.GetProxyConfig()
	var search_object_proxy common.SearchPostProxy
	search_object_proxy.UserID = uid
	search_object_proxy.Latitude = req.Latitude
	search_object_proxy.Longitude = req.Longitude

	body, contentType, err := common.CreateSearchObjectBody(search_object_proxy, *fileheader)
	if err != nil {
		fmt.Println("CreateBodyError:", err)
		return common.SearchObjectResponse{}, err // エラーを返す
	}

	send, err := http.NewRequest("POST", c.GetString("proxy.objectUpload")+"api/objects/search/spot", body)
	if err != nil {
		fmt.Println("SendError:", err)
		return common.SearchObjectResponse{}, err // エラーを返す
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
		return common.SearchObjectResponse{}, err // エラーを返す
	}
	defer resp.Body.Close()

	// レスポンスのボディを読み取ります
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadError:", err)
		return common.SearchObjectResponse{}, err // エラーを返す
	}
	var response common.SearchObjectResponse
	fmt.Println("Response Body:", string(res))
	err = json.Unmarshal(res, &response)
	if err != nil {
		fmt.Println("JsonError:", err)
		return common.SearchObjectResponse{}, err // エラーを返す
	}
	return response, nil
}
