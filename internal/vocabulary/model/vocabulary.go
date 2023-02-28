package model

import common "github.com/minhquy1903/flashcard-api/internal/common/model"

type Vocabulary struct {
	common.BaseModel
	NewWord     string
	UserId      int32
	Status      int8
	CountStatus int8
	Image       string
	Meaning     string
	Sequence    string
	Topic       string
}

func (d Vocabulary) TableName() string {
	return "vocabulary"
}
