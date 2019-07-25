// Ping pong
package main

import (
	"fmt"
	"time"
)

func pinger(ping chan int, pong chan int) {
	for {
		<-ping
		fmt.Println("ping")
		time.Sleep(1 * time.Second)
		pong <- 1
	}
}

func ponger(ping chan int, pong chan int) {
	for {
		<-pong
		fmt.Println("pong")
		time.Sleep(1 * time.Second)
		ping <- 1
	}
}

func runPingPong() {
	ping, pong := make(chan int), make(chan int)
	go pinger(ping, pong)
	go ponger(ping, pong)
	ping <- 1
	fmt.Scanln()
}

/*==============================================================*/

// this one is not finish yet
// reference: https://blog.golang.org/pipelines
package main

import (
	"fmt"
	"time"
)

func spread(main chan int, a chan int, b chan int, c chan int) {
	for {
		value := <-main
		a <- value
		b <- value
		c <- value
	}
}

func fanIn(out chan int) {
	var total int
	for num := range out {
		total += num
		fmt.Println("Total =", total)
	}
	fmt.Println("Finnal total =", total)
}

func printChannel(c chan int, out chan int, f func(int) int) {
	for {
		value := <-c
		x := f(value)
		out <- x
	}
}

func runFanIn() {
	main := make(chan int)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	out := make(chan int)

	go spread(main, a, b, c)
	go printChannel(a, out, func(i int) int {
		return i * 2
	})
	go printChannel(b, out, func(i int) int {
		return i * 3
	})
	go printChannel(c, out, func(i int) int {
		return i * 4
	})
	go fanIn(out)
	for i := 1; i <= 1; i++ {
		main <- i
	}
	time.Sleep(3 * time.Second)
	close(out)
	time.Sleep(3 * time.Second)
}

/*===============================================================*/

// go routing / search go scheduler
// channel
// blocking / non blocking channel
// waitGroup in package sync


// select sử dụng trong gửi nhận dữ liệu channel 


// fan out: input vao 1 so -> channel a print ra 1
//						   -> channel b print ra 2
//						   -> channel c print ra 3
// can paste value and action into channel

// fan in (merge channel)

// worker pool