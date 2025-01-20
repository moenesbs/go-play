package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	dsn := "host=localhost user=admin password=admin dbname=admin port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return nil
	}
	fmt.Println("Connection to database established")
	return db
}

func InsertUser(user UserStruct) int {
	db := CreateConnection()
	id := db.Create(&user)
	if id.Error != nil {
		panic(id.Error)
	}
	fmt.Println("User inserted successfully to DB, id: ", user.ID)
	return int(user.ID)
}

func GetUsers() []UserStruct {
	db := CreateConnection()
	var usersdb []UserStruct
	db.Select("Name").Find(&usersdb)
	return usersdb
}
