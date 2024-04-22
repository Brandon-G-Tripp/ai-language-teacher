package models

import "time"

type Message struct {
    ID uint `gorm:"primaryKey"`
    ConversationID uint `gorm:"not null"`
    UserID uint `gorm:"not null"`
    Content string `gorm:"type:text;not null"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
