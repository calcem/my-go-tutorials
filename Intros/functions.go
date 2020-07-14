package main

import "fmt"
/*
func add(x int, y int ) int{
	return x+y
}
*/
//this is to show that you can declare the type of multiple arguments in sequence
func add(x,y int) int{
	return x+y
}

func main(){
	fmt.Println(add(2,2))
}
