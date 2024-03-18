package model

type Object struct {
	ID     string `gorm:"primarykey;type:varchar(26)" json:"id"`
	LabID  string `json:"labId"`
	Aspect int    `json:"aspect"`
	Height int    `json:"height"`
	Size   int    `json:"size"`
}

func InsertObject(o Object) error {
	if err := db.Create(&o).Error; err != nil {
		return err
	}
	return nil
}
