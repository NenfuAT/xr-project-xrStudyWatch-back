package model

type User struct {
	ID         string `gorm:"primarykey;type:varchar(26)" json:"id"`
	UserName   string `json:"userName"`
	Mail       string `json:"mail"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Password   string `json:"password"`
}

func InsertUser(u User) error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}
