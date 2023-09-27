package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	post "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	log.Println("create pool database connection")

	dbURL := os.Getenv("DATABASE_URL")
	sqlDb, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Panicln("failed to connect database", err)
		return nil
	}

	sqlDb.SetConnMaxIdleTime(30)
	sqlDb.SetMaxOpenConns(50)
	sqlDb.SetConnMaxLifetime(2 * time.Minute)

	log.Println("pool database connection is created")

	ormDb, err := gorm.Open(post.New(post.Config{
		Conn: sqlDb,
	}), &gorm.Config{})

	if err != nil {
		log.Println("error on creating gorm connection")
		return nil
	}

	log.Println("gorm connection is created")

	return ormDb
}
