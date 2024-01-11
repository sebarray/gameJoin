package user

import (
	"fmt"
	"gameJoin/pkg/domain/user"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (u User) CreateUserRegistry(userRegistry user.UserRegistryRequest) (user.UserRegistryRequest, error) {
	dsn := os.Getenv("db")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return user.UserRegistryRequest{}, err
	}
	// AutoMigrate will ONLY create tables, missing columns and missing indexes, and WON’T change existing column’s types or delete unused columns
	db.AutoMigrate(&user.UserRegistryRequest{})

	if err := db.Create(userRegistry).Error; err != nil {
		return user.UserRegistryRequest{}, fmt.Errorf("failed to create user registry: %w", err)
	}

	return userRegistry, nil
}
