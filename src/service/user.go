package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/common"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/conf"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/model"
)

func CreateUser(req model.User) (model.User, error) {
	c := conf.GetProxyConfig()
	body, _ := json.Marshal(req)
	send, err := http.NewRequest("POST", c.GetString("proxy.objectUpload")+"api/user/create", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("SendError:", err)
	}
	// ベーシック認証の文字列を作成
	authString := c.GetString("proxy.ACCESS_KEY") + ":" + c.GetString("proxy.SECRET_KEY")

	// Base64エンコード
	authEncoded := base64.StdEncoding.EncodeToString([]byte(authString))
	send.Header.Set("Authorization", "Basic "+authEncoded)

	fmt.Println("Request Line:", send.Method, send.URL, send.Header, send.Body)

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
	var response common.UserUploadResponse
	fmt.Println("Response Body:", string(res))
	err = json.Unmarshal(res, &response)
	if err != nil {
		fmt.Println("JsonError:", err)
	}
	var user model.User
	user.ID = response.ID
	user.Name = response.Name
	user.Email = response.Email
	user.Gender = response.Gender
	user.Height = response.Height
	user.Weight = response.Weight
	user.Occupation = response.Occupation
	user.Address = response.Address
	user.Password = req.Password
	model.InsertUser(user)
	return user, nil
}
