package models

import (
    db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

type CreateConversationRequest struct {
    UserID uint `json:"user_id"`
    Title string `json:"title"`
} 

type CreateConversationResponse struct {
    Conversation *db_models.Conversation `json:"conversation"`
} 

type GetConversationRequest struct {
    Conversation *db_models.Conversation `json:"conversation"`
} 

type GetConversationResponse struct {
    Conversation *db_models.Conversation `json:"conversation"`
} 

type UpdateConversationRequest struct {
    Title string `json:"title"`
} 

type UpdateConversationResponse struct {
    Conversation *db_models.Conversation `json:"conversation"`
} 

type DeleteConversationResponse struct {
    Message string `json:"message"`
} 

