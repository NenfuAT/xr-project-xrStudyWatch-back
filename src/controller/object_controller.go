package controller

import (
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/common"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/model"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/service"
	"github.com/gin-gonic/gin"
)

func PostObject(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	// "Basic " 接頭辞を削除して、Base64でエンコードされた文字列を取得
	authValue := strings.TrimPrefix(authHeader, "Basic ")

	// Base64デコード
	decoded, err := base64.StdEncoding.DecodeString(authValue)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	// デコードされた文字列を取得
	credentials := string(decoded)

	// ユーザー名とパスワードの分割
	split := strings.SplitN(credentials, ":", 2)
	uid := split[0]
	result := model.GetUserByID(uid)
	if result.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		return
	}

	var req common.ObjectPost

	if err := c.Bind(&req); err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fileHeaders []*multipart.FileHeader

	// "objectFile" フィールドから画像ファイルを取得
	imageFile, imageHeader, err := c.Request.FormFile("objectFile")
	if err != nil {
		fmt.Println("GetImgError:", err)
	}
	defer imageFile.Close()

	// "rawDataFile" フィールドから CSV ファイルを取得
	csvFile, csvHeader, err := c.Request.FormFile("rawDataFile")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer csvFile.Close()

	// ファイルのヘッダーを作成してスライスに追加
	fileHeaders = append(fileHeaders, imageHeader, csvHeader)

	if err := service.CreateObject(uid, req, fileHeaders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func SearchObject(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	// "Basic " 接頭辞を削除して、Base64でエンコードされた文字列を取得
	authValue := strings.TrimPrefix(authHeader, "Basic ")

	// Base64デコード
	decoded, err := base64.StdEncoding.DecodeString(authValue)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	// デコードされた文字列を取得
	credentials := string(decoded)

	// ユーザー名とパスワードの分割
	split := strings.SplitN(credentials, ":", 2)
	uid := split[0]
	result := model.GetUserByID(uid)
	if result.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		return
	}
	var req common.SearchPost

	if err := c.Bind(&req); err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Latitude", req.Latitude)
	fmt.Println("Longitude", req.Longitude)

	// "rawDataFile" フィールドから CSV ファイルを取得
	csvFile, csvHeader, err := c.Request.FormFile("rawDataFile")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer csvFile.Close()

	response, err := service.SearchObject(uid, req, csvHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
