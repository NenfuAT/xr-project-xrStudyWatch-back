package model

type Laboratory struct {
	ID              string `gorm:"primarykey;type:varchar(26)" json:"id"`
	UserID          string `json:"userId"`
	UndergraduateID string `json:"undergraduateId"`
	LocationID      string `json:"locationId"`
	Name            string `json:"name"`
	Homepage        string `json:"homepage"`
}

func InsertLaboratory(l Laboratory) error {
	if err := db.Create(&l).Error; err != nil {
		return err
	}
	return nil
}
