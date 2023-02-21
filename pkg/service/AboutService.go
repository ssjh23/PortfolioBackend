package service

import (
	"gorm.io/gorm"
)
type AboutService struct {
	db *gorm.DB
}

func NewAboutService(db *gorm.DB) *AboutService {
	return &AboutService{
		db: db,
	}
}
