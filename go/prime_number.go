package main

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

//func main() {
//	ch := make(chan int) // Create a new channel.
//	go generate(ch)      // Start generate() as a goroutine.
//	for {
//		prime := <-ch
//		fmt.Println(prime)
//		ch1 := make(chan int)
//		go filter(ch, ch1, prime)
//		ch = ch1
//	}
//}

//func main() {
//	origin, wait := make(chan int), make(chan struct{})
//	Processor(origin, wait)
//	for num := 2; num < 10000; num++ {
//		origin <- num
//	}
//	close(origin)
//	<-wait
//}

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
	}()
}
