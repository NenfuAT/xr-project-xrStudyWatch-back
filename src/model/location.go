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

func GetLocationIdByName(name string) (uint64, error) {
	var l Location
	if err := db.Where("building = ?", name).First(&l).Error; err != nil {
		return 0, err
	}
	return l.ID, nil
}

func GetLocationByID(id uint64) (Location, error) {
	var l Location
	if err := db.Where("id = ?", id).First(&l).Error; err != nil {
		return Location{}, err
	}
	return l, nil
}
