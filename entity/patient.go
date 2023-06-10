package entity

type Patient struct {
	Id          int              `gorm:"primaryKey" json:"id"`
	Nik         int              `gorm:"unique, not null" json:"nik"`
	Name        string           `gorm:"not null" json:"name"`
	Address     string           `gorm:"not null" json:"address"`
	Phonenumber string           `gorm:"not null" json:"phonenumber"`
	Bpjs        BpjsRensponse    `json:"bpjs"`
	Recipe      []RecipeResponse `json:"recipe"`
}

type PatientResponse struct {
	Id          int    `json:"-"`
	Nik         int    `json:"nik"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phonenumber string `json:"phonenumber"`
}

func (PatientResponse) TableName() string {
	return "patients"
}
