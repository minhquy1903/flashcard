package dto

import "github.com/minhquy1903/flashcard-api/internal/vocabulary/model"

type CreateVocabularyInput struct {
	NewWord  string
	UserId   int32
	Image    string
	Meaning  string
	Sequence string
	Topic    string
}

func (d CreateVocabularyInput) ToModel() model.Vocabulary {
	return model.Vocabulary{
		NewWord:  d.NewWord,
		Meaning:  d.Meaning,
		UserId:   d.UserId,
		Image:    d.Image,
		Topic:    d.Topic,
		Sequence: d.Sequence,
	}
}
