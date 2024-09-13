package database

import (
	getdollarexchange "github.com/edsonjuniordev/client-server-challenge-fullcycle/server/get-dollar-exchange"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabase() (*gorm.DB, error) {
	if db == nil {
		db, err := gorm.Open(sqlite.Open("server/usd.db"), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		err = db.AutoMigrate(&getdollarexchange.USDBRL{})
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	return db, nil
}
