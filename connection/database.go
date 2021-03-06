package connection

import (
	"demo/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&entity.User{})

	DB = database
}
