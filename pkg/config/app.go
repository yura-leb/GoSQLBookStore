package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("postgres", "yura-leb:1234/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("postgres", "host='localhost' port=5432 user=postgres dbname='postgres' password='1234' sslmode=disable")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
