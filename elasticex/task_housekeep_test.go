package elasticex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHousekeepGetHousekeepDate(t *testing.T) {
	// Set test cases
	testCases := []struct {
		time      time.Time
		retension int
		expected  string
	}{
		{
			time:      time.Date(2018, 5, 23, 0, 0, 0, 0, time.UTC),
			retension: 30,
			expected:  "2018.04.22",
		},
		{
			time:      time.Date(2017, 5, 23, 0, 0, 0, 0, time.UTC),
			retension: 30,
			expected:  "2017.04.22",
		},
		{
			time:      time.Date(2018, 3, 1, 0, 0, 0, 0, time.UTC),
			retension: 30,
			expected:  "2018.01.29",
		},
		{
			time:      time.Date(2018, 1, 31, 0, 0, 0, 0, time.UTC),
			retension: 30,
			expected:  "2017.12.31",
		},
		{
			time:      time.Date(2018, 4, 30, 0, 0, 0, 0, time.UTC),
			retension: 30,
			expected:  "2018.03.30",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			task := &HousekeepTask{
				Time:      testCase.time,
				Retention: testCase.retension,
			}
			housekeepDate := task.Time.AddDate(0, 0, -(task.Retention + 1)).Format("2006.01.02")
			assert.Equal(t, testCase.expected, housekeepDate)
		})
	}
}
