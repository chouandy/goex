package stringsex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifferences(t *testing.T) {
	// Set test cases
	testCases := []struct {
		sliceA   []string
		sliceB   []string
		expected struct {
			differenceA  []string
			intersection []string
			differenceB  []string
		}
	}{
		{
			sliceA: []string{},
			sliceB: []string{},
			expected: struct {
				differenceA  []string
				intersection []string
				differenceB  []string
			}{
				[]string{},
				[]string{},
				[]string{},
			},
		},
		{
			sliceA: []string{"a", "b", "c", "d", "e", "f"},
			sliceB: []string{},
			expected: struct {
				differenceA  []string
				intersection []string
				differenceB  []string
			}{
				[]string{"a", "b", "c", "d", "e", "f"},
				[]string{},
				[]string{},
			},
		},
		{
			sliceA: []string{},
			sliceB: []string{"a", "b", "c", "d", "e", "f"},
			expected: struct {
				differenceA  []string
				intersection []string
				differenceB  []string
			}{
				[]string{},
				[]string{},
				[]string{"a", "b", "c", "d", "e", "f"},
			},
		},
		{
			sliceA: []string{"a", "b", "c", "d", "e", "f"},
			sliceB: []string{"a", "b", "c", "d", "e", "f"},
			expected: struct {
				differenceA  []string
				intersection []string
				differenceB  []string
			}{
				[]string{},
				[]string{"a", "b", "c", "d", "e", "f"},
				[]string{},
			},
		},
		{
			sliceA: []string{"a", "b", "c"},
			sliceB: []string{"d", "e", "f"},
			expected: struct {
				differenceA  []string
				intersection []string
				differenceB  []string
			}{
				[]string{"a", "b", "c"},
				[]string{},
				[]string{"d", "e", "f"},
			},
		},
		{
			sliceA: []string{"a", "b", "c", "d"},
			sliceB: []string{"c", "d", "e", "f"},
			expected: struct {
				differenceA  []string
				intersection []string
				differenceB  []string
			}{
				[]string{"a", "b"},
				[]string{"c", "d"},
				[]string{"e", "f"},
			},
		},
		{
			sliceA: []string{"a", "b", "c", "d", "e"},
			sliceB: []string{"b", "c", "d", "e", "f"},
			expected: struct {
				differenceA  []string
				intersection []string
				differenceB  []string
			}{
				[]string{"a"},
				[]string{"b", "c", "d", "e"},
				[]string{"f"},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			differenceA, intersection, differenceB := Differences(testCase.sliceA, testCase.sliceB)
			assert.Equal(t, testCase.expected.differenceA, differenceA)
			assert.Equal(t, testCase.expected.intersection, intersection)
			assert.Equal(t, testCase.expected.differenceB, differenceB)
		})
	}
}

func TestIntersection(t *testing.T) {
	// Set test cases
	testCases := []struct {
		sliceA   []string
		sliceB   []string
		expected struct {
			intersection []string
		}
	}{
		{
			sliceA: []string{},
			sliceB: []string{},
			expected: struct {
				intersection []string
			}{
				[]string{},
			},
		},
		{
			sliceA: []string{"a", "b", "c", "d", "e", "f"},
			sliceB: []string{},
			expected: struct {
				intersection []string
			}{
				[]string{},
			},
		},
		{
			sliceA: []string{},
			sliceB: []string{"a", "b", "c", "d", "e", "f"},
			expected: struct {
				intersection []string
			}{
				[]string{},
			},
		},
		{
			sliceA: []string{"a", "b", "c", "d", "e", "f"},
			sliceB: []string{"a", "b", "c", "d", "e", "f"},
			expected: struct {
				intersection []string
			}{
				[]string{"a", "b", "c", "d", "e", "f"},
			},
		},
		{
			sliceA: []string{"a", "b", "c"},
			sliceB: []string{"d", "e", "f"},
			expected: struct {
				intersection []string
			}{
				[]string{},
			},
		},
		{
			sliceA: []string{"a", "b", "c", "d"},
			sliceB: []string{"c", "d", "e", "f"},
			expected: struct {
				intersection []string
			}{
				[]string{"c", "d"},
			},
		},
		{
			sliceA: []string{"a", "b", "c", "d", "e"},
			sliceB: []string{"b", "c", "d", "e", "f"},
			expected: struct {
				intersection []string
			}{
				[]string{"b", "c", "d", "e"},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			intersection := Intersection(testCase.sliceA, testCase.sliceB)
			assert.Equal(t, testCase.expected.intersection, intersection)
		})
	}
}

func TestSplitToInt32(t *testing.T) {
	// Set test cases
	testCases := []struct {
		s        string
		expected []int32
	}{
		{
			s:        "1,2,3,4",
			expected: []int32{1, 2, 3, 4},
		},
		{
			s:        "1,a,3,4",
			expected: []int32{1, 0, 3, 4},
		},
		{
			s:        "1,2,a,4",
			expected: []int32{1, 2, 0, 4},
		},
		{
			s:        "1,a,b,4",
			expected: []int32{1, 0, 0, 4},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			is := SplitToInt32(testCase.s, ",")
			assert.Equal(t, testCase.expected, is)
		})
	}
}
