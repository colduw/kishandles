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

	const usernameToDelete = "username"

	if deleteErr := database.Db().Unscoped().Delete(&database.DiscordHandle{}, "user_name = ?", usernameToDelete).Error; deleteErr != nil {
		panic(deleteErr)
	}
}
