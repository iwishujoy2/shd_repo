package main

import (
	"context"
	"fmt"
	"os"
)

func fibs(ctx context.Context, n int) chan int {
	ch := make(chan int) // Create channel to pass result back to initial thread
	go func() {          // Create a lightweight thread
		defer close(ch) // Clean up and, close this channel when we exit
		a, b := 1, 1
		for i := 0; i < n; i++ {
			select {
			case ch <- a:
				a, b = b, a+b
			case <-ctx.Done():
				fmt.Println(" cancelled")
			}
		}
	}()
	return ch
}

func main() {
	n := 0

	fmt.Println("Enter in the Fib sequence number")
	fmt.Scanf("%d", &n)

	if n > 92 {
		fmt.Println("Input number is to large to calculate squence!")
		fmt.Println("Try a number less then 93")
		os.Exit(-1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := fibs(ctx, n)
	for i := 0; i < n; i++ {
		val := <-ch
		fmt.Printf("% d ", val)
	}

	fmt.Println()
	cancel()
}
