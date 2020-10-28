package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
	testChan()
	test2()
	test3()
}

func (ip IPAddr) String() string {
	var str []string
	for _, v := range ip {
		str = append(str, fmt.Sprint(v))
	}
	return strings.Join(str, ".")
}

var lock sync.Mutex

func testChan() {
	c := make(chan int)
	go func() {
		lock.Lock()
		fmt.Println("1a")
		c <- 1
		lock.Unlock()
	}()
	go func() {
		fmt.Println("2b")
		c <- 2
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func test2() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func test3() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
