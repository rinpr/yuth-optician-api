package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	ID			int	`json:"id"`
	Name        string	`json:"name"`
	Author      string	`json:"author"`
	Publisher   string	`json:"publisher"`
	Description string	`json:"description"`
	Price       uint	`json:"price"`
}

func CreateBook(db *gorm.DB, book *Book) error {
	result := db.Create(&book)
	if result.Error != nil {
		// fmt.Printf("Error creating book: %v", result.Error)
		return result.Error
	}
	// fmt.Println("Create Book Successful")
	return nil
}

func GetBook(db *gorm.DB, id int) *Book {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		// log.Fatalf("Error getting book: %v", result.Error)
		fmt.Printf("Error getting book: %v", result.Error)
	}
	return &book
}

func GetAllBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		fmt.Printf("Error getting all books: %v", result.Error)
	}
	return books
}

func UpdateBook(db *gorm.DB, book *Book) error {
	result := db.Updates(&book)
	if result.Error != nil {
		// fmt.Printf("Error updating book: %v", result.Error)
		return result.Error
	}
	// fmt.Println("Update Book Successful")
	return nil
}

func DeleteBook(db *gorm.DB, id int) error {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		// fmt.Printf("Error deleting book: %v", result.Error)
		return result.Error
	}
	// fmt.Println("Delete Book Successful")
	return nil
}

func SearchBook(db *gorm.DB, bookName string) []Book {
	var books []Book
	result := db.Where("name = ?", bookName).Order("price").Find(&books)
	if result.Error != nil {
		fmt.Printf("Error searching book: %v", result.Error)
	}
	return books
}