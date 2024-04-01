package model

type Object struct {
	ID     string  `gorm:"primarykey;type:varchar(26)" json:"id"`
	LabID  string  `json:"labId"`
	Height float32 `json:"height"`
	Width  float32 `json:"weight"`
	Size   string  `json:"size"`
}

func InsertObject(o Object) error {
	if err := db.Create(&o).Error; err != nil {
		return err
	}
	return nil
}

func GetObjectByID(id string) (Object, error) {
	var object Object
	if err := db.Where("id = ? ", id).First(&object).Error; err != nil {
		return Object{}, err
	}
	return object, nil
}
