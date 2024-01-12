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
}

func ValidateVideoData(video *Video) error {
	if err := validator.Validate(video); err != nil {
		return err
	}

	return nil
}
