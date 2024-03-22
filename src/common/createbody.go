package common

import (
	"bytes"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/controller"
)

func CreatePostObjectBody(object controller.ObjectPostProxy, fileHeader *multipart.FileHeader) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)

	// 構造体の各フィールドをフォームデータに追加
	fields := map[string]string{
		"userId":       object.UserID,
		"extension":    object.Extension,
		"spotName":     object.SpotName,
		"floor":        strconv.Itoa(object.Floor),
		"locationType": object.LocationType,
		"latitude":     strconv.FormatFloat(object.Latitude, 'f', -1, 64),
		"longitude":    strconv.FormatFloat(object.Longitude, 'f', -1, 64),
	}
	for key, value := range fields {
		if err := mw.WriteField(key, value); err != nil {
			return nil, "", err
		}
	}

	// ファイルを追加
	filePart, err := mw.CreateFormFile("rawDataFile", fileHeader.Filename)
	if err != nil {
		return nil, "", err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	if _, err := io.Copy(filePart, file); err != nil {
		return nil, "", err
	}

	// マルチパートの末尾を追加
	if err := mw.Close(); err != nil {
		return nil, "", err
	}

	// Content-Type ヘッダーを取得
	contentType := mw.FormDataContentType()

	return body, contentType, nil
}
