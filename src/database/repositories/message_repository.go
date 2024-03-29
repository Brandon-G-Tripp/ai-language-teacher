package repositories

import (
	"errors"
	"log"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	database_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	"gorm.io/gorm"
)

var ErrMessageNotFound = errors.New("message not found")

type MessageRepository struct {
    db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
    return &MessageRepository{
        db: db,
    }
} 

func (r *MessageRepository) Create(message *database_models.Message) error {
    return r.db.Create(message).Error
}

func (r *MessageRepository) GetById(id uint) (*database_models.Message, error) {
    var message models.Message
    err := r.db.First(&message, id).Error
    log.Printf("Message: %+v", message)
    log.Printf("Error: %+v", err)

    if err == gorm.ErrRecordNotFound {
        log.Print("err is gorm not found")
        return &models.Message{}, ErrMessageNotFound
    } 

    return &message, err
} 

func (r *MessageRepository) GetByConversationId(conversationId uint) ([]*models.Message, error) {
    var messages []*models.Message
    err := r.db.Where("conversation_id = ?", conversationId).Find(&messages).Error
    return messages, err
}

func (r *MessageRepository) Update(message *models.Message) error {
    err := r.db.Save(message).Error
    return err
} 


func (r *MessageRepository) Delete(message *models.Message) error {
    err := r.db.Delete(message).Error
    return err
} 
