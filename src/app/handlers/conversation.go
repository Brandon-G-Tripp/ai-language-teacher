package handlers

import (
    "net/http"

    handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
    "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
    db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
    conversation_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
)

type ConversationHandler struct {
    ConversationRepo *conversation_repo.ConversationRepository
    AuthService *auth.AuthService
} 

func NewConversationHandler(repo *conversation_repo.ConversationRepository, authService *auth.AuthService) *ConversationHandler {
    return &ConversationHandler{
        ConversationRepo: repo,
        AuthService: authService,
    }
} 

func (h *ConversationHandler) CreateConversation(req handler_models.CreateConversationRequest) (*db_models.Conversation, error) {
    // implement create logic 
    return nil, nil
} 

func (h *ConversationHandler) GetConversation(conversationID uint) (*db_models.Conversation, error) {
    return nil, nil
}

func (h *ConversationHandler) UpdateConversation(conversationID uint, req handler_models.UpdateConversationRequest) (*db_models.Conversation, error) {
    return nil, nil
}

func (h *ConversationHandler) DeleteConversation(conversationID uint) error {
    return nil
} 
