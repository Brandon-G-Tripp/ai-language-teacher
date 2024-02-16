package repositories

import (
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
} 

func (r *UserRepository) Create(user *models.User) error {
    // TODO implement logic 
    return r.db.Create(user).Error
} 

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    return &user, err
} 


