package main

import "fmt"

type rect struct {
	width, height int
}

func (r * rect) area() int  {
	return r.width * r.height
}

func ( r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r rect) change1() {
	r.width = 0
}

func (r *rect) change2(){
	r.width = 0
}

func main() {
	r := rect{width : 10, height : 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	r.change1()
	fmt.Println(r)

	rp.change2()
	fmt.Println(r)
	fmt.Println(rp)
	fmt.Println(*rp)
}