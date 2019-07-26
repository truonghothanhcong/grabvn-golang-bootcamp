package main

import (
	"fmt"
	"time"
)

func gen() <- chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 1; i <= 100; i++ {
			c <- i
		}
	}()

	return c
}

func do(n int) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Got: ", n)
}

func main()  {
	nums := gen()
	i 
	for num := range nums {
		// fmt.Println(num)
		go do(num)
	}
	time.Sleep(2 * time.Second)
}
