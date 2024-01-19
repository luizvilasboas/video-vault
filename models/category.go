package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title  string  `json:"title" validate:"nonzero,nonnil"`
	Color  string  `json:"color" validate:"nonzero,nonnil"`
	Videos []Video `json:"videos"`
}

func (c *Category) BeforeDelete(tx *gorm.DB) (err error) {
	var videos []Video

	tx.Model(&c).Association("Videos").Find(&videos)
	tx.Model(&videos).Update("CategoryID", 1)

	return nil
}

func ValidateCategoryData(category *Category) error {
	if err := validator.Validate(category); err != nil {
		return err
	}

	return nil
}
