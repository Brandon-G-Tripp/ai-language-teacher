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

func (r *UserRepository) GetById(id uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    return &user, err
} 

func (r *UserRepository) Update(user *models.User) error {
    err := r.db.Save(user).Error
    return err
}

func (r *UserRepository) Delete(user *models.User) error {
    err := r.db.Delete(user).Error
    return err
} 

func (r *UserRepository) GetAll() ([]*models.User, error) {
    var users []*models.User
    err := r.db.Find(&users).Error
    return users, err
} 

