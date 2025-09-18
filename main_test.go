package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateRandomElements(t *testing.T) {
	tests := []struct {
		name      string
		size      int
		want      []int
		want_size int
	}{
		{name: "SizeZero", size: 0, want: nil},
		{name: "LessThanZero", size: -3, want: nil},
		{name: "SmallPositiveSize", size: 7, want_size: 7},
		{name: "MediumPositiveSize", size: 734, want_size: 734},
		{name: "BigPositiveSize", size: 30_001, want_size: 30_001},
	}
	for _, tt := range tests {
		if tt.want_size == 0 {
			assert.Equal(t, tt.want, generateRandomElements(tt.size), "Test: %s", tt.name)
		}
		assert.Len(t, generateRandomElements(tt.size), tt.want_size, "Test: %s", tt.name)
	}

}

// func Test_generateRandomElements_ReallyRandom(t *testing.T) {
// 	assert.NotEqual(t, generateRandomElements(789), generateRandomElements(789), "Test: ReallyRandom")
// }

func Test_maximum(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{"SmallLenArray", []int{9, 87, 4, 0}, 87},
		{"MediumLenArray", []int{999, 9, 87, 4, 0, 8, 67, 7, 7, 59, 2, 999, 1002}, 1002},
		{"WithOneValueArray", []int{9}, 9},
		{"WithOutValueArray", []int{}, 0},
		{"WithNilArray", nil, 0},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, maxChunks(tt.array), "Test: %s", tt.name)
	}
}

func Test_maxChunks(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{"SmallLenArray", []int{9, 87, 4, 0}, 87},
		{"MediumLenArray", []int{999, 9, 87, 4, 0, 8, 67, 7, 7, 59, 2, 999, 1002}, 1002},
		{"WithOneValueArray", []int{9}, 9},
		{"WithOutValueArray", []int{}, 0},
		{"WithNilArray", nil, 0},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, maxChunks(tt.array), "Test: %s", tt.name)
	}
}
