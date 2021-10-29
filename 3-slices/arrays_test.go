package main

import "reflect"
import "testing"

func testSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := sum(numbers)
	expected := 6

	if got != expected {
		t.Errorf("got %d, expected %d for %v", got, expected, numbers)
	}
}

func testSumAll(t *testing.T) {
	got := sumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func testSumAllTails(t *testing.T) {
	got := sumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
