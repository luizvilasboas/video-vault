package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title       string `json:"title" validate:"nonzero,nonnil"`
	Description string `json:"description" validate:"nonzero,nonnil"`
	URL         string `json:"url" validate:"nonzero,nonnil"`
	CategoryID  uint   `json:"category_id" validate:"nonzero,nonnil" gorm:"index;default:1"`
}

func ValidateVideoData(video *Video) error {
	return validator.Validate(video)
}
