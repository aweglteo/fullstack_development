package postgres

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

var GormDB *gorm.DB

func Connect() *gorm.DB {
	loadEnv()
	PLOTOCOL := "postgres"
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USERNAME := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := "host=" + HOST + " port=" + PORT + " user=" + USERNAME + " dbname=" + DBNAME + " password=" + PASSWORD + " sslmode=disable"

	db, err := gorm.Open(PLOTOCOL, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	GormDB = db
	return GormDB
}

func CloseConn() {
	GormDB.Close()
}
