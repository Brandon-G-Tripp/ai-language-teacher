package models

import "time"

type User struct {
    ID uint `gorm:"primaryKey"`
    Name string `gorm:"not null;size:255"`
    Email string `gorm:"uniqueIndex;not null;size:255"`
    Password string `gorm:"not null;size:255"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
} 
