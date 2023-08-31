package db

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mutex sync.Mutex
var db *gorm.DB

// model
type Tiny2LongURL struct {
	Tinyurl string `gorm:"primaryKey" json:"tinyurl"`
	Longurl string `json:"longurl"`
}

func ConnectToDB() {
	defer mutex.Unlock()
	mutex.Lock()
	if db != nil {
		return
	}
	dsn := "admin:password@/tinyurl?"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect DB")
	}
	database.AutoMigrate(&Tiny2LongURL{})
	db = database
	log.Println("connected to databse")

}

func GetDB() *gorm.DB {
	return db
}

func AddURL(record *Tiny2LongURL) *Tiny2LongURL {
	defer mutex.Unlock()
	mutex.Lock()
	result := db.Create(record)
	if result.RowsAffected == 0 {
		return nil
	}
	return record
}

func GetFullURL(tinyurl string) *Tiny2LongURL {
	defer mutex.Lock()
	mutex.Lock()
	var url Tiny2LongURL
	db.Where("tinyurl = ?", tinyurl).Find(&url)
	return &url
}
