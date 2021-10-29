package main

import "reflect"
import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d, expected %d for %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
