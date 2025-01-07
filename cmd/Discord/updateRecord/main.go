package main

import (
	"main/database"

	"github.com/joho/godotenv"
)

func main() {
	if loadErr := godotenv.Load("../../../.env"); loadErr != nil {
		panic(loadErr)
	}

	database.SetupDatabase()

	const usernameToUpdate = "username"
	const newDHCode = "dh=newcode"

	if updateErr := database.Db().Model(&database.DiscordHandle{}).Where("user_name = ?", usernameToUpdate).Update("dh_code", newDHCode).Error; updateErr != nil {
		panic(updateErr)
	}
}
