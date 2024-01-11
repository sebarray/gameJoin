package user

import (
	"errors"
	"fmt"
	"gameJoin/internal/data/repository/user/dto"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (u User) Login(userDTO dto.User) (*dto.User, error) {
	dsn := os.Getenv("db") // e.g. "user:password@localhost:5432/dbname?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Default().Println(err.Error())
		return nil, err
	}

	var userDB dto.User
	if err := db.Where("email = ?", userDTO.Email).First(&userDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Default().Println(err.Error())
			return nil, fmt.Errorf("no such user")
		} else {
			log.Default().Println(err.Error())
			return nil, fmt.Errorf("database error")
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.PasswordHash), []byte(userDTO.PasswordHash)); err != nil {

		return nil, fmt.Errorf("wrong password")
	}

	return &userDB, nil
}
