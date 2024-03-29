package repositories

import (
    "errors"
    "log"

    "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
    "gorm.io/gorm"
)

var ErrConversationNotFound = errors.New("conversation not found")

type ConversationRepository struct {
    db *gorm.DB
} 

func NewConversationRepository(db *gorm.DB) *ConversationRepository {
    return &ConversationRepository{
        db: db,
    }
} 

func (r *ConversationRepository) Create(conversation *models.Conversation) error {
    return r.db.Create(conversation).Error
} 

func (r *ConversationRepository) GetById(id uint) (*models.Conversation, error) {
    var conversation models.Conversation
    err := r.db.First(&conversation, id).Error
    log.Printf("Conversation: %+v", conversation)
    log.Printf("Error: %+v", err)

    if err == gorm.ErrRecordNotFound {
        log.Print("err is gorm not found")
        return &models.Conversation{}, ErrConversationNotFound
    } 

    return &conversation, err
} 

func (r *ConversationRepository) GetByUserId(userId uint) ([]*models.Conversation, error) {
    var conversations []*models.Conversation
    err := r.db.Where("user_id =?", userId).Find(&conversations).Error
    return conversations, err
}

func (r *ConversationRepository) Update(conversation *models.Conversation) error {
    err := r.db.Save(conversation).Error
    return err
}

func (r *ConversationRepository) Delete(conversation *models.Conversation) error {
    err := r.db.Delete(conversation).Error
    return err
}
