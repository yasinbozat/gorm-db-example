package main

import (
	"db_connect_app/models"
	_ "db_connect_app/models"
)

func main() {

	//database.DB.Debug().AutoMigrate(&models.User{}, &models.Key{}, &models.UserKey{}) // Auto Migration User Table
	models.Login("yasinbozatr@gmail.com", "123456789")
	//fmt.Println(utils.GetHWID())
	//models.CreateUser("Yasin", "Bozat", "yasinbozatr@gmail.com", "123456789", "5318332425", "Turkey", "Istanbul")
}
