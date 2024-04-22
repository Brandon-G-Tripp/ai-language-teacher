package handlers

import (
	"net/http"

	handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	message_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
)

type MessageHandler struct {
    MessageRepo *message_repo.MessageRepository
    AuthService *auth.AuthService
} 

func NewMessageHandler(
    repo *message_repo.MessageRepository,
    authService *auth.AuthService,
) *MessageHandler {
    return &MessageHandler{
        MessageRepo: repo,
        AuthService: authService,
    }
} 

func (h *MessageHandler) CreateMessage(req handler_models.CreateMessageRequest) (*db_models.Message, error) {
    if req.ConversationID <= 0 {
        return nil, handler_models.ApiError{
            Message: "Conversation ID is required",
            Code:    http.StatusBadRequest,
        }
    }
    if req.UserID <= 0 {
        return nil, handler_models.ApiError{
            Message: "User ID is required",
            Code:    http.StatusBadRequest,
        }
    }
    if req.Content == "" {
        return nil, handler_models.ApiError{
            Message: "Message content is required",
            Code:    http.StatusBadRequest,
        }
    }

    message := &db_models.Message{
        ConversationID: req.ConversationID,
        UserID:         req.UserID,
        Content:        req.Content,
    }

    if err := h.MessageRepo.Create(message); err != nil {
        return nil, err
    }

    return message, nil
}

func (h *MessageHandler) GetMessagesByConversationID(conversationID uint) ([]*db_models.Message, error) {
    messages, err := h.MessageRepo.GetByConversationID(conversationID)
    if err != nil {
        if err == message_repo.ErrConversationNotFound {
            return nil, handler_models.ApiError{
                Message: "Conversation not found",
                Code:    http.StatusNotFound,
            }
        }
        return nil, handler_models.ApiError{
            Message: "Failed to retrieve messages",
            Code:    http.StatusInternalServerError,
        }
    }

    return messages, nil
}

func (h *MessageHandler) UpdateMessage(messageID uint, req handler_models.UpdateMessageRequest) (*db_models.Message, error) {
    message, err := h.MessageRepo.GetByID(messageID)
    if err != nil {
        if err == message_repo.ErrMessageNotFound {
            return nil, handler_models.ApiError{
                Message: "Message not found",
                Code:    http.StatusNotFound,
            }
        }
        return nil, handler_models.ApiError{
            Message: "Failed to retrieve the message",
            Code:    http.StatusInternalServerError,
        }
    }

    message.Content = req.Content

    if err := h.MessageRepo.Update(message); err != nil {
        return nil, handler_models.ApiError{
            Message: "Failed to update message",
            Code:    http.StatusInternalServerError,
        }
    }

    return message, nil
}

func (h *MessageHandler) DeleteMessage(messageID uint) (*db_models.Message, error) {
    message, err := h.MessageRepo.GetByID(messageID)
    if err != nil {
        if err == message_repo.ErrMessageNotFound {
            return nil, handler_models.ApiError{
                Message: "Message not found",
                Code:    http.StatusNotFound,
            }
        }
        return nil, handler_models.ApiError{
            Message: "Failed to retrieve the message",
            Code:    http.StatusInternalServerError,
        }
    }

    if err := h.MessageRepo.Delete(message); err != nil {
        return nil, handler_models.ApiError{
            Message: "Failed to delete message",
            Code:    http.StatusInternalServerError,
        }
    }

    return message, nil
}
