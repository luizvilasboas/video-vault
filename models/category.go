package models

import (
	"fmt"

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

	if err := tx.Model(&c).Association("Videos").Find(&videos); err != nil {
		return fmt.Errorf("erro ao buscar vídeos associados à categoria: %v", err)
	}

	tx.Model(&videos).Update("CategoryID", 1)

	return nil
}

func ValidateCategoryData(category *Category) error {
	return validator.Validate(category)
}
