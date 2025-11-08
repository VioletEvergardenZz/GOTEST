package chanTest

import (
	"fmt"
	"time"
)

func ChanTest2() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "🐑"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "🐂"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
	//阻塞
	/**
	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
	*/
}
