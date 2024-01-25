package main
import "fmt"

//shape   interface
type shape interface

{

	Area() float64
}



//shape2 interface
type shape2 interface {
Area2() float64
perim() float64

}

// struct rectangle
type Rect	struct{
	width, height float64
}
// struct triangle
type Triangle	struct{
	base, height float64
}


func (t Triangle) Area() float64{
	return t.base * t.height / 2
}



 func (r Rect) Area() float64 {
	return r.width * r.height
 }



// calculate
func calculateArea(s shape) float64 {

    return s.Area()
}

func main() {

fmt.Println("let the execution begin ...")
Rect := Rect{10,5}
Triangle := Triangle{10,5}
fmt.Println(calculateArea(Rect))
fmt.Println(calculateArea(Triangle))

}