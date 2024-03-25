package model

type User struct {
	ID         string  `gorm:"primarykey;type:varchar(26)" json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Gender     string  `json:"gender"`
	Age        int     `json:"age"`
	Height     int     `json:"height"`
	Weight     float32 `json:"weight"`
	Occupation string  `json:"occupation"`
	Address    string  `json:"address"`
	Password   string  `json:"password"`
}

func InsertUser(u User) error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}
