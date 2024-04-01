package model

type Laboratory struct {
	ID              string `gorm:"primarykey;type:varchar(26)" json:"id"`
	UserID          string `json:"userId"`
	UndergraduateID uint64 `json:"undergraduateId"`
	LocationID      uint64 `json:"locationId"`
	Name            string `json:"name"`
}

func InsertLaboratory(l Laboratory) error {
	if err := db.Create(&l).Error; err != nil {
		return err
	}
	return nil
}

func GetLaboratoryByID(id string) (Laboratory, error) {
	var l Laboratory
	if err := db.Where("id = ?", id).First(&l).Error; err != nil {
		return Laboratory{}, err
	}
	return l, nil
}
