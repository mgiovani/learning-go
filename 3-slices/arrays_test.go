package main

import "reflect"
import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := sum(numbers)
	expected := 6

	if got != expected {
		t.Errorf("got %d, expected %d for %v", got, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
    checkSums := func(t testing.TB, got []int, want []int) {
        t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
    }

	t.Run("return sums of slices", func(t *testing.T) {
		got := sumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

        checkSums(t, got, want)
	})

	t.Run("return sums for empty slices", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}

        checkSums(t, got, want)
	})
}
