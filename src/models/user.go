package model

type User struct {
	ID       string `gorm:"primarykey;type:varchar(28)" json:"id"`
	UserName string `json:"userName"`
}
