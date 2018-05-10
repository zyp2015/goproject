package main

import (
	"os"
	"fmt"
)

func main() {
	bufer:=make([]byte,100)
	f,_:=os.Open("D:/testgo/test.txt")
	defer f.Close()
	for {
		n, _ := f.Read(bufer)
		if n != 0 {
			os.Stdout.Write(bufer[:n])
		}else if n == 0 {
			fmt.Println("read 0 byte")
			break
		}
	}
	fmt.Println(string(bufer))
}
