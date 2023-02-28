package repository

import (
	"context"

	"github.com/minhquy1903/flashcard-api/internal/auth/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id int32) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
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

func (r *userRepository) FindByID(ctx context.Context, id int32) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Table(user.TableName()).Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Table(user.TableName()).Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
