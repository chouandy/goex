package elasticex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunks := make([][]int, 0)
	chunkSize := 3
	for i := 0; i < len(array); i += chunkSize {
		end := i + chunkSize
		if end > len(array) {
			end = len(array)
		}
		chunks = append(chunks, array[i:end])
	}
	assert.Equal(t, []int{0, 1, 2}, chunks[0])
	assert.Equal(t, []int{3, 4, 5}, chunks[1])
	assert.Equal(t, []int{6, 7, 8}, chunks[2])
	assert.Equal(t, []int{9, 10}, chunks[3])
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunks := make([][]int, 0)
	chunkSize := 3
	for i := 0; i < len(array); i += chunkSize {
		<-ticker.C
		fmt.Println(i)
		end := i + chunkSize
		if end > len(array) {
			end = len(array)
		}
		chunks = append(chunks, array[i:end])
	}
	ticker.Stop()
	fmt.Println("ticker stopped")
	assert.Equal(t, []int{0, 1, 2}, chunks[0])
	assert.Equal(t, []int{3, 4, 5}, chunks[1])
	assert.Equal(t, []int{6, 7, 8}, chunks[2])
	assert.Equal(t, []int{9, 10}, chunks[3])
}
