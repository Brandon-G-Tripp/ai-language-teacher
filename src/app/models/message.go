package models

import (
    db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

type CreateMessageRequest struct {
    ConverstationID uint `json:"conversation_id"`
    UserID uint `json:"user_id"`
    Content string `json:"content"`
} 

type CreateMessageResponse struct {
    Message *db_models.Message `json:"message"`
} 

type GetMessagesByConversationIDRequest struct {
    ConversationID uint `json:"conversation_id"`
} 

type UpdateMessageRequest struct {
    Content string `json:"content"`
} 

type UpdateMessageResponse struct {
    Message *db_models.Message `json:"message"`
} 

type DeleteMessageResponse struct {
    Message string `json:"message"`
} 

