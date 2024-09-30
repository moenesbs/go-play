package database

import "gorm.io/gorm"

type UserStruct struct {
	gorm.Model
	Name       string
	FamilyName string
}
