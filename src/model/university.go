package model

type University struct {
	ID           uint64 `gorm:"primarykey;autoIncrement" json:"id"`
	UniversityID string `json:"universityId"`
	Name         string `json:"name"`
}

func InsertUniversity(u University) error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func GetUniversityIdByName(name string) (uint64, error) {
	var u University
	if err := db.Where("name = ?", name).First(&u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}
