package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDatabase(drive string) *gorm.DB {
	dsn := ""
	dialector := gorm.Dialector(nil)

	switch drive {
	case "postgres":
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", Get().DBUsername, Get().DBPassword, Get().DBHost, Get().DBPort, Get().DBName)
		dialector = postgres.Open(dsn)
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Get().DBPassword, Get().DBUsername, Get().DBHost, Get().DBPort, Get().DBName)
		dialector = mysql.Open(dsn)
	default:
		panic("invalid drive")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}
