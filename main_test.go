package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateRandomElements_ReturnNil(t *testing.T) {
	tests := []struct {
		name string
		size int
		want []int
	}{
		{"SizeZero", 0, nil},
		{"LessThanZero", -3, nil},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, generateRandomElements(tt.size), "Test: %s", tt.name)
	}

}

func Test_generateRandomElements_Size(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"SmallPositiveSize", 7, 7},
		{"MediumPositiveSize", 734, 734},
		{"BigPositiveSize", 30_001, 30_001},
	}
	for _, tt := range tests {
		assert.Len(t, generateRandomElements(tt.size), tt.want, "Test: %s", tt.name)
	}
}
func Test_generateRandomElements_ReallyRandom(t *testing.T) {
	assert.NotEqual(t, generateRandomElements(789), generateRandomElements(789), "Test: ReallyRandom")
}

func Test_maximum_EmptyOrNil(t *testing.T) {
	var data []int
	assert.Equal(t, -1, maximum(data))
	assert.Equal(t, -1, maximum(nil))
}
func Test_maximum_ReallyMax(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{"ArrayWithNegativeValue", []int{9, -87, 4, 0, 8, 67, -7, 7}, 67},
		{"ArrayWithOutNegativeValue", []int{9, 87, 4, 0, 8, 67, 7, 7}, 87},
		{"ArrayWithOneValue", []int{9}, 9},
		{"ArrayWithOneNegativeValue", []int{-19}, -19},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, maximum(tt.array), "Test: %s", tt.name)
	}
}

func Test_maxChunks_WithDifferentArray(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{"SmallLenArray", []int{9, 87, 4, 0}, 87},
		{"MediumLenArray", []int{999, 9, 87, 4, 0, 8, 67, 7, 7, 59, 2, 999, 1002}, 1002},
		{"WithNegativeValueArray", []int{9, -87, 4, 0, 8, 67, 7, 7}, 67},
		{"WithOneValueArray", []int{9}, 9},
		{"WithOneNegativeValueArray", []int{-19}, -19},
		{"WithOutValueArray", []int{}, -1},
		{"WithNilArray", nil, -1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, maxChunks(tt.array), "Test: %s", tt.name)
	}
}
