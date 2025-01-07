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

	dData := database.DiscordHandle{
		UserName: "username",
		DHCode:   "dh=code",
	}

	if createErr := database.Db().Model(&database.DiscordHandle{}).Create(&dData).Error; createErr != nil {
		panic(createErr)
	}
}
