package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"yuth-optician-api/models"
	"yuth-optician-api/router"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	loadENV()

	db := initDB()

	r := router.NewRouter(db)
	r.Listen(":8080")


	// book := models.GetBook(db, 1)
	// fmt.Println(book)

	// book.Price = 9999
	// models.UpdateBook(db, book)

	// models.DeleteBook(db, 1)

	// models.CreateBook(db, &models.Book{
	// 	Name: "7 Habits of highly effective people",
	// 	Author: "Stephen R. Covy",
	// 	Publisher: "dmb Books",
	// 	Description: "",
	// 	Price: 899,
	// })


	// currentBook := models.SearchBook(db, "7 Habits of highly effective people")
	// for _, v := range currentBook {
	// 	fmt.Println(v.ID, v.Name, v.Author, v.Price)
	// }
	// fmt.Println(currentBook)
}

func loadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}
}

func initDB() *gorm.DB {
	// load env
	host, user, password, dbname := os.Getenv("POSTGRES_HOST"), 
	os.Getenv("POSTGRES_USER"), 
	os.Getenv("POSTGRES_PASSWORD"),
	os.Getenv("POSTGRES_DB")
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic("INVALID PORT FROM .ENV")
	}

	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Logger config for gorm
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			//IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			//ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema
	// db.Migrator().DropColumn(&models.Book{}, "name")
	db.AutoMigrate(&models.Book{}, models.User{})
	fmt.Println("Database migration completed!")
	return db
}