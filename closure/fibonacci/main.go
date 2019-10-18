package main

import "fmt"

func fibonacci() func() int64 {
	x := int64(0)
	y := int64(1)
	return func() int64 {
		r := x
		x, y = y, x+y
		return r
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
