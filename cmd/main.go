package main

import (
	"github.com/nadeem-baig/MHPS-backend/cmd/api"
	"github.com/nadeem-baig/MHPS-backend/db"
	"github.com/nadeem-baig/MHPS-backend/utils/logger"
)

func main() {

	db, err := db.ConnectDB()
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
		return
	}
	api.StartServer(db)
}
