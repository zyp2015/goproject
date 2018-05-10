package main

import (
	"fmt"
)

//e.g:	//f:=[]float64{10.1,1.51,55.5}
func avg(x []float64) float64 {
	var sum float64
	for _,v :=range x{
		sum+=v
	}
	return sum/float64(len(x))
}
//e.g:fmt.Println(order(551,66))
func order(x, y int) (r1,r2 int){
	if x>y{
		return y,x
	}else {
		return x,y
	}
}

//stack
type stack struct {
	length int;

}

func pr(value ...int) {
	for _,v:=range value{
		fmt.Println(v)
	}
}
func fn(x int)  {
	fmt.Println(x)
	var sum,next,pre int = x,x,x
	for i := 0;i<8 ; i++ {
		sum=pre+next
		pre=next
		next=sum
		fmt.Println(sum)
	}
}
func xx2(x int) int {
	return x*x
}
func mymap(f func(x int) int,arr []int)[]int {
	temp:=make([]int,len(arr))
	for i,v:=range arr{
		temp[i]=f(v)
	}
	return temp
}
func main() {
	fmt.Println("learn start")
	//pr(11,22,33,44,55,66,77)
	fn(8)
	a:=[]int{1,3,4,5,6,7,8}
	fmt.Println(a)
	fmt.Println(mymap(xx2,a))
	fmt.Println(a)
}
