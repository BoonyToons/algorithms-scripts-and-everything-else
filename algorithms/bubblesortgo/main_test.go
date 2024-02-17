package main_test

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Test sorting with unsorted array"), func(t *testing.T) {
		input := []int{6, 5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5, 6}

		bubbleSort(input)

		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Expected %v, got %v", expected, input)
		}
	}
}