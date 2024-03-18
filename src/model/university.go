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
