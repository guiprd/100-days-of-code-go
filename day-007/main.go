package main

import (
	"fmt"
	"time"
)

/*
Efficient Hamming number generator using the 3-pointer algorithm.
This implementation generates numbers in a separate goroutine and
streams them over a channel (concurrent producer). The algorithm is
O(n) time and O(n) memory, and scales well for large `n`.
*/

func main() {
	t := time.Now()
	fmt.Println(Hammer(1000))
	fmt.Println(time.Since(t).Seconds())
}

// Hammer returns the n-th smallest Hamming number (1-based n).
func Hammer(n int) uint {
	if n <= 0 {
		return 0
	}
	ch := generateHamming(n)
	var val uint
	for v := range ch {
		val = v
	}
	return val
}

// generateHamming produces the first n Hamming numbers on a channel.
// It runs the generation in a goroutine (producer), so the caller can
// consume concurrently if desired.
func generateHamming(n int) <-chan uint {
	ch := make(chan uint)
	go func() {
		defer close(ch)
		if n <= 0 {
			return
		}
		h := make([]uint, n)
		h[0] = 1
		ch <- 1
		i2, i3, i5 := 0, 0, 0
		for idx := 1; idx < n; idx++ {
			next2 := h[i2] * 2
			next3 := h[i3] * 3
			next5 := h[i5] * 5
			next := min(next2, next3, next5)
			h[idx] = next
			if next == next2 {
				i2++
			}
			if next == next3 {
				i3++
			}
			if next == next5 {
				i5++
			}
			ch <- next
		}
	}()
	return ch
}

func min(a, b, c uint) uint {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}
