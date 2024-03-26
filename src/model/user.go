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

func GetUserByEmail(email string) User {
	var user User
	result := db.Where("email = ? ", email).First(&user)
	if result.Error != nil {
		return User{}
	}
	return user
}
func GetUserByEmailAndPassword(email, password string) (User, error) {
	var user User
	result := db.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
