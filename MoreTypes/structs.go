package main
import "fmt"

type Vertex struct{
	X int
	Y int
}

var(
	v1 = Vertex{1,2}
	v2 = Vertex{X: 1} //so you will have X = 1 and Y = 0
	v3 = Vertex{}
	p = &Vertex{1,2}
)	

func main() {
//	v:=Vertex{1,2}
//	p:=&v //create a pointer and point to v
//	p.X = 1e9 //change the value of X from 1 to 1e9
//	fmt.Println(v)
	fmt.Println(v1,p,v2,v3)
}
