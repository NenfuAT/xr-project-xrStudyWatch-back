package model

type Object struct {
	ID     string `gorm:"primarykey;type:varchar(26)" json:"id"`
	LabID  string `json:"labId"`
	Aspect string `json:"aspect"`
	Height int    `json:"height"`
	Size   string `json:"size"`
}

func InsertObject(o Object) error {
	if err := db.Create(&o).Error; err != nil {
		return err
	}
	return nil
}
