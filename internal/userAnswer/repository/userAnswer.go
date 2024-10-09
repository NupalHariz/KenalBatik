package repository

import (
	"context"
	"kenalbatik-be/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userAnswerRepository struct {
	db *gorm.DB
}

type UserAnswerRepositoryInterface interface{
	Create(ctx context.Context, userAnswer *domain.UserAnswer) error
	FindUserAnswer(ctx context.Context, userAnswer *[]domain.UserAnswer, userId uuid.UUID) error
}

func NewUserAnswerRepository(db *gorm.DB) UserAnswerRepositoryInterface {
	return &userAnswerRepository{db}
}

func (r *userAnswerRepository) Create(ctx context.Context, userAnswer *domain.UserAnswer) error {
	err := r.db.WithContext(ctx).Create(userAnswer).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userAnswerRepository) FindUserAnswer(ctx context.Context, userAnswer *[]domain.UserAnswer, userId uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(userAnswer).Error
	if err != nil {
		if len(*userAnswer) == 0 {
			return domain.ErrRecordNotFound
		}

		return err
	}

	return nil
}
