package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev:=0
	current := 0
	next:= prev + current
	return func () int{
		if prev == 0 && current == 0 {
			current += 1
		} else{
			prev = current
			current = next
			next = prev + current
		}
		return next
	}
			
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

