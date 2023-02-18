package db

import "gorm.io/gorm"

type SQL struct {
	orm *gorm.DB
}

func NewSQLStore(orm *gorm.DB) *SQL {
	return &SQL{orm: orm}
}

func (s *SQL) GetORM() *gorm.DB {
	return s.orm
}
