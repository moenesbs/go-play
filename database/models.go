package database

import "gorm.io/gorm"

type UserStruct struct {
	ID         uint `json:"id" gorm:"primary_key"`
	Name       string
	FamilyName string
	gorm.Model
}
