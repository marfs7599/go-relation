package entity

type Bpjs struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	NoCard    string `gorm:"unique, not null" json:"no_card"`
	Class     int    `gorm:"not null" json:"class"`
	PatientId int    `gorm:"not null" json:"patient_id"`
}
