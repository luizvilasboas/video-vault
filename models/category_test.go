package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryValidation(t *testing.T) {
	validCategory := &Category{
		Title:  "Valid Title",
		Color:  "Valid Color",
		Videos: []Video{},
	}

	invalidCategory := &Category{}

	t.Run("Valid Category", func(t *testing.T) {
		err := ValidateCategoryData(validCategory)
		assert.NoError(t, err, "Expected no validation error for a valid category")
	})

	t.Run("Invalid Category", func(t *testing.T) {
		err := ValidateCategoryData(invalidCategory)
		assert.Error(t, err, "Expected validation error for an invalid category")
	})
}
