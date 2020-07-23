package main 

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8{
	pic:=make([][]uint 8,dy)
	
	for i:=0; i<dy;i++{
		pic[i] = make([]uint8,dx)
	}
	for y,_:= range pic{
	    for x,_:=range pic[y]{
		pic[y][x] =  uint8(x)^2*(uint8(x)*uint8(y))
	     }
	}
	return pic
}

func main(){
	pic.Show(Pic)
}
