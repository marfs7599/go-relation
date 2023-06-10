package config

import (
	"fmt"
	"relation/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	newDB := "root:@tcp(127.0.0.1:3306)/relation_poli"
	DB, err = gorm.Open(mysql.Open(newDB), &gorm.Config{})
	if err != nil {
		panic("Connection failure!")
	}
	fmt.Println("Connection success!")
}

func Migration() {
	err := DB.AutoMigrate(
		&entity.Patient{},
		&entity.Bpjs{},
	)

	if err != nil {
		fmt.Println("Migration failed!")
	}

	fmt.Println("Migration success!")
}
