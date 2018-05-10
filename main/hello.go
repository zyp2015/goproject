package main

import (
	"fmt"
	"stringkoko"
	"math/rand"
	"math"
	"runtime"
	"strings"
)
func add(x ,y int) int{
	return x+y
}

func swap(s1,s2 string) (string,string){
	return s2,s1
}
type point struct {
	x int
	y int
}
func (v *point)add() int {
	return v.x+v.y
}
var a1, a2  = 1, 2

func adder() func(int) int {
	sum:=0
	return func (i int) int {
		sum+=i
		return sum
	}
}
func main() {
	str:=111

	fmt.Println(stringkoko.Reverse("hello sbsbsbsb"))
	fmt.Println(math.Pi)
	fmt.Printf("%d %f",rand.Intn(110),math.Nextafter(2,3))
	fmt.Println(add(1,3))
	fmt.Println(swap("ni ge sb","nige cao nima"))
	fmt.Println(a1,a2)
	fmt.Printf("this is %T\n",str)
	for count := 0; count <= 10; count++{
		//do no thing
		if count == 5 {
			fmt.Println(count)
		}
		//defer fmt.Println(count)
	}
	//defer fmt.Println("you are sb")
	switch os:=runtime.GOOS;os {
	case "linux":
		fmt.Println("this is linux")
	default:
		fmt.Println(os)
	}
	fmt.Println(point{55,22})

	p:=[]int{22,3,5,6,8,9,10,12}
	fmt.Println("p==",p)
	/*for i := 0; i < len(p); i++ {
		fmt.Println("p[i]=",p[i])
	}*/
	fmt.Println("p[1:3]=",p[1:3])
	fmt.Println("p[1:]=",p[1:])
	fmt.Println("p[3:]=",p[3:])

	a:=make([][6]int,6)
	fmt.Println(a,len(a),cap(a))
	for i,v:=range a{
		for j,_:= range v{
			a[i][j]=i*j
		}
	}
	fmt.Println(a)
	var m = map[int]int{
		1:11,
		3:22,
	}
	//m=make(map[string]point)
	//m["sb"]=point{78,88}
	m[2]=122
	delete(m,1)
	elem,ok:=m[3]
	fmt.Println(elem,ok)
	fmt.Println(m[3])
	fmt.Println(WordCount("you are sb sb is you"))

	dd:= func(x,y int) int {
		return (x*x+y*y)
	}
	fmt.Println(dd(12,33))
	pos,res:=adder(),adder()
	for i := 0; i < 4; i++ {
		fmt.Println(pos(i),res(-2*i))
	}
	fmt.Println(res(9))
	f:=fin()
	fmt.Println(f(),f(),f())
	f=fin()
	fmt.Println(f(),f(),f())

	test:=&point{99,44}
	fmt.Println(test.add())
	ip := map[int](IPAddr){
		1:{1,1,1,1},
		2:{2,2,2,2},
	}
	for i,v:=range ip{
		fmt.Printf("%v:%v\n",i,v)

	}
	r:=strings.NewReader("A")
	b:=make([]byte, 2)
	n,err :=r.Read(b)
	fmt.Printf("%c",b[:n])
	if err==nil{
		fmt.Println("error")
	}

}
func fin() func( ) int{
	pre:=0
	next:=1
	count:=0
	sum:=0
	return func() int {
		count++
		if count<2{
			return count
		}
		sum=pre+next
		pre=next
		next=sum
		return sum
	}
}
func WordCount(str string) map[string] int {
	s_array:=strings.Fields(str)
	m:=make(map[string]int)
	for i:=0;i<len(s_array);i++ {
		_, ok := m[s_array[i]]
		if ok == false {
			m[s_array[i]] = 1
		} else {
			m[s_array[i]]++
		}
	}
	return  m
}

type IPAddr [4]byte

func (addr IPAddr)String() string  {
	return  fmt.Sprintf("%d.%d.%d.%d",addr[0],addr[1],addr[2],addr[3])
}