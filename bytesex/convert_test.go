package bytesex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	// Set test cases
	testCases := []struct {
		b        uint64
		expected struct {
			KiB float64
			MiB float64
			GiB float64
			TiB float64
		}
	}{
		{
			b: 1099511627776,
			expected: struct {
				KiB float64
				MiB float64
				GiB float64
				TiB float64
			}{
				KiB: 1073741824,
				MiB: 1048576,
				GiB: 1024,
				TiB: 1,
			},
		},
		{
			b: 1073741824,
			expected: struct {
				KiB float64
				MiB float64
				GiB float64
				TiB float64
			}{
				KiB: 1048576,
				MiB: 1024,
				GiB: 1,
				TiB: 0,
			},
		},
		{
			b: 1048576,
			expected: struct {
				KiB float64
				MiB float64
				GiB float64
				TiB float64
			}{
				KiB: 1024,
				MiB: 1,
				GiB: 0,
				TiB: 0,
			},
		},
		{
			b: 1024,
			expected: struct {
				KiB float64
				MiB float64
				GiB float64
				TiB float64
			}{
				KiB: 1,
				MiB: 0,
				GiB: 0,
				TiB: 0,
			},
		},
		{
			b: 130000150000,
			expected: struct {
				KiB float64
				MiB float64
				GiB float64
				TiB float64
			}{
				KiB: 126953271.48,
				MiB: 123977.80,
				GiB: 121.07,
				TiB: 0.12,
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected.KiB, ToKiB(testCase.b))
			assert.Equal(t, testCase.expected.MiB, ToMiB(testCase.b))
			assert.Equal(t, testCase.expected.GiB, ToGiB(testCase.b))
			assert.Equal(t, testCase.expected.TiB, ToTiB(testCase.b))
		})
	}
}
