package model

type Undergraduate struct {
	ID           string `json:"id"`
	UniversityID string `json:"universityId"`
	Name         string `json:"name"`
	Department   string `json:"department"`
	Major        string `json:"major"`
}
