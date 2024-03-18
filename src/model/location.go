package model

type Location struct {
	ID       uint64 `gorm:"primarykey;autoIncrement" json:"id"`
	Building string `json:"building"`
	Room     string `json:"room"`
}

func InsertLocation(l Location) error {
	if err := db.Create(&l).Error; err != nil {
		return err
	}
	return nil
}
