package main

import "testing"

func TestSum(t *testing.T) {
    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}

        got := sum(numbers)
        expected := 15

        if got != expected {
            t.Errorf("got %d, expected %d for %v", got, expected, numbers)
        }
    })
}
