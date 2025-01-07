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

	const handleToUpdate = "handle-goes-here"
	const newDID = "did:plc:newstuff"

	if updateErr := database.Db().Model(&database.BskyHandle{}).Where("handle = ?", handleToUpdate).Update("did", newDID).Error; updateErr != nil {
		panic(updateErr)
	}
}
