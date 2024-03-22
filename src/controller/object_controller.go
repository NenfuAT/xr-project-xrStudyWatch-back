package controller

import (
	"mime/multipart"
	"net/http"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/common"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/service"
	"github.com/gin-gonic/gin"
)

func PostObject(c *gin.Context) {

	uid := "hoge"

	var req common.ObjectPost

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var fileHeaders []*multipart.FileHeader

	// "objectFile" フィールドから画像ファイルを取得
	imageFile, imageHeader, err := c.Request.FormFile("objectFile")
	if err != nil {
		// エラーハンドリング
	}
	defer imageFile.Close()

	// "rawDataFile" フィールドから CSV ファイルを取得
	csvFile, csvHeader, err := c.Request.FormFile("rawDataFile")
	if err != nil {
		// エラーハンドリング
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
