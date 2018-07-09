package timex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBeginningOfDay(t *testing.T) {
	// Set test cases
	testCases := []struct {
		datetime string
		expected string
	}{
		{
			datetime: "2018-07-09T05:40:58.321957447Z",
			expected: "2018-07-09T00:00:00Z",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			datetime, err := time.Parse(time.RFC3339Nano, testCase.datetime)
			assert.Nil(t, err)
			datetime = BeginningOfDay(datetime)
			assert.Equal(t, testCase.expected, datetime.Format(time.RFC3339Nano))
		})
	}
}

func TestEndOfDay(t *testing.T) {
	// Set test cases
	testCases := []struct {
		datetime string
		expected string
	}{
		{
			datetime: "2018-07-09T05:40:58.321957447Z",
			expected: "2018-07-09T23:59:59.999999999Z",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			datetime, err := time.Parse(time.RFC3339Nano, testCase.datetime)
			assert.Nil(t, err)
			datetime = EndOfDay(datetime)
			assert.Equal(t, testCase.expected, datetime.Format(time.RFC3339Nano))
		})
	}
}
