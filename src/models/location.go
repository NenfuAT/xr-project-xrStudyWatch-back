package model

type Location struct {
	ID       string `gorm:"primarykey;type:varchar(26)" json:"id"`
	Building string `json:"building"`
	Room     string `json:"room"`
}
