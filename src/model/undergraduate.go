package model

type Undergraduate struct {
	ID           uint64 `gorm:"primarykey;autoIncrement" json:"id"`
	UniversityID string `json:"universityId"`
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
