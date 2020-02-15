package models

import (
	"mirror/Databases"
)

type User struct {
	Id       int    `gorm:"column:id"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	Name     string `gorm:"column:name"`
}

type UserInfo struct {
	Name     string
	Password string
}

func (this *User) Insert() {
	Databases.DB.Create(&this)
}

func (this *User) FindByEmail(email string) (User, error) {
	var r User
	err := Databases.DB.Where("email = ?", email).Find(&r).Error

	return r, err
}

func (this *User) FindById(id int) User {
	var r User
	Databases.DB.Where("id = ?", id).Find(&r)
	return r
}

func (this *User) FindByName(name string) (User, error) {
	var r User
	err := Databases.DB.Where("name = ?", name).Find(&r).Error

	return r,err
}

func (this *User) Update() {
	Databases.DB.Update(&this)

}
