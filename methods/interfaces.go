package main
//this program demonstrates how to use interfaces
import( 
"fmt"
"math"
)


type Abser interface{
	Abs() float64
}

func main(){
	var a Abser
	f := MyFloat(-math,Sqrt2)
	v := Vertex{3,4}

	a = f //a MyFloat implements Abser
	a = &v //a vertex pointer implements Abser

	a = v //this V and since it's not a pointer it doesn't implment Abser
	//because its the actual v Vector

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) ABs() float64{
	if f< 0{
		return float64(-f)
	}
	return float64(f)
}

type Vertex strcut{
	X,Y float64
}

func (v *Vertex) Abs()float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}
