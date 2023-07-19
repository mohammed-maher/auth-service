package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect(connectionString string) error {
	instance, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	Instance = instance
	return err
}

func Migrate() error {
	return Instance.AutoMigrate(&User{})
}
