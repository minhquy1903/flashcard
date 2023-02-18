package repository

import (
	"context"

	"github.com/minhquy1903/flashcard-api/internal/auth/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int32) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.db.Table(user.TableName()).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Table(user.TableName()).Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Table(user.TableName()).Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
