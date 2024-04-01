package model

import (
	"errors"

	"gorm.io/gorm"
)

type Result struct {
	UserID string `json:"userId"`
	Json   string `json:"json"`
}

func InsertOrUpdateResult(r Result) error {
	// トランザクションを開始
	tx := db.Begin()

	// ユーザーIDに対する結果を検索
	var existingResult Result
	if err := tx.Where("user_id = ?", r.UserID).First(&existingResult).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// エラーが発生した場合はロールバックしてエラーを返す
			tx.Rollback()
			return err
		}
	}

	// レコードが存在する場合は結果を更新
	if existingResult.UserID != "" {
		if err := tx.Model(&existingResult).Where("user_id = ?", r.UserID).Update("json", r.Json).Error; err != nil {
			// エラーが発生した場合はロールバックしてエラーを返す
			tx.Rollback()
			return err
		}
	} else {
		// レコードが存在しない場合は新しいレコードを挿入
		if err := tx.Create(&r).Error; err != nil {
			// エラーが発生した場合はロールバックしてエラーを返す
			tx.Rollback()
			return err
		}
	}

	// トランザクションをコミット
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func GetResultByID(id string) (string, error) {
	var resul Result
	if err := db.Where("user_id = ? ", id).First(&resul).Error; err != nil {
		return "", err
	}
	return resul.Json, nil
}
