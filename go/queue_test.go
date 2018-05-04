package main

import (
	"testing"
)

func TestJosephus(t *testing.T) {
	//r := ring.New(6)
	//
	//// Get the length of the ring
	//n := r.Len()
	//
	//// Initialize the ring with some integer values
	//for i := 0; i < n; i++ {
	//	r.Value = i
	//	r = r.Next()
	//}
	//
	//// Unlink three elements from r, starting from r.Next()
	//r.Unlink(2)
	//
	//// Iterate through the remaining ring and print its contents
	//r.Do(func(p interface{}) {
	//	fmt.Println(p.(int))
	//})

	Josephus(6, 3, 1)
}
