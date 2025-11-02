package models

import (
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Модель user
type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username" binding:"max=255"`
	Password string `gorm:"size:255;not null;" json:"password" binding:"min=8;max=255"`
}

// Хэширование пароля
func (user *User) HashPassword() error {
	// Получение хэша
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Ошибка хэширования пароля")
		// return err
	}

	// Запись хэша
	user.Password = string(hashedPassword)

	return nil
}

func (user *User) HtmlEscapeUsername() {
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
}
