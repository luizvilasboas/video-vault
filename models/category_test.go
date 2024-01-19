package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryValidation(t *testing.T) {
	t.Run("Valid Category", func(t *testing.T) {
		validCategory := &Category{
			Title:  "Valid Title",
			Color:  "Valid Color",
			Videos: []Video{},
		}

		err := ValidateCategoryData(validCategory)

		assert.NoError(t, err, "No validation error expected for a valid category")
	})

	t.Run("Invalid Category - Missing Title", func(t *testing.T) {
		invalidCategory := &Category{
			Color:  "Valid Color",
			Videos: []Video{},
		}

		err := ValidateCategoryData(invalidCategory)

		assert.Error(t, err, "Expected validation error for a category with missing title")
		assert.Contains(t, err.Error(), "Title", "Error message should mention the missing Title field")
	})

	t.Run("Invalid Category - Missing Color", func(t *testing.T) {
		invalidCategory := &Category{
			Title:  "Valid Title",
			Videos: []Video{},
		}

		err := ValidateCategoryData(invalidCategory)

		assert.Error(t, err, "Expected validation error for a category with missing color")
		assert.Contains(t, err.Error(), "Color", "Error message should mention the missing Color field")
	})
}
