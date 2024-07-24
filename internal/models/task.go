package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
    ID              uint    `gorm:"primaryKey"`
    Name            string  `gorm:"size:255"`
    Description     string  `gorm:"type:text"`
    Completed       bool
    CreatedAt       time.Time
    UpdatedAt       time.Time
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&Task{})
}
