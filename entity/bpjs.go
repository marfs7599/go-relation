package entity

type Bpjs struct {
	Id        int             `gorm:"primaryKey" json:"id"`
	NoCard    string          `gorm:"unique, not null" json:"no_card"`
	Class     int             `gorm:"not null" json:"class"`
	PatientId int             `gorm:"not null" json:"patient_id"`
	Patient   PatientResponse `json:"patient"`
}

type BpjsRensponse struct {
	Id        int    `json:"id"`
	NoCard    string `json:"no_card"`
	Class     int    `json:"class"`
	PatientId int    `json:"-"`
}

func (BpjsRensponse) TableName() string {
	return "bpjs"
}
