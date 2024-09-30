package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	dsn := "host=localhost user=example password=example dbname=example port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return nil
	}
	fmt.Println("Connection to database established")
	return db
}

func PopulateDatabase() {
	fmt.Println("Populating database")
	db := CreateConnection()
	err := db.AutoMigrate(&UserStruct{})
	if err != nil {
		panic(err)
		return
	}
	db.Create(&UserStruct{
		Name:       "John",
		FamilyName: "Doe",
	})
	db.Create(&UserStruct{
		Name:       "Doe",
		FamilyName: "John",
	})

}
