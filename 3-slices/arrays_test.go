package main

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
