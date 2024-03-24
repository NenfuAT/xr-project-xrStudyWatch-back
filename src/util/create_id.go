package util

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// CreateUserID は新しいULIDを生成して文字列として返します
func CreateUserID() string {
	// 現在の時間を取得
	t := time.Now()

	// 新しいULIDを生成
	entropy := ulid.Monotonic(rand.Reader, 0)
	ulidObj := ulid.MustNew(ulid.Timestamp(t), entropy)

	// ULIDを文字列に変換して返す
	return ulidObj.String()
}
