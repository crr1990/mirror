package models

import (
	"mirror/Databases"
)

type Register struct {
	Id    int    `gorm:"column:id"`
	Code  string `gorm:"column:code"`
	Email string `gorm:"column:email"`
}

func (this *Register) Insert() {
	Databases.DB.Create(&this)
}

func (this *Register) FindByEmail(email string) (Register, error) {
	var r Register
	err := Databases.DB.Where("email = ?", email).Find(&r).Error

	if err != nil {
		return r, err
	}

	return r, err
}
