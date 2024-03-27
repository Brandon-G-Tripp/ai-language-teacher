package models

import "time"

type Conversation struct {
    ID uint `gorm:"primaryKey"`
    UserID uint `gorm:"not null"`
    Title string `gorm:"type:varchar(255)"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
} 
