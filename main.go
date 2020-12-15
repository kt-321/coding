package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	//fmt.Printf("%T: %v\n", <-ch, <-ch)
	fmt.Println(<-ch)
	//fmt.Printf("%T: %v\n", <-ch, <-ch)

	fmt.Println("++++++++++++++++")

	//c := make(chan int, 10)
	//go fibonacci(cap(c), c)
	//for i:= range c {
	//	fmt.Printf("%T: %v\n", i, i)
	//}


	//c2 := make(chan int)
	//quit := make(chan int)
	//
	//go func() {
	//	for i:=0; i<10; i++ {
	//		fmt.Println(<-c2)
	//	}
	//	quit <- 0
	//}()
	//fibonacci2(c2, quit)



	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick")
		case <- boom:
			fmt.Println("boom")
		default: // defaultがあることで、ブロックせずに送受信する
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i:=0; i<n; i++ {
		fmt.Println("i", i)
		c <- x
		x, y = y, x+y
	}
	close(c) // これがないとmainルーチンのrangeループを終了できず all goroutines are asleep - deadlock!
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	// 複数あるcaseのいずれかが準備できるようになるまでブロック
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
