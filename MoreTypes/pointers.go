package main

import "fmt"

func main(){
	i,j:=42,2701
	p:=&i //create a pointer p and point it at i
	fmt.Println(*p) //print the value stored in p
        *p = 21 //set i to 21 through p
	fmt.Println(i) //see the value of i

	p = &j//point to j
	*p=*p/37 //divide j through the pointer
	fmt.Println(j)
}
