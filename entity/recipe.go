package entity

type Recipe struct {
	Id        int             `gorm:"primaryKey" json:"id"`
	Name      string          `gorm:"not null" json:"name"`
	Dose      string          `gorm:"not null" json:"dose"`
	Type      string          `gorm:"not null" json:"type"`
	PatientId int             `json:"patient_id"`
	Patient   PatientResponse `json:"patient"`
}

type RecipeResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	PatientId int    `json:"-"`
}

func (RecipeResponse) TableName() string {
	return "recipes"
}
