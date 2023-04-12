package models

import (
	"html"
	"strings"

	"RESTAPI_Gin/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User struct contains gorm.Model struct along with Username and Password both of which are strings and Entries which specify 1:n relationship
type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"` //json binding for password is -. This ensures that it is not returned in json response.
	Entries  []Entry
}

func (user *User) Save() (*User, error) {
	err := db.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(db2 *gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (*User, error) {
	var user User
	err := db.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func FindUserById(userID uint) (*User, error) {
	var user User
	err := db.Database.Preload("Entries").Where("ID=?", userID).Find(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}
