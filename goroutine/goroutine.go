package main

import (
	"time"
	"fmt"
)
var c chan int
func pr(str string, i int) {
	time.Sleep(time.Second*time.Duration(i))
	fmt.Println(str,"is ready")
	c<-1
}
func main() {
	i:=1
	c=make(chan int)
	go pr("tea",3)
	go pr("water",2)
	go pr("orange",5)
	fmt.Println("im ok")
	L:
		for {
			select {
			case <-c:
				i++
				if i > 3 {
					break L
				}
			}
		}
	close(c)
}
