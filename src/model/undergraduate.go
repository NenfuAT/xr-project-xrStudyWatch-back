package model

type Undergraduate struct {
	ID           uint64 `gorm:"primarykey;autoIncrement" json:"id"`
	UniversityID uint64 `json:"universityId"`
	Name         string `json:"name"`
	Department   string `json:"department"`
	Major        string `json:"major"`
}

func InsertUndergraduate(u Undergraduate) error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func GetUndergraduateIdByName(name string, uId uint64) (uint64, error) {
	var u Undergraduate
	if err := db.Where("name = ? AND university_id = ?", name, uId).First(&u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}

func GetUndergraduateByID(id uint64) (Undergraduate, error) {
	var u Undergraduate
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return Undergraduate{}, err
	}
	return u, nil
}
