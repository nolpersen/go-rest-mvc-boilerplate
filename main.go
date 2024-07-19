package main

import (
	"github.com/nolpersen/src/config"
	"github.com/nolpersen/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDb(db)

	routes.Routes()
}
