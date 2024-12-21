package models

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func CreateUser(db *gorm.DB, user *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing password: %v", err)
		return err
	}
		
	user.Password = string(hash)
	result := db.Create(user)
	if result.Error != nil {
		fmt.Printf("Error creating user: %v", result.Error)
		return result.Error
	}
	return nil		
}

func LoginUser(db *gorm.DB, user *User) (string, error) {
	selectedUser := new(User)
	// find user by email
	result := db.Where("email = ?", user.Email).First(selectedUser)
	if result.Error != nil {
		return fmt.Sprintf("Error logging in: %v", result.Error), result.Error
	}

	// check password
	err := bcrypt.CompareHashAndPassword(
		[]byte(selectedUser.Password), 
		[]byte(user.Password),
	)
	if err != nil {
		return fmt.Sprintf("Invalid password: %v", result.Error), result.Error
	}

	// Create JWT token
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY") // do not leak this

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
  
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
	  return fmt.Sprintf("invalid tokens: %v", err), err
	}
	
	return t, nil
}