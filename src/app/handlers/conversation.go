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
    if req.UserID <= 0 {
        return nil, handler_models.ApiError{
            Message: "User ID is required",
            Code: http.StatusBadRequest,
        }
    }
    conversation := &db_models.Conversation{
        UserID: req.UserID, 
        Title: req.Title, 
    }

    if err := h.ConversationRepo.Create(conversation); err != nil {
        return nil, err
    } 

    return conversation, nil
} 

func (h *ConversationHandler) GetConversation(conversationID uint) (*db_models.Conversation, error) {
    conversation, err := h.ConversationRepo.GetById(conversationID)
    if err != nil {
        if err == conversation_repo.ErrConversationNotFound {
            return nil, handler_models.ApiError{
                Message: "Conversation not found", 
                Code: http.StatusNotFound,
            } 
        } 
        return nil, handler_models.ApiError{
            Message: "Failed to retrieve the conversation",
            Code: http.StatusInternalServerError,
        }
    } 

    return conversation, nil
}

func (h *ConversationHandler) UpdateConversation(conversationID uint, req handler_models.UpdateConversationRequest) (*db_models.Conversation, error) {
    conversation, err := h.ConversationRepo.GetById(conversationID)
    if err != nil {
        if err == conversation_repo.ErrConversationNotFound {
            return nil, handler_models.ApiError{
                Message: "Conversation not found", 
                Code: http.StatusNotFound,
            } 
        } 
        return nil, handler_models.ApiError{
            Message: "Failed to retrieve the conversation",
            Code: http.StatusInternalServerError,
        }
    } 

    conversation.Title = req.Title 

    if err := h.ConversationRepo.Update(conversation); err != nil {
        return nil, handler_models.ApiError{
            Message: "Failed to update conversation",
            Code: http.StatusInternalServerError, 
        }
    } 

    return conversation, nil
}

func (h *ConversationHandler) DeleteConversation(conversationID uint) (*db_models.Conversation, error) {
    conversation, err := h.ConversationRepo.GetById(conversationID)
    if err != nil {
        if err == conversation_repo.ErrConversationNotFound {
            return nil, handler_models.ApiError{
                Message: "Conversation not found", 
                Code: http.StatusNotFound,
            } 
        } 
        return nil, handler_models.ApiError{
            Message: "Failed to retrieve the conversation",
            Code: http.StatusInternalServerError,
        }
    } 

    if err := h.ConversationRepo.Delete(conversation); err != nil {
        return nil, handler_models.ApiError{
            Message: "Failed to delete conversation",
            Code: http.StatusInternalServerError, 
        }
    } 

    return conversation, nil
} 
