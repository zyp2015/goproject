package main

import (
	"fmt"
)

type S struct {
	i int
}

func (p *S)Get() int  {
	return  p.i
}

func (p *S)Set(v int)  {
	p.i=v
}

type I interface {
	Get() int
	Set(int)
}

func f(i I)  {
	fmt.Println(i.Get())
	i.Set(100)
	fmt.Println(i.Get())
}
func main() {
	i:=new(S)
	f(i)
}
