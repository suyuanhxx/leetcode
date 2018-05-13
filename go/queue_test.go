package main

import (
	"testing"
	"fmt"
	"container/ring"
)

func TestJosephus(t *testing.T) {
	r := ring.New(6)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Unlink two elements from r, starting from r.Next()
	r.Unlink(2)

	// Iterate through the remaining ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	fmt.Println("-------------")
	fmt.Println(Josephus(6, 3, 1))
}
