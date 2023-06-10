package entity

type Patient struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Nik         int    `gorm:"unique, not null" json:"nik"`
	Name        string `gorm:"not null" json:"name"`
	Address     string `gorm:"not null" json:"address"`
	Phonenumber string `gorm:"not null" json:"phonenumber"`
}
