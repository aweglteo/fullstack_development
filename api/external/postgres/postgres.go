package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var GormDB *gorm.DB

func main() *gorm.DB {
	DBMS := "postgres"
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USERNAME := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	PROTOCOL := "tcp(" + HOST + ":" + PORT + ")"

	CONNECT := USERNAME + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME

	GormDB, err := gorm.Open(PROTOCOL, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return GormDB
}
