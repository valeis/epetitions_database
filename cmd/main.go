package main

import (
	"E-PETIONS_DATABASE/internal/models"
	"E-PETIONS_DATABASE/pkg/database/postgres"
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func main(){

}

func init(){
	db=postgres.ConnectDatabase()

	err := db.AutoMigrate(models.UserModel{}) 
	if err != nil {
		log.Fatal("failed to migrate user model:", err)
	}
	
	err = db.AutoMigrate(models.PetitionModel{})
	if err != nil {
		log.Fatal("failed to migrate petition model:", err)
	}

	err = db.AutoMigrate(models.VoteModel{})
	if err != nil {
		log.Fatal("failed to migrate vote model:", err)
	}
}