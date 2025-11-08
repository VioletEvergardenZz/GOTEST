package chanTest

import (
	"fmt"
	"time"
)

func ChanTest1() {
	c := make(chan string)
	go count(3, "🐑", c)
	for message := range c {
		fmt.Println(message)
	}
}
func count(num int, animal string, c chan string) {
	for i := 0; i < num; i++ {
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
