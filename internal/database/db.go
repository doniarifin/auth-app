package database

import (
	"auth-app/internal/model"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	appenv := os.Getenv("APP_ENV")
	// for sqlserver
	// dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
	// 	user, pass, host, port, name,
	// )
	// db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	//for postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, pass, dbname, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//run migrate
	stingconn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, dbname,
	)

	wd, _ := os.Getwd()
	sourceURL := fmt.Sprintf("file://%s/../internal/migrations", wd)

	if appenv == "local" {
		sourceURL = fmt.Sprintf("file://%s/../internal/migrations", wd)
	}

	m, err := migrate.New(
		sourceURL,
		stingconn)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("db connected!")
	return db, nil
}

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
