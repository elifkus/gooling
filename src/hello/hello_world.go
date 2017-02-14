package main 

import (
	"fmt"
)

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	fmt.Println("Hello world!")
	
	var xPointer = new(int)
	var yPointer = new(int)
	*xPointer = 3
	*yPointer = 4
	
	swap(xPointer, yPointer)
	
	fmt.Printf("x: %d y: %d", *xPointer, *yPointer)
	
}

