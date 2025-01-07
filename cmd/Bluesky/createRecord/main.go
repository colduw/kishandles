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

	bData := database.BskyHandle{
		Handle: "handle-goes-here",
		DID:    "did:plc:stuff",
	}

	if createErr := database.Db().Model(&database.BskyHandle{}).Create(&bData).Error; createErr != nil {
		panic(createErr)
	}
}
