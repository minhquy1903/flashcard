package repository

import (
	"context"

	"github.com/minhquy1903/flashcard-api/internal/vocabulary/model"
	"github.com/minhquy1903/flashcard-api/internal/vocabulary/presenter"
	"gorm.io/gorm"
)

type VocabRepository interface {
	Create(c context.Context, vocabulary *model.Vocabulary) error
	// FindByID(c context.Context, id int32) (model.Vocabulary, error)
	GetList(c context.Context, condition presenter.GetListRequest) (interface{}, error)
}

type vocabRepository struct {
	db *gorm.DB
}

func NewVocabRepository(db *gorm.DB) *vocabRepository {
	return &vocabRepository{db: db}
}

func (r *vocabRepository) Create(c context.Context, vocab *model.Vocabulary) error {
	if err := r.db.Table(vocab.TableName()).Create(vocab).Error; err != nil {
		return err
	}

	return nil
}

func (r *vocabRepository) GetList(c context.Context, condition presenter.GetListRequest) (interface{}, error) {

	condition.Limit = 10
	condition.Page = 1

	var vocabs []model.Vocabulary

	offset := condition.Limit * (condition.Page - 1)

	qb := r.db.Table(model.Vocabulary{}.TableName())

	if condition.Topic != "" {
		qb.Where("topic = ?", condition.Topic)
	}

	if err := qb.Offset(offset).Limit(condition.Limit).Find(&vocabs).Error; err != nil {
		return nil, err
	}

	return vocabs, nil
}
