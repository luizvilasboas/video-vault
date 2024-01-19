package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/olooeez/video-vault/models"
)

func TestValidateVideoValid(t *testing.T) {
	t.Run("Valid Video", func(t *testing.T) {
		video := &models.Video{
			Title:       "Sample Title",
			Description: "Sample Description",
			URL:         "http://example.com",
			CategoryID:  1,
		}

		err := models.ValidateVideoData(video)

		assert.NoError(t, err, "No validation error expected for a valid video")
	})
}

func TestValidateVideoInvalid(t *testing.T) {
	t.Run("Invalid Video - Missing Title", func(t *testing.T) {
		video := &models.Video{
			Description: "Sample Description",
			URL:         "http://example.com",
			CategoryID:  1,
		}

		err := models.ValidateVideoData(video)

		assert.Error(t, err, "Validation error expected for a video with missing title")
		assert.Contains(t, err.Error(), "Title", "Error message should mention the missing Title field")
	})

	t.Run("Invalid Video - Missing URL", func(t *testing.T) {
		video := &models.Video{
			Title:       "Sample Title",
			Description: "Sample Description",
			CategoryID:  1,
		}

		err := models.ValidateVideoData(video)

		assert.Error(t, err, "Validation error expected for a video with missing URL")
		assert.Contains(t, err.Error(), "URL", "Error message should mention the missing URL field")
	})

	t.Run("Invalid Video - Missing CategoryID", func(t *testing.T) {
		video := &models.Video{
			Title:       "Sample Title",
			Description: "Sample Description",
			URL:         "http://example.com",
		}

		err := models.ValidateVideoData(video)

		assert.Error(t, err, "Validation error expected for a video with missing CategoryID")
		assert.Contains(t, err.Error(), "CategoryID", "Error message should mention the missing CategoryID field")
	})
}
