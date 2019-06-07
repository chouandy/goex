package gormex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAttachment(t *testing.T) {
	HashSecret = []byte("hash_secret")
	updatedAt, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")

	// Set test cases
	testCases := []struct {
		attachment *Attachment
		expected   struct {
			updatedAtTimestampString string
			fileNameHashKey          string
			hashedFileName           string
			fileExtension            string
			path                     string
		}
	}{
		{
			attachment: &Attachment{
				ResourceName: "Activity",
				ResourceID:   1,
				Name:         "Poster",
				FileName:     "poster.jpg",
				UpdatedAt:    updatedAt,
			},
			expected: struct {
				updatedAtTimestampString string
				fileNameHashKey          string
				hashedFileName           string
				fileExtension            string
				path                     string
			}{
				updatedAtTimestampString: "1546300800",
				fileNameHashKey:          "activities/posters/1/original/1546300800",
				hashedFileName:           "28a8cae489a94085bac99294f61732bb939f67df",
				fileExtension:            ".jpg",
				path:                     "activities/posters/1/original/28a8cae489a94085bac99294f61732bb939f67df.jpg",
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected.updatedAtTimestampString, testCase.attachment.GetUpdatedAtTimestampString())
			assert.Equal(t, testCase.expected.fileNameHashKey, testCase.attachment.GetFileNameHashKey())
			assert.Equal(t, testCase.expected.hashedFileName, testCase.attachment.GetHashedFileName())
			assert.Equal(t, testCase.expected.fileExtension, testCase.attachment.GetFileExtension())
			assert.Equal(t, testCase.expected.path, testCase.attachment.GetPath())
		})
	}
}
