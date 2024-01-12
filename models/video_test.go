package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/olooeez/video-vault/models"
)

func TestValidateVideoValid(t *testing.T) {
	video := &models.Video{
		Title:       "Sample Title",
		Description: "Sample Description",
		URL:         "http://example.com",
	}

	err := models.ValidateVideoData(video)

	assert.NoError(t, err, "Validation error not expected")
}

func TestValidateVideoEmpty(t *testing.T) {
	video := &models.Video{}

	err := models.ValidateVideoData(video)

	assert.Error(t, err, "Validation error expected")
	assert.Contains(t, err.Error(), "zero value", "Validation error 'zero value' message expected")
}
